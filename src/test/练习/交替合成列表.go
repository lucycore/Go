package main 


import (
	"fmt"
)


func main() {

	var list = [5] string {"1","2","3","4","5"}
	var list2 = [5] string {"段子杰","张宗旭","薛晓阳","汪文博","宫逸君"}
	var newlist = [] string {}

	for x,y := range list {
		newlist = append(newlist,y)
		newlist = append(newlist,list2[x])
	}
	fmt.Println(newlist)


}
