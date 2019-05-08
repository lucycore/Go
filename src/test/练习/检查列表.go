package main 


import (
	"fmt"
)


func main() {

	var list = [5] string {"陈世霖","陈锦泽","段子杰","张宗旭","李嘉熙"}

	for _,y := range list {
		if y == "张宗旭" {
			fmt.Println("ok")
		} else {
			fmt.Println("no")
		}
	}


}
