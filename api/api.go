package api

import (
	"net/http"

	"github.com/clintjedwards/scheduler/config"
	"github.com/clintjedwards/scheduler/storage"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// API represents the grpc backend service
type API struct {
	config  *config.Config
	storage storage.Engine
}

// NewAPI inits a new api service
func NewAPI(config *config.Config, storage storage.Engine) *API {
	return &API{
		config:  config,
		storage: storage,
	}
}

// RegisterEmployeeRoutes registers /employees with a given router
func (api *API) RegisterEmployeeRoutes(router *mux.Router) {
	router.Handle("/api/employees", handlers.MethodHandler{
		"GET":  http.HandlerFunc(api.ListEmployeesHandler),
		"POST": http.HandlerFunc(nil),
	})

	router.Handle("/api/employees/{id}", handlers.MethodHandler{
		"GET":    http.HandlerFunc(nil),
		"DELETE": http.HandlerFunc(nil),
	})
}

// RegisterPositionRoutes registers /positions with a given router
func (api *API) RegisterPositionRoutes(router *mux.Router) {
	router.Handle("/api/positions", handlers.MethodHandler{
		"GET":  http.HandlerFunc(nil),
		"POST": http.HandlerFunc(nil),
	})

	router.Handle("/api/positions/{id}", handlers.MethodHandler{
		"GET":    http.HandlerFunc(nil),
		"DELETE": http.HandlerFunc(nil),
	})
}

// RegisterScheduleRoutes registers /schedules with a given router
func (api *API) RegisterScheduleRoutes(router *mux.Router) {
	router.Handle("/api/schedules", handlers.MethodHandler{
		"GET":  http.HandlerFunc(nil),
		"POST": http.HandlerFunc(nil),
	})

	router.Handle("/api/schedules/{id}", handlers.MethodHandler{
		"GET":    http.HandlerFunc(nil),
		"DELETE": http.HandlerFunc(nil),
	})
}

// RegisterSystemRoutes registers /system/info with a given router
func (api *API) RegisterSystemRoutes(router *mux.Router) {
	router.Handle("/api/system/info", handlers.MethodHandler{
		"GET": http.HandlerFunc(nil),
	})
}
