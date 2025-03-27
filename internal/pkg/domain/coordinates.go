package domain

type Coordinates struct {
	latitude  float64 // 위도, ex> 37.48184253930266
	longitude float64 // 경도, ex> 127.01572695888012
}

func NewCoordinates(latitude float64, longitude float64) *Coordinates {
	return &Coordinates{
		latitude:  latitude,
		longitude: longitude,
	}
}
