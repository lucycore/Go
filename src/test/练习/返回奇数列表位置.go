package main 


import (
	"fmt"
)


func main() {

	var list = [10] int {1,2,3,4,5,6,7,8,9,10}

	for x,y := range list {
		if x % 2 == 0 {
			fmt.Println(y)
		}
	}


}
