package flightinfo

import (
	"encoding/json"

	"github.com/scottybarr/inbound/request"
)

// FlightInfoResult is the exported struct we use
// that contains flight information.
type FlightInfoResult struct {
	InFlightInfoResult flightInfo `json:"inFlightInfoResult"`
}

// flightInfo struct is responsible for providing
// details for a particular flight ID
type flightInfo struct {
	FaFlightID        string `json:"faFlightID"`
	Ident             string `json:"ident"`
	prefix            string
	suffix            string
	Type              string `json:"type"`
	Origin            string `json:"origin"`
	Destination       string `json:"destination"`
	timeout           string
	timestamp         int
	DepartureTime     int
	firstPositionTime int
	ArrivalTime       int     `json:"arrivalTime"`
	Longitude         float64 `json:"longitude"`
	Latitude          float64 `json:"latitude"`
	lowLongitude      float64
	lowLatitude       float64
	highLongitude     float64
	highLatitude      float64
	Groundspeed       int `json:"groundspeed"`
	Altitude          int `json:"altitude"`
	Heading           int `json:"heading"`
	altitudeStatus    string
	updateType        string
	altitudeChange    string
	waypoints         string
}

// GetFlight Makes a request to the FlightAware API for a
// particular flight number.
// Returns a FlightInfoResult
func GetFlight(flightNum string) FlightInfoResult {
	var data []byte
	var jsontype FlightInfoResult
	data = request.GetFlightData(flightNum)
	json.Unmarshal(data, &jsontype)
	return jsontype
}
