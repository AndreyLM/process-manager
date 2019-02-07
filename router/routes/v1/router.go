package v1

import (
	"net/http"

	"github.com/andreylm/process-manager/models/process"
	"github.com/andreylm/process-manager/router/routes"
	V1ProcessRoutes "github.com/andreylm/process-manager/router/routes/v1/process"
)

// Middleware - sub route middleware
func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Actions before running routes
		next.ServeHTTP(w, r)
	})
}

// GetRoutes - get v1 routes
func GetRoutes(collection *process.Collection) (SubRoute map[string]routes.SubRoutePackage) {
	SubRoute = map[string]routes.SubRoutePackage{
		"/v1": routes.SubRoutePackage{
			Routes: routes.Routes{
				routes.Route{
					Name:        "V1ProcessCreateRoute",
					Method:      "POST",
					Pattern:     "/create",
					HandlerFunc: V1ProcessRoutes.Create(collection),
				},
			},
			Middleware: Middleware,
		},
	}
	return SubRoute
}
