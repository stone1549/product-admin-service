package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/stone1549/product-admin-service/common"
	"github.com/twinj/uuid"
	"io/ioutil"
)

type inMemoryProductRepository struct {
	products map[string]*common.Product
}

// NewProduct adds a product to the repo.
func (impr *inMemoryProductRepository) NewProduct(ctx context.Context, product common.Product) (string, error) {
	if product.Name == "" {
		return "", newErrRepository("product name is required")
	}
	id := uuid.NewV4().String()

	_, ok := impr.products[id]; if ok {
		return "", newErrRepository(fmt.Sprintf("product with id %s already in repo", id))
	}

	impr.products[id] = &product

	return id, nil
}

// UpdateProduct updates a product stored in the repo.
func (impr *inMemoryProductRepository) UpdateProduct(ctx context.Context, id string, product common.Product) error {
	if product.Name == "" {
		return newErrRepository("product name is required")
	}
	_, ok := impr.products[id]; if !ok {
		return newErrRepository(fmt.Sprintf("product with id %s not found", id))
	}

	impr.products[id] = &product

	return nil
}

// MakeInMemoryRepository constructs an in memory backed ProductRepository from the given configuration.
func MakeInMemoryRepository(config common.Configuration) (ProductRepository, error) {
	var err error

	products, err := loadInitInMemoryDataset(config.GetInitDataSet())

	return &inMemoryProductRepository{products }, err
}

type storedProduct struct {
	common.Product
	Id string `json:"id"`
}

func loadInitInMemoryDataset(dataset string) (map[string]*common.Product, error) {
	if dataset == "" {
		return nil, nil
	}

	var err error
	storedProducts := make([]storedProduct, 0)

	if err != nil {
		return nil, err
	}

	jsonBytes, err := ioutil.ReadFile(dataset)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(jsonBytes, &storedProducts)

	if err != nil {
		return nil, err
	}

	products := make(map[string]*common.Product)

	for _, storedProduct := range storedProducts {
		products[storedProduct.Id] = &common.Product{Name: storedProduct.Name, DisplayImage: storedProduct.DisplayImage,
		Thumbnail: storedProduct.Thumbnail, Price: storedProduct.Price, Description: storedProduct.Description,
		ShortDescription: storedProduct.ShortDescription, QtyInStock: storedProduct.QtyInStock}
	}

	return products, err
}
