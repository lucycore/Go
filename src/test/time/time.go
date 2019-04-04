package main

import (
	"time"
	"fmt"
)

func main() {
	a := time.Now().Format("2006-01-02-15-04")
	fmt.Println(a)
}