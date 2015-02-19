package request

import (
	"io/ioutil"
	"net/http"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// GetFlightData makes a request to the flight aware API
func GetFlightData(flightNum string) []byte {
	url := os.Getenv("FLIGHTAWARE_API_URL")
	url += "/InFlightInfo?ident=" + flightNum
	return makeRequest(url)
}

func makeRequest(url string) []byte {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	check(err)

	setAuth(req)
	resp, err := client.Do(req)
	check(err)

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	check(err)

	return body
}

func setAuth(req *http.Request) {
	username := os.Getenv("FLIGHTAWARE_API_USERNAME")
	apiKey := os.Getenv("FLIGHTAWARE_API_KEY")
	req.SetBasicAuth(username, apiKey)
}
