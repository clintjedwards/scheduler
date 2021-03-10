package api

import (
	"errors"
	"net/http"
	"strings"

	"github.com/rs/zerolog/log"
)

var appVersion = "v0.0.dev_<build_time>_<commit>"

// Info represents application information gleaned from the version string
type info struct {
	semver string
	epoch  string
	hash   string
}

// Parse takes a very specific version string format:
// <semver>_<epoch_time>_<git_hash>
// and returns its individual parts
func parse(version string) (info, error) {
	versionTuple := strings.Split(version, "_")

	if len(versionTuple) != 3 {
		return info{},
			errors.New("version not in correct format: <semver>_<epoch_time>_<git_hash>")
	}

	return info{
		semver: versionTuple[0],
		epoch:  versionTuple[1],
		hash:   versionTuple[2],
	}, nil
}

// GetSystemInfoHandler returns system information and health
func (api *API) GetSystemInfoHandler(w http.ResponseWriter, r *http.Request) {

	info, err := parse(appVersion)
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
		BuildTime:       info.epoch,
		Commit:          info.hash,
		DebugEnabled:    api.config.Debug,
		FrontendEnabled: api.config.Frontend,
		Semver:          info.semver,
	}

	sendResponse(w, http.StatusOK, systemInfo)
}
