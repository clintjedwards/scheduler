package api

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/clintjedwards/scheduler/models"
	"github.com/clintjedwards/scheduler/utils"
	"github.com/gorilla/mux"
	"github.com/mitchellh/copystructure"
	"github.com/rs/zerolog/log"
)

func (api *API) generateSchedule(sch *models.Schedule) error {

	startDate, err := time.Parse("01-02-2006", sch.Start)
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

	currentDate := startDate

	//TODO(clintjedwards): settings can be nil which can cause panics, we should validate settings and create
	// non nil values before passing it to other functions
	// We start from the start date and then iterate through each day until we hit the desired length
	for {
		if currentDate.After(endDate) {
			break
		}

		datef := startDate.Format("01-02-2006")

		switch weekday := startDate.Weekday(); weekday {
		case time.Monday:
			day, err := scheduleDay(employees, sch.Program.Monday)
			if err != nil {
				return err
			}
			sch.TimeTable[datef] = day
		case time.Tuesday:
			day, err := scheduleDay(employees, sch.Program.Tuesday)
			if err != nil {
				return err
			}
			sch.TimeTable[datef] = day
		case time.Wednesday:
			day, err := scheduleDay(employees, sch.Program.Wednesday)
			if err != nil {
				return err
			}
			sch.TimeTable[datef] = day
		case time.Thursday:
			day, err := scheduleDay(employees, sch.Program.Thursday)
			if err != nil {
				return err
			}
			sch.TimeTable[datef] = day
		case time.Friday:
			day, err := scheduleDay(employees, sch.Program.Friday)
			if err != nil {
				return err
			}
			sch.TimeTable[datef] = day
		case time.Saturday:
			day, err := scheduleDay(employees, sch.Program.Saturday)
			if err != nil {
				return err
			}
			sch.TimeTable[datef] = day
		case time.Sunday:
			day, err := scheduleDay(employees, sch.Program.Sunday)
			if err != nil {
				return err
			}
			sch.TimeTable[datef] = day
		default:
			log.Error().Msgf("could not generate day %q; not a valid weekday", weekday)
		}
		currentDate = currentDate.AddDate(0, 0, 1)
	}

	return nil
}

// scheduleDay will insert employees into roles they are best suited for
// this alters the given map.
// We return a new copy of what is the positions structure because we want to leave the positions data structure intact
func scheduleDay(employees map[string]models.Employee, shifts map[string][]models.Shift) (map[string][]models.Shift, error) {
	if shifts == nil || employees == nil {
		return map[string][]models.Shift{}, nil
	}
	availableEmployees := map[string]models.Employee{}

	rawDay, err := copystructure.Copy(shifts)
	if err != nil {
		log.Error().Err(err).Msg("could not copy days")
		return nil, err
	}
	day := rawDay.(map[string][]models.Shift)

	for employeeID, employee := range employees {
		availableEmployees[employeeID] = employee
	}

	for positionID, shifts := range day {
		for _, shift := range shifts {
			for id, employee := range availableEmployees {
				if _, exists := employee.Positions[positionID]; !exists {
					continue
				}

				shift.Employee = id
				delete(availableEmployees, id)
				break
			}
			if shift.Employee == "" {
				// TODO(clintjedwards): Return a collection of errors to the calling function,
				/// so it can be assigned to the day and reported back to the user
				log.Warn().Msgf("could not find eligible employee for postion: %s; shift %s-%s",
					positionID, shift.Start, shift.End)
			}
		}
	}

	return day, nil
}

// getEmployees returns a map of eligible employees for future scheduling.
// TODO(clintjedwards): This datastructure should be more advanced, something like a heap would work nicely here
// since eventually employees will have a weight. We might have to calculate weights for each
// position ahead of time and then use that here later
func (api *API) getEmployees(employeeFilter []string) (map[string]models.Employee, error) {

	eligibleEmployees := map[string]models.Employee{}

	employees, err := api.storage.GetAllEmployees()
	if err != nil {
		return nil, err
	}

	if len(employeeFilter) != 0 {
		for _, employee := range employeeFilter {
			if employees[employee].Status != models.EmployeeActive {
				continue
			}
			eligibleEmployees[employee] = *employees[employee]
		}
		return eligibleEmployees, nil
	}

	for id, employee := range employees {
		if employee.Status != models.EmployeeActive {
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

	newSchedule := models.Schedule{}

	err := parseJSON(r.Body, &newSchedule)
	if err != nil {
		log.Warn().Err(err).Msg("could not parse json request")
		sendErrResponse(w, http.StatusBadRequest, err)
		return
	}
	defer r.Body.Close()

	newSchedule.ID = string(utils.GenerateRandString(api.config.IDLength))
	newSchedule.TimeTable = map[string]map[string][]models.Shift{}
	err = api.generateSchedule(&newSchedule)
	if err != nil {
		sendErrResponse(w, http.StatusBadGateway, err)
		return
	}

	err = api.storage.AddSchedule(newSchedule.ID, &newSchedule)
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
