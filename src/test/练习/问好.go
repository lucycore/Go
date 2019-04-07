package main

import (
	"fmt"
)

func main() {

	name := []string {}

	for {
		namea := ""
		fmt.Println("请输入名称：")
		fmt.Scanln(&namea)
		if namea == "q" {
			break
		}
		name = append(name,namea)
	}
	
	for _, y := range name {
		switch {
		case y == "wwb":
			fmt.Println("你好！", y)
		case y == "gyj":
			fmt.Println("你好！", y) 
		}
	}

}