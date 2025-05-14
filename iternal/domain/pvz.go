package domain

import (
	"errors"
	"github.com/google/uuid"
	"time"
)

type City string

var ErrCityNotAllowed = errors.New("city is not allowed")

const (
	CityMoscow City = "Москва"
	CityKazan  City = "Казань"
	CitySPB    City = "Санкт-Петербург"
)

func (c City) IsAllowed() bool {
	if c == CityMoscow || c == CityKazan || c == CitySPB {
		return true
	}
	return false
}

type PVZ struct {
	Id               uuid.UUID
	RegistrationDate time.Time
	City             City
}

func NewPVZ(city City) (*PVZ, error) {
	if !city.IsAllowed() {
		return nil, ErrCityNotAllowed
	}
	return &PVZ{
		Id:               uuid.New(),
		RegistrationDate: time.Now(),
		City:             city,
	}, nil
}
