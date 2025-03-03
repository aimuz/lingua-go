package sv

import (
	"embed"
	"github.com/pemistahl/lingua-go"
)

//go:embed *.zip
var model embed.FS

func init() {
	lingua.Register("sv", model)
}
