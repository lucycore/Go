package main

import (
	"fmt"
	"strconv"
	"time"
	"math/rand"
)

func sjs(fw int) int {

    rand.Seed(time.Now().Unix())

    rnd := rand.Intn(fw)

    return rnd
} 

func main() {
	fmt.Println("请输入随机数生成范围:")
	fw := ""
	fmt.Scanln(&fw)

	fwi,_ := strconv.Atoi(fw)

	sjss := sjs(fwi)

	js := 0

	zccs := 0
	pd := 0

	for pd == 0 {

		fmt.Println("请输入您猜测的数字：")
		css := ""
		fmt.Scanln(&css)
		cs,_ := strconv.Atoi(css)

		if zccs == cs {

			fmt.Println("猜重复了！")

		} else {

			switch {
			case sjss < cs :
				fmt.Println("大了！")
				js += 1

			case sjss > cs :
				fmt.Println("小了！")
				js += 1

			case sjss == cs :
				fmt.Println("猜对了！")
				js += 1
				pd = 1
			}
			zccs = cs

		}

	}
	fmt.Println("一共猜了",js,"次")

}