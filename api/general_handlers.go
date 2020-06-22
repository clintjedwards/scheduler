package api

import (
	"net/http"

	"github.com/clintjedwards/toolkit/version"
	"github.com/rs/zerolog/log"
)

var appVersion = "v0.0.dev_<build_time>_<commit>"

// GetSystemInfoHandler returns system information and health
func (api *API) GetSystemInfoHandler(w http.ResponseWriter, r *http.Request) {

	info, err := version.Parse(appVersion)
	if err != nil {
		log.Error().Err(err).Msg("could not parse version string")
		sendErrResponse(w, http.StatusBadGateway, err)
		return
	}

	systemInfo := struct {
		BuildTime       string `json:"build_time"`
		Commit          string `json:"commit"`
		DebugEnabled    bool   `json:"debug_enabled"`
		FrontendEnabled bool   `json:"frontend_enabled"`
		Semver          string `json:"semver"`
	}{
		BuildTime:       info.Epoch,
		Commit:          info.Hash,
		DebugEnabled:    api.config.Debug,
		FrontendEnabled: api.config.Frontend,
		Semver:          info.Semver,
	}

	sendResponse(w, http.StatusOK, systemInfo)
}
