package utils

import (
	"fmt"
	"github.com/libra412/orange/iof"
	"io/ioutil"
	"net/http"
	"regexp"
	"runtime"
	"strconv"
	"sync"
	"time"
)

func Grep(url, rule string) (strs [][]string) {

	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("User-Agent", `Mozilla/5.0 (iPhone; CPU iPhone OS 8_0 like Mac OS X) AppleWebKit/600.1.3 (KHTML, like Gecko) Version/8.0 Mobile/12A4345d Safari/600.1.4`)
	resp, err := client.Do(req)
	defer resp.Body.Close()
	if err == nil {
		respBody, _ := ioutil.ReadAll(resp.Body)
		body := string(respBody)

		reg := regexp.MustCompile(rule)
		strs = reg.FindAllStringSubmatch(body, -1)
		//iof.WriteToFile(body, "/Users/apple/Desktop/test.html")
		// count := len(strs)
		// content := ""
		// for i := 0; i < count; i++ {
		// 	content = content + strs[i][1] + "\n"
		// }
		// fmt.Println(content)
		return
	}
	return

}

func grepMp4(filename, url string) {
	fmt.Println("begin download " + url)
	resp, err := http.Get(url)
	defer resp.Body.Close()
	if err == nil {
		size := iof.WriteToBinaryFile(filename, resp.Body)
		fmt.Println("下载完成  " + fmt.Sprint(size/1024) + "m")
		waitgroup.Done()
		runtime.Goexit()
	}
}

var waitgroup sync.WaitGroup

func GoDownload(strs []string) {
	count := len(strs)
	fmt.Println(count)
	for i := 0; i < count; i++ {
		filename := "/Users/apple/Desktop/" + strconv.FormatInt(time.Now().UnixNano(), 10) + ".mp4"
		waitgroup.Add(1)
		go grepMp4(filename, strs[i])
	}
	waitgroup.Wait()

}
