package iof

import (
	"encoding/base64"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"time"
)

// 文件base64序列化
func FileToBase64String(path string) string {
	stream := Read(path)
	dst := base64.StdEncoding.EncodeToString([]byte(stream))
	return dst
}

// 文件转base64
func FilepathToBase64String(filepath string) string {
	stream, err := ioutil.ReadFile(filepath)
	if err != nil {
		return ""
	}
	dst := base64.StdEncoding.EncodeToString(stream)
	return dst
}

// 该路径是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// 获取文件修改时间 返回unix时间戳
func GetFileModTime(path string) int64 {
	f, err := os.Open(path)
	if err != nil {
		log.Println("open file error")
		return time.Now().Unix()
	}
	defer f.Close()

	fi, err := f.Stat()
	if err != nil {
		log.Println("stat fileinfo error")
		return time.Now().Unix()
	}
	return fi.ModTime().Unix()
}

// 读取文件夹下所有文件
func ReadFileListFromDir(path string) ([]fs.FileInfo, int) {
	fileInfoList, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}
	return fileInfoList, len(fileInfoList)
}
