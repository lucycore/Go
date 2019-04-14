package download

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

