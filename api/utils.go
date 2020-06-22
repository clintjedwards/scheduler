package api

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/rs/zerolog/log"
)

// sendResponse converts raw objects and parameters to a json response
// and passes it to a provided writer.
func sendResponse(w http.ResponseWriter, httpStatusCode int, payload interface{}) {
	w.WriteHeader(httpStatusCode)

	enc := json.NewEncoder(w)
	err := enc.Encode(payload)
	if err != nil {
		log.Error().Err(err).Msgf("could not encode json response: %v", err)
	}
}

// sendErrResponse converts raw objects and parameters to a json response specifically for erorrs
// and passes it to a provided writer. The creation of a separate function for just errors,
// is due to how they are handled differently from other payload types.
func sendErrResponse(w http.ResponseWriter, httpStatusCode int, appErr error) {
	w.WriteHeader(httpStatusCode)

	enc := json.NewEncoder(w)
	err := enc.Encode(map[string]string{"err": appErr.Error()})
	if err != nil {
		log.Error().Err(err).Msgf("could not encode json response: %v", err)
	}
}

// parseJSON parses the given json request into interface
func parseJSON(rc io.Reader, object interface{}) error {
	decoder := json.NewDecoder(rc)
	err := decoder.Decode(object)
	if err != nil {
		return err
	}
	return nil
}

//DefaultHeaders is a wrapper function setting the reponse headers
// func DefaultHeaders(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		w.Header().Set("Content-Type", "application/json")
// 		w.Header().Set("Access-Control-Allow-Origin", "*")

// 		next.ServeHTTP(w, r)
// 	})
// }
