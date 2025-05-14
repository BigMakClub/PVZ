package domain

import (
	"errors"
	"github.com/google/uuid"
	"time"
)

type ProductType string

var ErrProductIsNotValid = errors.New("Product is not valid")

const (
	Electronica ProductType = "электроника"
	Clothes     ProductType = "одежда"
	Shoes       ProductType = "обувь"
)

func (pt ProductType) IsValid() bool {
	if pt == Electronica || pt == Clothes || pt == Shoes {
		return true
	} else {
		return false
	}
}

type Product struct {
	Id          uuid.UUID
	DateTime    time.Time
	Type        ProductType
	ReceptionID uuid.UUID
}

func NewProduct(productType ProductType, reception Reception) (*Product, error) {
	if !productType.IsValid() {
		return nil, ErrProductIsNotValid
	}
	return &Product{
		Id:          uuid.New(),
		DateTime:    time.Now(),
		Type:        productType,
		ReceptionID: reception.Id,
	}, nil
}
