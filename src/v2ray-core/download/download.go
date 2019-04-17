package download

import (
    "io"
    "net/http"
    "os"
    "io/ioutil"
    "bytes"
)

func Download(url, lj string) {

    resp, _ := http.Get(url)
    body, _ := ioutil.ReadAll(resp.Body)
    out, _ := os.Create(lj)
    io.Copy(out, bytes.NewReader(body))


}