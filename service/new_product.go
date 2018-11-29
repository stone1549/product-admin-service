package service

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/go-chi/render"
	"github.com/stone1549/product-admin-service/common"
	"github.com/stone1549/product-admin-service/repository"
	"net/http"
)

type newProductResponse struct {
	Id               string     `json:"id"`
}

func (plr newProductResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

// NewProductMiddleware middleware to add a new product from the request parameters to the repo
func NewProductMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)

		var product common.Product
		err := decoder.Decode(&product)
		if err != nil {
			render.Render(w, r, errInvalidRequest(err))
			return
		}

		if product.Name == "" {
			render.Render(w, r, errInvalidRequest(errors.New("product name is required")))
			return
		}

		productRepo, ok := r.Context().Value("repo").(repository.ProductRepository)

		if !ok {
			render.Render(w, r, errRepository(errors.New("ProductRepository not found in context")))
			return
		}

		id, err := productRepo.NewProduct(r.Context(), product)

		if err != nil {
			render.Render(w, r, errRepository(err))
			return
		} else if id == "" {
			render.Render(w, r, errUnknown(errors.New("unable to determine whether or not product was inserted")))
			return
		}

		ctx := context.WithValue(r.Context(), "id", id)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// NewProduct renders the newly inserted products id.
func NewProduct(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id, ok := ctx.Value("id").(string)

	if !ok {
		render.Render(w, r, errUnknown(errors.New("unable to determine whether or not product was inserted")))
		return
	}

	if err := render.Render(w, r, newProductResponse{id}); err != nil {
		render.Render(w, r, errUnknown(err))
		return
	}
}
