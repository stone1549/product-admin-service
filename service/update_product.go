package service

import (
	"encoding/json"
	"errors"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/stone1549/product-admin-service/common"
	"github.com/stone1549/product-admin-service/repository"
	"net/http"
)

// UpdateProductMiddleware middleware to update a product from the request parameters
func UpdateProductMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "productId")

		if id == "" {
			render.Render(w, r, errInvalidRequest(errors.New("product id is required")))
			return
		}

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

		err = productRepo.UpdateProduct(r.Context(), id, product)

		if err != nil {
			render.Render(w, r, errRepository(err))
			return
		}

		next.ServeHTTP(w, r)
	})
}

// UpdateProduct renders the response to the product update request.
func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	render.Status(r, 200)
}
