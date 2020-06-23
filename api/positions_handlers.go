package api

import (
	"errors"
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
		sendErrResponse(w, http.StatusBadGateway, err)
		return
	}

	sendResponse(w, http.StatusOK, positions)
}

// AddPositionHandler adds a new position to the scheduler service
func (api *API) AddPositionHandler(w http.ResponseWriter, r *http.Request) {

	newPosition := models.Position{}

	err := parseJSON(r.Body, &newPosition)
	if err != nil {
		log.Warn().Err(err).Msg("could not parse json request")
		sendErrResponse(w, http.StatusBadRequest, err)
		return
	}
	defer r.Body.Close()

	newPosition.ID = string(utils.GenerateRandString(api.config.IDLength))

	err = api.storage.AddPosition(newPosition.ID, &newPosition)
	if err != nil {
		if errors.Is(err, utils.ErrEntityExists) {
			sendErrResponse(w, http.StatusConflict, err)
			return
		}
		log.Error().Err(err).Msg("could not add position")
		sendErrResponse(w, http.StatusBadGateway, err)
		return
	}

	log.Info().Interface("position", newPosition).Msg("created new position")
	sendResponse(w, http.StatusCreated, newPosition)
}

// GetPositionHandler returns a single position by id
func (api *API) GetPositionHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	position, err := api.storage.GetPosition(vars["id"])
	if err != nil {
		if errors.Is(err, utils.ErrEntityNotFound) {
			sendErrResponse(w, http.StatusNotFound, err)
			return
		}
		log.Error().Err(err).Msg("could not get position")
		sendErrResponse(w, http.StatusBadGateway, err)
		return
	}

	sendResponse(w, http.StatusOK, position)
}
