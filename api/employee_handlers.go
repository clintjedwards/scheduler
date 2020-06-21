package api

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/clintjedwards/scheduler/models"
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
		sendResponse(w, http.StatusBadGateway, nil, errors.New("failed to retrieve employees"))
		return
	}

	sendResponse(w, http.StatusOK, employees, nil)
}

// AddEmployeeHandler adds a new employee to the scheduler service
func (api *API) AddEmployeeHandler(w http.ResponseWriter, r *http.Request) {

	newEmployee := models.Employee{}

	err := parseJSON(r.Body, &newEmployee)
	if err != nil {
		log.Warn().Err(err).Msg("could not parse json request")
		sendResponse(w, http.StatusBadRequest, nil, fmt.Errorf("could not parse json request: %v", err))
		return
	}
	defer r.Body.Close()

	newEmployee.ID = string(utils.GenerateRandString(api.config.IDLength))
	newEmployee.Created = time.Now().Unix()
	newEmployee.Modified = time.Now().Unix()

	err = api.storage.AddEmployee(newEmployee.ID, &newEmployee)
	if err != nil {
		if errors.Is(err, utils.ErrEntityExists) {
			sendResponse(w, http.StatusConflict, nil, utils.ErrEntityExists)
			return
		}
		log.Error().Err(err).Msg("could not add employee")
		sendResponse(w, http.StatusBadGateway, nil, errors.New("could not add employee"))
		return
	}

	log.Info().Interface("employee", newEmployee).Msg("created new employee")
	sendResponse(w, http.StatusOK, newEmployee, nil)
}

// GetEmployeeHandler returns a single employee by id
func (api *API) GetEmployeeHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	employee, err := api.storage.GetEmployee(vars["id"])
	if err != nil {
		if errors.Is(err, utils.ErrEntityNotFound) {
			sendResponse(w, http.StatusNotFound, nil, utils.ErrEntityNotFound)
			return
		}
		log.Error().Err(err).Msg("could not get employee")
		sendResponse(w, http.StatusBadGateway, nil, utils.ErrEntityNotFound)
		return
	}

	sendResponse(w, http.StatusOK, employee, nil)
}
