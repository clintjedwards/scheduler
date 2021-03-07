package api

import (
	"errors"
	"net/http"

	"github.com/clintjedwards/scheduler/model"
	"github.com/clintjedwards/scheduler/utils"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
)

// ListProgramsHandler returns all programs unpaginated
func (api *API) ListProgramsHandler(w http.ResponseWriter, r *http.Request) {

	programs, err := api.storage.GetAllPrograms()
	if err != nil {
		log.Error().Err(err).Msg("failed to retrieve programs")
		sendErrResponse(w, http.StatusBadGateway, err)
		return
	}

	sendResponse(w, http.StatusOK, programs)
}

// AddProgramHandler adds a new program to the scheduler service
func (api *API) AddProgramHandler(w http.ResponseWriter, r *http.Request) {

	newProgram := model.Program{}

	err := parseJSON(r.Body, &newProgram)
	if err != nil {
		log.Warn().Err(err).Msg("could not parse json request")
		sendErrResponse(w, http.StatusBadRequest, err)
		return
	}
	defer r.Body.Close()

	newProgram.ID = string(utils.GenerateRandString(api.config.IDLength))

	err = api.storage.AddProgram(newProgram.ID, &newProgram)
	if err != nil {
		if errors.Is(err, utils.ErrEntityExists) {
			sendErrResponse(w, http.StatusConflict, err)
			return
		}
		log.Error().Err(err).Msg("could not add program")
		sendErrResponse(w, http.StatusBadGateway, err)
		return
	}

	log.Info().Interface("program", newProgram).Msg("created new program")
	sendResponse(w, http.StatusCreated, newProgram)
}

// GetProgramHandler returns a single program by id
func (api *API) GetProgramHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	program, err := api.storage.GetProgram(vars["id"])
	if err != nil {
		if errors.Is(err, utils.ErrEntityNotFound) {
			sendErrResponse(w, http.StatusNotFound, err)
			return
		}
		log.Error().Err(err).Msg("could not get program")
		sendErrResponse(w, http.StatusBadGateway, err)
		return
	}

	sendResponse(w, http.StatusOK, program)
}

// DeleteProgramHandler removes an employee by id
func (api *API) DeleteProgramHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	err := api.storage.DeleteProgram(vars["id"])
	if err != nil {
		if errors.Is(err, utils.ErrEntityNotFound) {
			http.Error(w, "program not found", http.StatusNotFound)
			return
		}
		log.Error().Err(err).Msg("could not delete program")
		sendErrResponse(w, http.StatusBadGateway, err)
		return
	}

	sendResponse(w, http.StatusOK, nil)
}
