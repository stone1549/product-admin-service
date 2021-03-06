package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"github.com/stone1549/product-admin-service/common"
	"github.com/stone1549/product-admin-service/repository"
	"github.com/stone1549/product-admin-service/service"
	"net/http"
)

func main() {
	flag.Parse()

	config, err := common.GetConfiguration()

	if err != nil {
		panic(fmt.Sprintf("Unable to load configuration: %s", err.Error()))
	}

	repo, err := repository.NewProductRepository(config)

	if err != nil {
		panic(fmt.Sprintf("Unable to configure repository: %s", err.Error()))
	}

	repoMiddleWare := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(r.Context(), "repo", repo)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)
	r.Use(render.SetContentType(render.ContentTypeJSON))
	r.Use(repoMiddleWare)

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	r.Use(middleware.Timeout(config.GetTimeout()))

	r.Route("/product", func(r chi.Router) {
		r.With(service.NewProductMiddleware).Post("/", service.NewProduct)
		r.Route("/{productId}", func(r chi.Router) {
			r.With(service.UpdateProductMiddleware).Patch("/", service.UpdateProduct)
		})
	})

	http.ListenAndServe(":3333", r)
}

