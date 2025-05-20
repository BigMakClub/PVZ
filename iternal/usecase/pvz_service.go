package usecase

import (
	"context"
	"pvz/iternal/domain"
	m "pvz/iternal/usecase/model"
	"time"
)

type PVZService interface {
	Create(ctx context.Context, city domain.City) (*domain.PVZ, error)
	List(ctx context.Context,
		start, end time.Time,
		page, limit int) ([]m.PVZAggregate, error)
}
