package flightinfo

import (
    "encoding/json"
    "github.com/scottybarr/inbound/request"
)

type flightInfo struct {
    FaFlightID, Ident string
    prefix, suffix string
    Type string
    Origin, Destination string
    timeout string
    timestamp int
    DepartureTime, firstPositionTime, ArrivalTime int
    Longitude, Latitude float64
    lowLongitude, lowLatitude float64
    highLongitude, highLatitude float64
    Groundspeed, Altitude, Heading int
    altitudeStatus, updateType, altitudeChange string
    waypoints string
}

type FlightInfoResult struct {
    InFlightInfoResult flightInfo
}

func GetFlight(flightNum string) FlightInfoResult {
    var data []byte
    var jsontype FlightInfoResult
    data = request.GetFlightData("")
    json.Unmarshal(data, &jsontype)
    return jsontype
}
