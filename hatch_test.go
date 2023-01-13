package hen

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"
	"testing"
)

// Test
func TestHatch(t *testing.T) {
	input := `{"foo":"egg(abc)"}`
	d := os.TempDir()
	filePath := path.Join(d, "test1.json")
	ioutil.WriteFile(filePath, []byte(input), 0o660)
	Hatch(filePath)

	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		t.Error(err)
	}

	dataStr := string(data)

	if strings.Contains(dataStr, "egg(") {
		t.Error("egg not replaced")
	}
	fmt.Println(dataStr)
}
