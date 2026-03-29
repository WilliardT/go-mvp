package statistics_service

import (
	"context"
	"time"

	"github.com/WilliardT/go-mvp/internal/core/domain"
)

type StatisticsService struct {
	statisticsRepository StatisticsRepository
}

type StatisticsRepository interface {
	GetProductsStatistics(
		ctx context.Context,
		authorUserID *int,
		createdFrom *time.Time,
		createdTo *time.Time,
	) (domain.Statistics, error)

	GetProductsPrices(
		ctx context.Context,
		authorUserID *int,
		createdFrom *time.Time,
		createdTo *time.Time,
	) ([]float64, error)
}

func NewStatisticsService(
	statisticsRepository StatisticsRepository,
) *StatisticsService {
	return &StatisticsService{
		statisticsRepository: statisticsRepository,
	}
}
