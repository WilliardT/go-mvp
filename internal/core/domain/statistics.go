package domain

type ProductPriceRanges struct {
	Cheap     int
	Medium    int
	Expensive int
}

type Statistics struct {
	ProductsCount      int
	ProductPriceMedian *float64
	ProductPriceRanges ProductPriceRanges
}

func NewStatistics(
	productsCount int,
	productPriceMedian *float64,
	productPriceRanges ProductPriceRanges,
) Statistics {
	return Statistics{
		ProductsCount:      productsCount,
		ProductPriceMedian: productPriceMedian,
		ProductPriceRanges: productPriceRanges,
	}
}
