package iof

import (
	"io/ioutil"
	"net/http"
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
	resp, err := http.Get(file_url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	fd, err := ioutil.ReadAll(resp.Body)
	return string(fd)
}
