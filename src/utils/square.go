package utils

import (
	e "math-coordenadas/src/entities"
)

func IsAt(alvo e.ICoordenada, lat, lon e.ICoordenada) bool {
	if alvo.Latitude() >= lat.MinValue() && alvo.Latitude() <= lat.MaxValue() &&
		alvo.Longitude() >= lon.MinValue() && alvo.Longitude() <= lon.MaxValue() {
		return true
	} else {
		return false
	}
}

func VerifyLocations() {}
