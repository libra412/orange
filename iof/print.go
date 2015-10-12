package iof

import (
	"io"
	"os"
)

func WriteToFile(body, path string) bool {
	fi, _ := os.Create(path)
	defer fi.Close()
	n, err := fi.WriteString(body)
	if n > 0 && nil == err {
		return true
	}
	return false
}

func WriteToBinaryFile(filename string, body io.Reader) int64 {
	file, _ := os.Create(filename)
	defer file.Close()
	if size, err := io.Copy(file, body); err == nil {
		return size
	} else {
		panic(err)
	}
	return 0
}

// 检查文件或目录是否存在
// 如果由 filename 指定的文件或目录存在则返回 true，否则返回 false
func Exist(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil || os.IsExist(err)
}
