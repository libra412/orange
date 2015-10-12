package system

import (
	"bytes"
	"fmt"
	"log"
	"os/exec" //这个包是主要用来调用cmd命令
)

//调用系统指令的方法，参数s 就是调用的shell命令
func system(s string) {
	cmd := exec.Command("/bin/sh", "-c", s) //调用Command函数
	var out bytes.Buffer                    //缓冲字节

	cmd.Stdout = &out //标准输出
	err := cmd.Run()  //运行指令 ，做判断
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", out.String()) //输出执行结果
}
