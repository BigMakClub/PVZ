package domain

import (
	"errors"
	"github.com/google/uuid"
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
	PvzId    uuid.UUID
	Status   Status
	DateTime time.Time
}

func NewReception(pvz PVZ) *Reception {
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
