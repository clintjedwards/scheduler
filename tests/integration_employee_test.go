package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/clintjedwards/scheduler/api"
	"github.com/clintjedwards/scheduler/app"
	"github.com/clintjedwards/scheduler/config"
	"github.com/clintjedwards/scheduler/storage"
	"github.com/rs/zerolog/log"
)

func TestListEmployees(t *testing.T) {
	config, err := config.FromEnv()
	if err != nil {
		log.Fatal().Err(err).Msg("could not get config in order to start services")
	}

	storage, err := app.InitStorage(storage.EngineType(config.Database.Engine))
	if err != nil {
		log.Panic().Err(err).Msg("could not init storage")
	}

	api := api.NewAPI(nil, storage)

	req, err := http.NewRequest("GET", "/employees", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(api.ListEmployeesHandler)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v; want %v", status, http.StatusOK)
	}
}
