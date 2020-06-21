package api

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/clintjedwards/scheduler/models"
	"github.com/clintjedwards/scheduler/utils"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
)

//TODO(clintjedwards): validate params of all handlers

// ListPositionsHandler returns all positions unpaginated
func (api *API) ListPositionsHandler(w http.ResponseWriter, r *http.Request) {

	positions, err := api.storage.GetAllPositions()
	if err != nil {
		log.Error().Err(err).Msg("failed to retrieve positions")
		sendResponse(w, http.StatusBadGateway, nil, fmt.Errorf("failed to retrieve positions"))
		return
	}

	sendResponse(w, http.StatusOK, positions, nil)
}

// AddPosition adds a new position to the scheduler service
func (api *API) AddPosition(w http.ResponseWriter, r *http.Request) {

	newPosition := models.Position{}

	err := parseJSON(r.Body, &newPosition)
	if err != nil {
		log.Warn().Err(err).Msg("could not parse json request")
		sendResponse(w, http.StatusBadRequest, nil, fmt.Errorf("could not parse json request: %v", err))
		return
	}
	defer r.Body.Close()

	newPosition.ID = string(utils.GenerateRandString(api.config.IDLength))

	err = api.storage.AddPosition(newPosition.ID, &newPosition)
	if err != nil {
		if errors.Is(err, utils.ErrEntityExists) {
			sendResponse(w, http.StatusConflict, nil, utils.ErrEntityExists)
			return
		}
		log.Error().Err(err).Msg("could not add position")
		sendResponse(w, http.StatusBadGateway, nil, fmt.Errorf("could not add position"))
		return
	}

	log.Info().Interface("position", newPosition).Msg("created new position")
	sendResponse(w, http.StatusOK, newPosition, nil)
}

// GetPosition returns a single position by id
func (api *API) GetPosition(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	position, err := api.storage.GetPosition(vars["id"])
	if err != nil {
		if errors.Is(err, utils.ErrEntityNotFound) {
			sendResponse(w, http.StatusNotFound, nil, utils.ErrEntityNotFound)
			return
		}
		log.Error().Err(err).Msg("could not get position")
		sendResponse(w, http.StatusBadGateway, nil, utils.ErrEntityNotFound)
		return
	}

	sendResponse(w, http.StatusOK, position, nil)
}
