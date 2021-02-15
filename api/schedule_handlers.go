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
	dates, err := getDates(sch.Start, sch.End)
	if err != nil {
		return err
	}

	employeePool, err := api.storage.GetAllEmployees()
	if err != nil {
		return err
	}

	// filter out employees that are ineligible for scheduling
	employeePool = filterEmployeesByID(sch.EmployeeFilter, employeePool)
	employeePool = filterInactiveEmployee(employeePool)

	for _, date := range dates {
		err := sch.ScheduleDay(date, employeePool)
		if err != nil {
			return err
		}
	}

	return nil
}

// getDates returns a list of consecutive dates as time.Time objects given a start date and an end date in format mm-dd-yyy
func getDates(start, end string) ([]time.Time, error) {

	currentDate, err := time.Parse("2006-01-02", start)
	if err != nil {
		return nil, fmt.Errorf("could not parse start date; should be in format mm-dd-yyyy: %v", err)
	}

	endDate, err := time.Parse("2006-01-02", end)
	if err != nil {
		return nil, fmt.Errorf("could not parse end date; should be in format mm-dd-yyyy: %v", err)
	}

	dates := []time.Time{}

	for {
		if currentDate.After(endDate) {
			break
		}
		dates = append(dates, currentDate)
		currentDate = currentDate.AddDate(0, 0, 1)
	}

	return dates, nil
}

func filterInactiveEmployee(employees map[string]*model.Employee) map[string]*model.Employee {
	filtered := map[string]*model.Employee{}

	for id, employee := range employees {
		if employee.Status != model.EmployeeActive {
			continue
		}
		filtered[id] = employee
	}

	return filtered
}

func filterEmployeesByID(filter []string, employees map[string]*model.Employee) map[string]*model.Employee {

	if len(filter) == 0 {
		return employees
	}

	filtered := map[string]*model.Employee{}
	filterMap := map[string]struct{}{}

	for _, id := range filter {
		filterMap[id] = struct{}{}
	}

	for id, employee := range employees {
		if _, exists := filterMap[id]; exists {
			filtered[id] = employee
		}
	}

	return filtered
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
