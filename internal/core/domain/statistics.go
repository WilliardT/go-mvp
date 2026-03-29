package domain

type Statistics struct {
	ProductsCount      int
	ProductPriceMedian *float64
}

func NewStatistics(
	productsCount int,
	productPriceMedian *float64,
) Statistics {
	return Statistics{
		ProductsCount:      productsCount,
		ProductPriceMedian: productPriceMedian,
	}
}
