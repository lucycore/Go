package main

import (
    //"os/exec"
    "io/ioutil"
    "fmt"
)

/*
func Dlef(lj) {
    dir := lj
    exec.Command("cmd", "/C", "rd", "/S", "/Q", dir).Start()
}
*/

func main() {
    files, _ := ioutil.ReadDir("C:/Users/lucycore/Desktop/归档")
    for _, f := range files {
        fmt.Println(f.Name())
    }
}
