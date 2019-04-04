package main

import (
    "os"
    "fmt"
)

func getcs() string {
	//获取启动参数
    argsWithProg := os.Args
    a := argsWithProg[1]
    return a
}


func main() {
	fmt.Println("v2ray 启动模块")

}
