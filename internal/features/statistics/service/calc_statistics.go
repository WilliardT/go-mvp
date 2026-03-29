package statistics_service

import "github.com/WilliardT/go-mvp/internal/core/domain"


func calcProductPriceRanges(prices []float64) domain.ProductPriceRanges {
	var ranges domain.ProductPriceRanges

	for _, price := range prices {
		switch {
		case price < 100:
			ranges.Cheap++
		case price <= 500:
			ranges.Medium++
		default:
			ranges.Expensive++
		}
	}

	return ranges
}
