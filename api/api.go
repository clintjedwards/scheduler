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

func unimplementedhandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, http.StatusText(http.StatusNotImplemented), http.StatusNotImplemented)
}

// RegisterEmployeeRoutes registers /employees with a given router
func (api *API) RegisterEmployeeRoutes(router *mux.Router) {
	router.Handle("/api/employees", handlers.MethodHandler{
		"GET":  http.HandlerFunc(api.ListEmployeesHandler),
		"POST": http.HandlerFunc(api.AddEmployeeHandler),
	})

	router.Handle("/api/employees/{id}", handlers.MethodHandler{
		"GET":    http.HandlerFunc(api.GetEmployeeHandler),
		"PUT":    http.HandlerFunc(api.UpdateEmployeeHandler),
		"DELETE": http.HandlerFunc(api.DeleteEmployeeHandler),
	})
}

// RegisterPositionRoutes registers /positions with a given router
func (api *API) RegisterPositionRoutes(router *mux.Router) {
	router.Handle("/api/positions", handlers.MethodHandler{
		"GET":  http.HandlerFunc(api.ListPositionsHandler),
		"POST": http.HandlerFunc(api.AddPositionHandler),
	})

	router.Handle("/api/positions/{id}", handlers.MethodHandler{
		"GET":    http.HandlerFunc(api.GetPositionHandler),
		"DELETE": http.HandlerFunc(unimplementedhandler),
	})
}

// RegisterScheduleRoutes registers /schedules with a given router
func (api *API) RegisterScheduleRoutes(router *mux.Router) {
	router.Handle("/api/schedules", handlers.MethodHandler{
		"GET":  http.HandlerFunc(api.ListSchedulesHandler),
		"POST": http.HandlerFunc(api.GenerateScheduleHandler),
	})

	router.Handle("/api/schedules/{id}", handlers.MethodHandler{
		"GET":    http.HandlerFunc(api.GetScheduleHandler),
		"DELETE": http.HandlerFunc(unimplementedhandler),
	})
}

// RegisterSystemRoutes registers /system/info with a given router
func (api *API) RegisterSystemRoutes(router *mux.Router) {
	router.Handle("/api/system/info", handlers.MethodHandler{
		"GET": http.HandlerFunc(api.GetSystemInfoHandler),
	})
}
