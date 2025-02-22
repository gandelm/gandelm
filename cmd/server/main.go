package server

import (
	"log"
	"net/http"

	"github.com/gandelm/gandelm/generated/protocol/catalog/v1/catalogv1connect"
	"github.com/gandelm/gandelm/internal/container"
	"github.com/gandelm/gandelm/internal/service/catalog"
	"github.com/rs/cors"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func Start(container container.Containerer) error {
	log.Println("Starting server on localhost:8080")

	mux := http.NewServeMux()
	mux.Handle(catalogv1connect.NewCatalogServiceHandler(catalog.NewCatalogService(container)))

	corsHandler := cors.AllowAll().Handler(h2c.NewHandler(mux, &http2.Server{}))

	return http.ListenAndServe(
		"localhost:8080",
		corsHandler,
	)
}
