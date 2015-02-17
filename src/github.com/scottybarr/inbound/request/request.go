package request

import (
    "io/ioutil"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func GetFlightData(flightNum string) []byte {
    var fName string
    fName = "github.com/scottybarr/inbound/samples/inFlightInfo.json"
    data, err := ioutil.ReadFile(fName)
    check(err)
    return data
}
