package main

import (
	"fmt"
	"strconv"
)

func main() {
	var nu string

	fmt.Println("请输入一个数字")
	fmt.Scanln(&nu)

	nuint,_ := strconv.Atoi(nu)

	var he int
	var js int

	for js := 0 ;  js <= nuint ; js += 1 {
		he = he + js 
	}

	fmt.Println(he)
	fmt.Println(js)
}