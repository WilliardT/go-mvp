package statistics_service

import (
	"context"
	"fmt"
	"time"

	"github.com/WilliardT/go-mvp/internal/core/domain"
	core_errors "github.com/WilliardT/go-mvp/internal/core/errors"
)

func (s *StatisticsService) GetProductsStatistics(
	ctx context.Context,
	authorUserID *int,
	createdFrom *time.Time,
	createdTo *time.Time,
) (domain.Statistics, error) {
	if authorUserID != nil && *authorUserID <= 0 {
		return domain.Statistics{}, fmt.Errorf(
			"author user id must be positive: %w",
			core_errors.ErrInvalidArgument,
		)
	}

	if createdFrom != nil && createdTo != nil && createdFrom.After(*createdTo) {
		return domain.Statistics{}, fmt.Errorf(
			"'created_from' must be less than or equal to 'created_to': %w",
			core_errors.ErrInvalidArgument,
		)
	}

	statistics, err := s.statisticsRepository.GetProductsStatistics(
		ctx,
		authorUserID,
		createdFrom,
		createdTo,
	)
	if err != nil {
		return domain.Statistics{}, fmt.Errorf(
			"get products statistics from repository: %w",
			err,
		)
	}

	prices, err := s.statisticsRepository.GetProductsPrices(
		ctx,
		authorUserID,
		createdFrom,
		createdTo,
	)
	if err != nil {
		return domain.Statistics{}, fmt.Errorf(
			"get product prices from repository: %w",
			err,
		)
	}

	statistics.ProductPriceRanges = calcProductPriceRanges(prices)

	return statistics, nil
}
