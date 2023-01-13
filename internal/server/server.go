package server

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

var (
	encryptionKey string
	tokens        []string
)

func Setup() error {
	encryptionKey = os.Getenv("ENCRYPTION_KEY")
	if encryptionKey == "" {
		return fmt.Errorf("variable not set ENCRYPTION_KEY")
	}
	tokens = append(tokens, "foobar")
	return nil
}

type ProtectRequest struct {
	Message string `json:"message" binding: "required"`
}

type HatchRequest struct {
	Egg string `json:"egg" binding: "required"`
}

type HatchResponse struct {
	Secret string `json:"secret"`
}

func SetupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	// Ping test
	r.GET("/health", func(c *gin.Context) {
		c.String(http.StatusOK, "ok")
	})
	// protect
	r.POST("/protect", func(c *gin.Context) {
		valid := false
		token := c.GetHeader("token")
		if token == "" {
			c.String(http.StatusBadRequest, "token missing")
			return
		}

		// find token
		for _, t := range tokens {
			if t == token {
				valid = true
				break
			}
		}

		if !valid {
			log.Println("Not valid token")
			c.String(http.StatusBadRequest, "token not found")
			return
		}

		var message ProtectRequest
		c.BindJSON(&message)

		myCipher, err := aes.NewCipher([]byte(encryptionKey))
		if err != nil {
			log.Println(err)
			c.String(http.StatusBadRequest, "bad")
			return
		}
		gcm, err := cipher.NewGCM(myCipher)
		if err != nil {
			log.Println(err)
			c.String(http.StatusBadRequest, "bad request")
		}
		nonce := make([]byte, gcm.NonceSize())

		if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
			log.Println(err)
			c.String(http.StatusBadRequest, "bad request")
		}
		secret := gcm.Seal(nonce, nonce, []byte(message.Message), nil)
		// hex := binary.BitConverter.ToString(secretText)
		secretHex := hex.EncodeToString(secret)
		log.Println(secretHex)
		// c.String(http.StatusOK, secretHex)
		eggTxt := fmt.Sprintf("egg(%s,%x)", secretHex, md5.Sum([]byte(message.Message)))
		c.JSON(200, gin.H{"egg": eggTxt})
	})

	r.POST("/hatch", func(c *gin.Context) {
		var hatchRequest HatchRequest
		c.BindJSON(&hatchRequest)
		if hatchRequest.Egg == "" {
			c.String(http.StatusBadRequest, "egg empty")
			return
		}
		secretBytes, err := hex.DecodeString(hatchRequest.Egg)
		// ciphertext, err := ioutil.ReadFile("myfile.data")
		// if our program was unable to read the file
		// print out the reason why it can't
		if err != nil {
			log.Println(err)
			c.String(http.StatusBadRequest, "hex decode error")
			return
		}
		// ciphertext, err := ioutil.ReadFile("")
		mycipher, err := aes.NewCipher([]byte(encryptionKey))
		if err != nil {
			fmt.Println(err)
		}

		gcm, err := cipher.NewGCM(mycipher)
		if err != nil {
			fmt.Println(err)
		}

		nonceSize := gcm.NonceSize()
		if len(secretBytes) < nonceSize {
			fmt.Println("bytes > nonceSize length")
		}

		nonce, secretBytes := secretBytes[:nonceSize], secretBytes[nonceSize:]
		plaintext, err := gcm.Open(nil, nonce, secretBytes, nil)
		if err != nil {
			fmt.Println(err)
			c.String(http.StatusBadRequest, "bad decrypt")
			return
		}
		log.Println(string(plaintext))
		c.JSON(200, gin.H{"secret": string(plaintext)})
	})
	return r
}
