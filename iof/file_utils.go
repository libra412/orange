package iof

import (
	"encoding/base64"
	// "io/ioutil"
	// "os"
)

//
func FileToBase64String(path string) string {
	stream := Read(path)
	var dst []byte
	base64.StdEncoding.Encode(dst, []byte(stream))
	return string(dst)
}
