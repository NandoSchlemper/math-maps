package utils

import (
	e "math-coordenadas/src/entities"
)

func IsAt(alvo e.ICoordenada, lat, lon e.ICoordenada) bool {
	return alvo.Latitude() >= lat.MinValue() && alvo.Latitude() <= lat.MaxValue() &&
		alvo.Longitude() >= lon.MinValue() && alvo.Longitude() <= lon.MaxValue()
}

func VerifyLocations() {}
