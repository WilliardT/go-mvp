package products_postgres_repository

import core_postgres_pool "github.com/WilliardT/go-mvp/internal/core/repository/postgres/pool"

type ProductsRepository struct {
	pool core_postgres_pool.Pool
}

func NewProductsRepository(
	pool core_postgres_pool.Pool,
) *ProductsRepository {
	return &ProductsRepository{
		pool: pool,
	}
}
