package main

import (
	"fmt"
	"strconv"

)

type food struct {
	name string
	number int
	jg int
}

type food_li struct {
	name string
	jg int
}

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


func food_list() []food{

	input := []food_li {
		food_li {"宫保鸡丁",30},
		food_li {"锅包肉",50},
	}

	for x , y := range input{
		fmt.Println("编号:", x+1 , "菜品名称：",y.name," 菜品价格：",y.jg)
	}

	fmt.Println("请问您要点什么？（在下方输入菜品编号即可，录入完成输入‘ok’结束）")

	var food_list_li []food 

	for {
		var user_in string
		var num string
		fmt.Println("编号：")
		fmt.Scanln(&user_in)

		if user_in == "ok"{
			break
		} else {
			fmt.Println("数量：")
			fmt.Scanln(&num)
		}


		ys,_ := strconv.Atoi(user_in)
		cpmc := input[ys-1].name
		cpjg := input[ys-1].jg
		b,_ := strconv.Atoi(num)
		jgz := cpjg * b

		a := food {cpmc,b,jgz}
		food_list_lizz := &food_list_li
		*food_list_lizz = append(*food_list_lizz,a)

		fmt.Println("菜品名称：",cpmc," 数量：",b ," 总价格：", jgz)
	}
	return food_list_li

}



func main() {
	name, nu :=welcome()
	fmt.Println("名称：",name)
	fmt.Println("人数：",nu)

	a := food_list()

}
