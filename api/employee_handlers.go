package api

import (
	"errors"
	"net/http"

	"github.com/clintjedwards/scheduler/model"
	"github.com/clintjedwards/scheduler/utils"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
)

// TODO(clintjedwards): Add request validation

// ListEmployeesHandler returns all employees unpaginated
func (api *API) ListEmployeesHandler(w http.ResponseWriter, r *http.Request) {

	employees, err := api.storage.GetAllEmployees()
	if err != nil {
		log.Error().Err(err).Msg("failed to retrieve employees")
		sendErrResponse(w, http.StatusBadGateway, err)
		return
	}

	sendResponse(w, http.StatusOK, employees)
}

// AddEmployeeHandler adds a new employee to the scheduler service
func (api *API) AddEmployeeHandler(w http.ResponseWriter, r *http.Request) {

	pendingEmployee := model.AddEmployee{}

	err := parseJSON(r.Body, &pendingEmployee)
	if err != nil {
		log.Warn().Err(err).Msg("could not parse json request")
		sendErrResponse(w, http.StatusBadRequest, err)
		return
	}
	defer r.Body.Close()

	newEmployee := model.NewEmployee(api.config.IDLength)
	pendingEmployee.ToEmployee(newEmployee)

	err = newEmployee.IsValid()
	if err != nil {
		sendErrResponse(w, http.StatusBadRequest, err)
	}

	err = api.storage.AddEmployee(newEmployee.ID, newEmployee)
	if err != nil {
		if errors.Is(err, utils.ErrEntityExists) {
			sendErrResponse(w, http.StatusConflict, err)
			return
		}
		log.Error().Err(err).Msg("could not add employee")
		sendErrResponse(w, http.StatusBadGateway, err)
		return
	}

	log.Info().Interface("employee", newEmployee).Msg("created new employee")
	sendResponse(w, http.StatusCreated, newEmployee)
}

// GetEmployeeHandler returns a single employee by id
func (api *API) GetEmployeeHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	employee, err := api.storage.GetEmployee(vars["id"])
	if err != nil {
		if errors.Is(err, utils.ErrEntityNotFound) {
			http.Error(w, "employee not found", http.StatusNotFound)
			return
		}
		log.Error().Err(err).Msg("could not get employee")
		sendErrResponse(w, http.StatusBadGateway, err)
		return
	}

	sendResponse(w, http.StatusOK, employee)
}
