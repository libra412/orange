package main

import (
	"fmt"
	"github.com/libra412/orange/utils"

	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	rule1 := "<iframe src='(.*?[^']?)'"
	rule2 := `<source src="(.*?[^"]?)"`
	//rule3 := `<video controls src="(.*?[^"]?)"`
	mainUrl := "http://sunshine-afternoon-tea.tumblr.com/" //"http://fulitv.tumblr.com/"
	strs := utils.Grep(mainUrl, rule1)
	var strs3 []string
	count := len(strs)
	for i := 0; i < count; i++ {
		strs2 := utils.Grep(strs[i][1], rule2)
		strs3 = append(strs3, strs2[0][1])
	}
	fmt.Println("==========begin download")
	fmt.Println(strs3)
	utils.GoDownload(strs3)
	fmt.Println("b")

}
