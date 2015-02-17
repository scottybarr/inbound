package inboundflight

type InboundFlight struct {
    FAFlightId, Ident, Origin ,Destination, AircraftType string
    ActualArrival, EstimatedArrival int
}
