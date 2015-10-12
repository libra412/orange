package iof

import (
	"io/ioutil"
	"os"
)

//read by the file path
func Read(file_path string) string {
	fi, err := os.Open(file_path)
	if err != nil {
		panic(err)
	}
	defer fi.Close()
	fd, err := ioutil.ReadAll(fi)
	return string(fd)
}

//read the file by the url
func ReadUrl(file_url string) string {
	//TODO
	return ""
}
