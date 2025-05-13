package reception

import (
	"github.com/google/uuid"
	"time"
)

type Reception struct {
	Id       uuid.UUID `json:"id"`
	DateTime time.Time `json:"dateTime"`
	PvzId    uuid.UUID `json:"pvzId"`
	Status   string    `json:"status"`
}
