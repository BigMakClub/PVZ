package reception

import (
	"errors"
	"github.com/google/uuid"
	"pvz/iternal/domain/pvz"
	"time"
)

type Status string

const (
	StatusInProgress Status = "IN_PROGRESS"
	StatusClosed     Status = "CLOSED"
)

var (
	ErrReceptionClosed = errors.New("Reception is already closed")
	ErrReceptionOpen   = errors.New("Reception is already open")
)

type Reception struct {
	Id       uuid.UUID
	DateTime time.Time
	PvzId    uuid.UUID
	Status   Status
}

func NewReception(pvz pvz.PVZ) *Reception {
	return &Reception{
		Id:       uuid.New(),
		DateTime: time.Now(),
		PvzId:    pvz.Id,
		Status:   StatusInProgress,
	}
}

func (r *Reception) Close() error {
	if r.Status == StatusClosed {
		return ErrReceptionClosed
	}
	r.Status = StatusClosed
	return nil
}

func (r *Reception) Open() error {
	if r.Status == StatusInProgress {
		return ErrReceptionOpen
	}
	r.Status = StatusInProgress
	return nil
}
