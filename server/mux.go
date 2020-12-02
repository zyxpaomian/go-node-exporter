package server

import (
	"github.com/gorilla/mux"
)

type WWWMux struct {
	r *mux.Router
}

func New() *WWWMux {
	return &WWWMux{r:mux.NewRouter()}
}

func (m *WWWMux) GetRouter() *mux.Router {
	return m.r
}