package repository_test

import (
	"github.com/stone1549/product-admin-service/repository"
	"testing"
)

func makeNewImRepo(t *testing.T) repository.ProductRepository {
	repo, err := repository.MakeInMemoryRepository(inMemorySmall)

	ok(t, err)
	return repo
}

