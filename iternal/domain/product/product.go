package product

import (
	"github.com/google/uuid"
	"time"
)

type Product struct {
	Id       uuid.UUID `json:"id"`
	DateTime time.Time `json:"dateTime"`
}
