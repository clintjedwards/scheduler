package frontend

import (
	"embed"
	"io/fs"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
	"github.com/shurcooL/httpgzip"
)

// We bake frontend files directly into the binary
// embeddedAssets is an implementation of an http.filesystem
// that points to the public folder
//
//go:embed public
var embeddedAssets embed.FS

//Frontend represents an instance of the frontend application
type Frontend struct{}

//NewFrontend initializes a new UI application
func NewFrontend() *Frontend {
	return &Frontend{}
}

// HistoryModeHandler is a hack so that our frontend can use history mode.
// We answer all requests for files normally but, any other path returns the normal
// index.html file. This is so that we can handle the rendering of a 404 with javascript instead
// of a separate handler.
// https://router.vuejs.org/guide/essentials/history-mode.html
func historyModeHandler(fileServerHandler http.Handler, indexFile []byte) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		path := req.URL.Path

		// serve static files as normal
		if strings.Contains(path, ".") || path == "/" {
			fileServerHandler.ServeHTTP(w, req)
			return
		}

		// return index.html for any 404s
		_, _ = w.Write(indexFile)
	})
}

// RegisterUIRoutes registers the endpoints needed for the frontend
// with an already established http router.
func (ui *Frontend) RegisterUIRoutes(router *mux.Router) {
	fsys, err := fs.Sub(embeddedAssets, "public")
	if err != nil {
		log.Fatal().Err(err).Msg("could not get embedded filesystem")
	}

	fileServerHandler := httpgzip.FileServer(http.FS(fsys), httpgzip.FileServerOptions{IndexHTML: true})

	file, err := fsys.Open("index.html")
	if err != nil {
		log.Fatal().Err(err).Msg("could not find index.html file")
	}
	defer file.Close()

	indexContent, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal().Err(err).Msg("could not read index.html file")
	}

	router.PathPrefix("/").Handler(historyModeHandler(fileServerHandler, indexContent))
}
