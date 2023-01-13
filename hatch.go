package hen

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
)

var eggMatch = regexp.MustCompile(`egg\(\S+(,\S+|)\)`)

// https://tutorialedge.net/golang/go-encrypt-decrypt-aes-tutorial/

func Hatch(filepath string) error {
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		fmt.Println("error openning file", filepath)
	}

	newData := eggMatch.ReplaceAllFunc(data, func(b []byte) []byte {
		fmt.Println(string(b))
		return []byte("xxx")
	})

	err = os.Rename(filepath, filepath+".bak")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(filepath, newData, 0o600)
	if err != nil {
		return fmt.Errorf("failed to write to %s %v", filepath, err)
	}

	return nil
}
