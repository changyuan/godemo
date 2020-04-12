package main

import (
    "fmt"
    "rsc.io/quote"
)

func main() {
    fmt.Println(quote.Hello())

    var a = App{}
    a.Initialize()
    a.Run(":8000")

}
