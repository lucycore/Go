package main

import (
	"fmt"
	"strconv"
)

func welcome() (string,int){
	fmt.Println("欢迎来到我们的餐厅！")
	fmt.Println("接下来我们将向您索取您的信息")
	name := ""
	number := 0
	fmt.Println("您的名称：")
	fmt.Scanln(&name)
	fmt.Println("一共有几位：")
	fmt.Scanln(&number)
	return name,number
}

type food struct {
	name string
	number int
	jg int
}

func get_food_list() {
	food_list := [3]string {"冰激凌","锦江肉丝","酸辣粉"}
	for x,y := range food_list {
		xa := x + 1
		xb := strconv.Itoa(xa)
		fmt.Println(xb,",",y)
	}
}
	

func main() {

	name, nu :=welcome()
	fmt.Println(name)
	fmt.Println(nu)
	get_food_list()

}
