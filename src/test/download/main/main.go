package main

import (
    "io"
    "net/http"
    "os"
)


func download(url, lj string) {
    res, err := http.Get(url)
    if err != nil {
        panic(err)
    }
    f, err := os.Create(lj)
    if err != nil {
        panic(err)
    }
    io.Copy(f, res.Body)
}


func main() {
	download("http://www.lucycore.top/v2ray/v2rayMacX.zip", "a.zip")
}