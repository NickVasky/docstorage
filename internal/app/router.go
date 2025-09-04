package app

import (
	"net/http"

	"github.com/NickVasky/docstorage/internal/codegen/apicodegen"
	"github.com/gorilla/mux"
)

func (s *App) newRouter(routers ...apicodegen.Router) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	for _, api := range routers {
		for name, route := range api.Routes() {
			var handler http.Handler = route.HandlerFunc

			//handler = s.SecurityController.AuthMiddleware(handler, name)
			handler = apicodegen.Logger(handler, name)

			router.
				Methods(route.Method).
				Path(route.Pattern).
				Name(name).
				Handler(handler)
		}
	}

	return router
}
