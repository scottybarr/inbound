package main

import (
    "fmt"
    "github.com/scottybarr/inbound/flightinfo"
)

func main() {
    flt := flightinfo.GetFlight("")

    fmt.Println(flt)
}
