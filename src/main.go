package main

import (
    "fmt"
    "time"
    "github.com/inancgumus/myhttp"
)
func main() {

    mh := myhttp.New(time.Second)
    resonse,_ := mh.Get("https://jsonip.com")
    fmt.Printf("hello, world\n")
    fmt.Println(resonse.StatusCode);
}
