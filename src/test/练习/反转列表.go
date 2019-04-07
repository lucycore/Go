package main

import (
	"fmt"
)

func fzlist(list []string) []string {
	sy := len(list) - 1
	js := 0
	var newlist []string

	for ; len(list)-1 >= js ; js += 1 {

		a := len(list)-1-js

		newlist = append(newlist,list[a])
	}
	_ = sy
	return newlist
}


func main() {
	a := []string {"h","e","l","l","o"}
	b := fzlist(a)
	fmt.Println(b)
}