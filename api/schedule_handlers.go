package api

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/clintjedwards/scheduler/model"
	"github.com/clintjedwards/scheduler/utils"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
)

func (api *API) generateSchedule(sch *model.Schedule) error {

	currentDate, err := time.Parse("01-02-2006", sch.Start)
	if err != nil {
		return fmt.Errorf("could not parse start date; should be in format mm-dd-yyy: %v", err)
	}

	endDate, err := time.Parse("01-02-2006", sch.End)
	if err != nil {
		return fmt.Errorf("could not parse end date; should be in format mm-dd-yyy: %v", err)
	}

	employees, err := api.getEmployees(sch.EmployeeFilter)
	if err != nil {
		return err
	}

	for {
		if currentDate.After(endDate) {
			break
		}
		err := sch.ScheduleDay(currentDate, employees)
		if err != nil {
			return err
		}
		currentDate = currentDate.AddDate(0, 0, 1)
	}

	return nil
}

// getEmployees returns a map of eligible employees for future scheduling.
// TODO(clintjedwards): This datastructure should be more advanced, something like a heap would work nicely here
// since eventually employees will have a weight. We might have to calculate weights for each
// position ahead of time and then use that here later
func (api *API) getEmployees(employeeFilter []string) (map[string]model.Employee, error) {

	eligibleEmployees := map[string]model.Employee{}

	employees, err := api.storage.GetAllEmployees()
	if err != nil {
		return nil, err
	}

	if len(employeeFilter) != 0 {
		for _, employee := range employeeFilter {
			if employees[employee].Status != model.EmployeeActive {
				continue
			}
			eligibleEmployees[employee] = *employees[employee]
		}
		return eligibleEmployees, nil
	}

	for id, employee := range employees {
		if employee.Status != model.EmployeeActive {
			continue
		}
		eligibleEmployees[id] = *employees[id]
	}

	return eligibleEmployees, nil
}

// ListSchedulesHandler returns all schedules unpaginated
func (api *API) ListSchedulesHandler(w http.ResponseWriter, r *http.Request) {

	schedules, err := api.storage.GetAllSchedules()
	if err != nil {
		log.Error().Err(err).Msg("failed to retrieve schedules")
		sendResponse(w, http.StatusBadGateway, err)
		return
	}

	sendResponse(w, http.StatusOK, schedules)
}

// GenerateScheduleHandler adds a new schedule to the scheduler service
func (api *API) GenerateScheduleHandler(w http.ResponseWriter, r *http.Request) {

	settings := model.Schedule{}

	err := parseJSON(r.Body, &settings)
	if err != nil {
		log.Warn().Err(err).Msg("could not parse json request")
		sendErrResponse(w, http.StatusBadRequest, err)
		return
	}
	defer r.Body.Close()

	newSchedule := model.NewSchedule(string(utils.GenerateRandString(api.config.IDLength)), settings)
	err = api.generateSchedule(newSchedule)
	if err != nil {
		sendErrResponse(w, http.StatusBadGateway, err)
		return
	}

	err = api.storage.AddSchedule(newSchedule.ID, newSchedule)
	if err != nil {
		if errors.Is(err, utils.ErrEntityExists) {
			sendErrResponse(w, http.StatusConflict, err)
			return
		}
		log.Error().Err(err).Msg("could not add schedule")
		sendErrResponse(w, http.StatusBadGateway, err)
		return
	}

	log.Info().Interface("schedule", newSchedule).Msg("created new schedule")
	sendResponse(w, http.StatusCreated, newSchedule)
}

// GetScheduleHandler returns a single schedule by id
func (api *API) GetScheduleHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	schedule, err := api.storage.GetSchedule(vars["id"])
	if err != nil {
		if errors.Is(err, utils.ErrEntityNotFound) {
			sendErrResponse(w, http.StatusNotFound, err)
			return
		}
		log.Error().Err(err).Msg("could not get schedule")
		sendErrResponse(w, http.StatusBadGateway, err)
		return
	}

	sendResponse(w, http.StatusOK, schedule)
}
