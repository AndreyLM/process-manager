package router

import (
	"net/http"

	"github.com/andreylm/process-manager/models/process"
	"github.com/andreylm/process-manager/router/routes"
	V1Routes "github.com/andreylm/process-manager/router/routes/v1"
	"github.com/gorilla/mux"
)

const (
	staticDir = "/static/"
)

// RouteHandler - the handler for go api routes
type RouteHandler struct {
	Router *mux.Router
}

// AttachSubRouterWithMiddleware - allows  you to attach subroutes to router
func (r *RouteHandler) AttachSubRouterWithMiddleware(
	path string,
	subroutes routes.Routes,
	Middleware mux.MiddlewareFunc,
) (SubRouter *mux.Router) {
	SubRouter = r.Router.PathPrefix(path).Subrouter()
	SubRouter.Use(Middleware)

	for _, route := range subroutes {
		SubRouter.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	return
}

// NewRouter - creates new route handler
func NewRouter(collection *process.Collection) *RouteHandler {
	var router RouteHandler

	router.Router = mux.NewRouter().StrictSlash(true)

	router.Router.PathPrefix(staticDir).
		Handler(http.StripPrefix(staticDir, http.FileServer(http.Dir("."+staticDir))))

	v1routes := V1Routes.GetRoutes(collection)
	for name, pack := range v1routes {
		router.AttachSubRouterWithMiddleware(name, pack.Routes, pack.Middleware)
	}

	return &router
}
