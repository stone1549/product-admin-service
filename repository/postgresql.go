package repository

import (
	"context"
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/stone1549/product-admin-service/common"
	"github.com/twinj/uuid"
)

const (
	insertProductQuery = "INSERT INTO product (id, name, description, short_description, display_image, thumbnail, " +
		"price, qty_in_stock) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)"
	updateProductQuery = "UPDATE product SET name=$1, description=$2, short_description=$3, display_image=$4, " +
		"thumbnail=$5, price=$6, qty_in_stock=$7 WHERE id=$8"
)

type postgresqlProductRepository struct {
	db *sql.DB
}

// NewProduct adds a product to the repo.
func (impr *postgresqlProductRepository) NewProduct(ctx context.Context, product common.Product) (string, error) {
	if product.Name == "" {
		return "", newErrRepository("product name is required")
	}

	id := uuid.NewV4().String()

	var pricePtr *string

	if product.Price != nil {
		priceStr := product.Price.StringFixed(6)
		if priceStr != "" {
			pricePtr = &priceStr
		}

	}

	_, err := impr.db.ExecContext(ctx, insertProductQuery, id, product.Name, product.Description, product.ShortDescription,
		product.DisplayImage, product.Thumbnail, pricePtr, product.QtyInStock)

	return id, err
}

// UpdateProduct updates a product stored in the repo.
func (impr *postgresqlProductRepository) UpdateProduct(ctx context.Context, id string, product common.Product) error {
	if product.Name == "" {
		return newErrRepository("product name is required")
	}

	var pricePtr *string

	if product.Price != nil {
		priceStr := product.Price.StringFixed(6)
		if priceStr != "" {
			pricePtr = &priceStr
		}
	}

	_, err := impr.db.ExecContext(ctx, updateProductQuery, product.Name, product.Description, product.ShortDescription,
		product.DisplayImage, product.Thumbnail, pricePtr, product.QtyInStock, id)

	return err
}

func loadInitPostgresqlData(db *sql.DB, dataset string) error {
	products, err := loadInitInMemoryDataset(dataset)

	if err != nil {
		return err
	}

	txn, err := db.Begin()

	if err != nil {
		return err
	}

	for id, product := range products {
		_, err = txn.Exec(insertProductQuery, id, product.Name, product.Description, product.ShortDescription,
			product.DisplayImage, product.Thumbnail, product.Price.StringFixed(6), product.QtyInStock)

		if err != nil {
			return err
		}
	}

	return txn.Commit()
}

// MakePostgresqlProductRespository constructs a PostgreSQL backed ProductRepository from the given params.
func MakePostgresqlProductRespository(config common.Configuration, db *sql.DB) (ProductRepository, error) {
	var err error
	if config.GetInitDataSet() != "" {
		err = loadInitPostgresqlData(db, config.GetInitDataSet())
	}

	if err != nil {
		return nil, err
	}

	return &postgresqlProductRepository{db}, nil
}
