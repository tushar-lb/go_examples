package server

import (
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)

	}
	return router
}

func init() {
	// Initialise file counter information
	fileExt := make(map[string]int)
	filePath := make([]string, 10)
	FileInfoCounter = FileStatsInfo{
		FileCount:       0,
		MaxFileSize:     0,
		AvgFileSize:     0,
		FileExtension:   fileExt,
		RecentFilePaths: filePath,
	}
}
