package frontend

import (
	"github.com/gorilla/mux"
	"github.com/shurcooL/httpgzip"
)

//Frontend represents an instance of the frontend application
type Frontend struct{}

//NewFrontend initializes a new UI application
func NewFrontend() *Frontend {
	return &Frontend{}
}

// RegisterUIRoutes registers the endpoints needed for the frontend
// with an already established http router.
func (ui *Frontend) RegisterUIRoutes(router *mux.Router) {

	// We bake frontend files directly into the binary
	// assets is an implementation of an http.filesystem created by
	// github.com/shurcooL/vfsgen that points to the "public" folder
	fileServerHandler := httpgzip.FileServer(assets, httpgzip.FileServerOptions{IndexHTML: true})

	router.PathPrefix("/").Handler(fileServerHandler)
}
