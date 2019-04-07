package main

import (
	"fmt"
)

func big(list [10]int) int {

	big := 0

	for x,y := range list {
		if len(list) == x+1{

			if big < y {
				big = y
			}
		} else {

			if y > list[x+1] && big < y {
				big = y
			}
		}

	}

	return big

}


func main() {
	a := [10] int {30,5,2,3,9,8,4,20,1,31}
	b := big(a)
	fmt.Println(b)

}