package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/scottybarr/inbound/flightinfo"
)

// check is responsible for error handling
func check(e error) {
	if e != nil {
		panic(e)
	}
}

// getFlight is the handler for the route:
// "/flight/BAW123"
// Obtains flight information from flightInfo API
func getFlight(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	flightID := ps.ByName("flightId")
	flt := flightinfo.GetFlight(flightID)

	js, err := json.Marshal(flt)
	check(err)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", js)
}

// ping is a healthcheck endpoint on the route:
// "/ping"
func ping(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	t := time.Now().UTC()
	fmt.Fprint(w, t.Format(time.RFC1123Z))
}

// main initialises our server
func main() {
	router := httprouter.New()

	router.GET("/ping", ping)
	router.GET("/flight/:flightId", getFlight)

	http.ListenAndServe(":8080", router)
}
