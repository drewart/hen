package hen

import (
	"fmt"
)

type Egg struct {
	secureText string
	Encrypted  string
	md5        string
}

func (e *Egg) ToString() string {
	return fmt.Sprintf("egg(%s,%s)", e.Encrypted, e.md5)
}

type Hen struct {
	Host     string
	FilePath string
	Token    string
	Eggs     []Egg
}
