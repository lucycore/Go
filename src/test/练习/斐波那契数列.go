package main 


import (
	"fmt"
)


func main() {

	var list = [] int {1,1}

	js := 100

	for js > 0 {

		js = js -1

		sl := len(list)
		sl = sl - 1

		za := list[sl]
		zb := list[sl-1]

		zc := za + zb
		list = append(list,zc)


	}
	fmt.Println(list)

}
