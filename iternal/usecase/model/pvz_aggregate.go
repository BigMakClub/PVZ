package model

import "pvz/iternal/domain"

type ReceptionAggregate struct {
	Reception *domain.Reception
	Products  []domain.Product
}

type PVZAggregate struct {
	PVZ        *domain.PVZ
	Receptions []ReceptionAggregate
}
