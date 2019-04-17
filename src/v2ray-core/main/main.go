package main

import (
    //"os"
    "fmt"
    "v2ray-core/core"
    "v2ray-core/download"
    "time"
    "os"
)

var sysname string

func main() {

	sysname = core.Check_system()

	if sysname == "darwin" {

		jsonurl := "http://www.lucycore.top/static/pz/v2_gy.json"
		v2zip := "http://www.lucycore.top/static/v2ray-download/v2rayMac.zip"

		fmt.Println("v2ray Golang macos 内部版本V3.0")
		fmt.Println("适用时间截至2019年5月31日")

		if core.Test_time() {
			fmt.Println("本地时间验证通过！")


			if core.Yz_v2_wz() {

				fmt.Println("正在尝试获取配置文件")
				bdlj,_,_ := core.Get_lj()
				download.Download(jsonurl,bdlj)
				core.Start_v2()
				fmt.Println("已启动v2ray")

			} else {

				fmt.Println("正在下载v2ray需要一些时间")
				_,ysb,_ := core.Get_lj()
				download.Download(v2zip,ysb)
				_,_,jyf := core.Get_lj()
				fmt.Println(v2zip + " " + ysb +  " " + jyf)
				core.Unzip(ysb,jyf)
				fmt.Println("以解压完成！")
				fmt.Println("请您手动重启此程序！")
				fmt.Println("正在退出程序！3秒后退出！")
				time.Sleep(time.Duration(3)*time.Second)
				os.Exit(10)

			}

		} else {
			fmt.Println("本地时间验证失败！")
			fmt.Println("正在退出程序！3秒后退出！")
			time.Sleep(time.Duration(3)*time.Second)
			os.Exit(10)

		}
		

	}

	if sysname == "windows"{

		jsonurl := "http://www.lucycore.top/static/pz/v2_gy.json"
		v2zip := "http://www.lucycore.top/static/v2ray-download/v2rayWin.zip"

		fmt.Println("v2ray Golang windows 内部版本V3.0")
		fmt.Println("适用时间截至2019年5月31日")

		if core.Test_time() {
			fmt.Println("本地时间验证通过！")


			if core.Yz_v2_wz() {

				fmt.Println("正在尝试获取配置文件")
				bdlj,_,_ := core.Get_lj()
				download.Download(jsonurl,bdlj)
				core.Start_v2()
				fmt.Println("已启动v2ray")

			} else {

				fmt.Println("正在下载v2ray需要一些时间")
				_,ysb,_ := core.Get_lj()
				download.Download(v2zip,ysb)
				_,_,jyf := core.Get_lj()
				core.Unzip(ysb,jyf)
				fmt.Println("以解压完成！")
				fmt.Println("请您手动重启此程序！")
				fmt.Println("正在退出程序！3秒后退出！")
				time.Sleep(time.Duration(3)*time.Second)
				os.Exit(10)

			}

		} else {
			fmt.Println("本地时间验证失败！")
			fmt.Println("正在退出程序！3秒后退出！")
			time.Sleep(time.Duration(3)*time.Second)
			os.Exit(10)

		}


	}


}
