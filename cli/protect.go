package cli

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/drewart/hen/internal/server"
	"github.com/spf13/cobra"
)

var protectCmd = &cobra.Command{
	Use:   "protect",
	Short: "protect",
	Long:  `protect encrypts a string`,
	Run: func(cmd *cobra.Command, args []string) {
		host, _ := cmd.Flags().GetString("host")
		message, _ := cmd.Flags().GetString("message")
		runProtect(host, message)
	},
}

func init() {
	rootCmd.AddCommand(protectCmd)
	protectCmd.Flags().StringP("message", "m", "", "message to encrypt")
	protectCmd.Flags().StringP("host", "H", "http://localhost:3000", "host name and port")
}

func runProtect(host, message string) {
	var msg server.ProtectRequest
	msg.Message = message

	body, _ := json.Marshal(msg)

	postBody := bytes.NewBuffer(body)

	req, _ := http.NewRequest("POST", fmt.Sprintf("%s/protect", host), postBody)
	req.Header.Add("token", "foobar")
	// resp, err := http.Post(fmt.Sprintf("%s/protect", host), "application/json", postBody)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	// fmt.Println(resp.Status)

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}

	sb := string(respBody)
	fmt.Println(sb)
}
