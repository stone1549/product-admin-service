package repository

import (
	"context"
	"database/sql"
	"github.com/stone1549/product-admin-service/common"
)


// ProductRepository represents a data source through which products can be managed.
type ProductRepository interface {
	// NewProduct adds a product to the repo.
	NewProduct(ctx context.Context, product common.Product) (string, error)
	// UpdateProduct updates a product stored in the repo.
	UpdateProduct(ctx context.Context, id string, product common.Product) error
}

// NewProductRepository constructs a ProductRepository from the given configuration.
func NewProductRepository(config common.Configuration) (ProductRepository, error) {
	var err error
	var repo ProductRepository
	var db *sql.DB
	switch config.GetRepoType() {
	case common.InMemoryRepo:
		repo, err = MakeInMemoryRepository(config)
	case common.PostgreSqlRepo:
		db, err = sql.Open("postgres", config.GetPgUrl())

		if err != nil {
			return nil, err
		}
		repo, err = MakePostgresqlProductRespository(config, db)
	default:
		err = newErrRepository("repository type unimplemented")
	}

	return repo, err
}
