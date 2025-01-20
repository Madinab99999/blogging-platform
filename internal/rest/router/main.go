package router

import (
	"context"
	"net/http"

	"github.com/Madinab99999/blogging-platform/internal/rest/handler"
)

type Router struct {
	router  *http.ServeMux
	handler *handler.Handler
}

func New(handler *handler.Handler) *Router {
	mux := http.NewServeMux()

	return &Router{
		router:  mux,
		handler: handler,
	}
}

func (r *Router) Start(ctx context.Context) *http.ServeMux {
	return r.router
}
