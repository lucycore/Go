package main

import (
    //"os"
    "fmt"
    "v2ray-nb/core"
    "v2ray-nb/download"
    "time"
    "os"
)


func main() {

	//jsonurl := "http://www.lucycore.top/static/"
	jsonurl := "http://www.lucycore.top/zzz/v2ray/v2_sn_1.json"


	fmt.Println("v2ray Golang 内部版本V3.0")
	fmt.Println("适用时间截至2019年5月31日")

	if core.Test_time() {
		fmt.Println("本地时间验证通过！")
		fmt.Println("正在尝试获取配置文件")

	bdlj := core.Get_lj()
	download.Download(jsonurl,bdlj)
	core.Start_v2()
	fmt.Println("已启动v2ray")


	} else {
		fmt.Println("本地时间验证失败！")
		fmt.Println("正在退出程序！3秒后退出！")
		time.Sleep(time.Duration(3)*time.Second)
		os.Exit(10)

	}



}
