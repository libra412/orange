package iof

import (
	"bufio"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
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

//read file by line and handler func to do it.
func ReadLine(fileName string, handler func(string)) error {
	f, err := os.Open(fileName)
	if err != nil {
		return err
	}
	buf := bufio.NewReader(f)
	for {
		line, err := buf.ReadString('\n')
		line = strings.TrimSpace(line)
		handler(line)
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
	}
	return nil
}
