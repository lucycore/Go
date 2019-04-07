package main

import (
	"fmt"
)

func main() {
	a := [9]int {1,2,3,4,5,6,7,8,9} 
	b := [9]int {1,2,3,4,5,6,7,8,9} 

	for _,x := range a {
		for _,y := range b {

			fmt.Println(x,"*",y,"=",x*y)

		}
	}

}