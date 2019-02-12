package v1

import (
	"net/http"

	"github.com/andreylm/process-manager/models/task"
	"github.com/andreylm/process-manager/router/routes"
	V1ProcessRoutes "github.com/andreylm/process-manager/router/routes/v1/process"
	V1TaskRoutes "github.com/andreylm/process-manager/router/routes/v1/task"
)

// Middleware - sub route middleware
func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Actions before running routes
		next.ServeHTTP(w, r)
	})
}

// GetRoutes - get v1 routes
func GetRoutes(collection *task.Collection) (SubRoute map[string]routes.SubRoutePackage) {
	SubRoute = map[string]routes.SubRoutePackage{
		"/v1": routes.SubRoutePackage{
			Routes: routes.Routes{
				// Task's routes
				routes.Route{
					Name:        "V1TaskListRoute",
					Method:      "GET",
					Pattern:     "/tasks",
					HandlerFunc: V1TaskRoutes.List(collection),
				},
				routes.Route{
					Name:        "V1TaskCreateRoute",
					Method:      "POST",
					Pattern:     "/task",
					HandlerFunc: V1TaskRoutes.Create(collection),
				},
				routes.Route{
					Name:        "V1TaskExistRoute",
					Method:      "GET",
					Pattern:     "/task/exist/{task}",
					HandlerFunc: V1TaskRoutes.Exist(collection),
				},
				routes.Route{
					Name:        "V1TaskInfoRoute",
					Method:      "GET",
					Pattern:     "/task/{task}",
					HandlerFunc: V1TaskRoutes.Info(collection),
				},
				routes.Route{
					Name:        "V1TaskDeleteRoute",
					Method:      "DELETE",
					Pattern:     "/task/{task}",
					HandlerFunc: V1TaskRoutes.Delete(collection),
				},
				routes.Route{
					Name:        "V1TaskDeleteAllRoute",
					Method:      "DELETE",
					Pattern:     "/tasks",
					HandlerFunc: V1TaskRoutes.DeleteAll(collection),
				},
				// Process routes
				routes.Route{
					Name:        "V1ProcessCreateRoute",
					Method:      "POST",
					Pattern:     "/task/{name}",
					HandlerFunc: V1ProcessRoutes.Create(collection),
				},
			},
			Middleware: Middleware,
		},
	}
	return SubRoute
}
