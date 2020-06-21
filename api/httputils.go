package api

import (
	"encoding/json"
	"io"
	"net/http"
)

// sendResponse formats and sends a message to supplied writer in json format
func sendResponse(w http.ResponseWriter, httpStatusCode int, data interface{}, httpErr error) error {
	w.Header().Set("Content-Type", "application/json")

	var errStr string
	if httpErr != nil {
		errStr = httpErr.Error()
	}

	err := json.NewEncoder(w).Encode(struct {
		StatusText string      `json:"status_text"`
		Data       interface{} `json:"data"`
		Error      string      `json:"error"`
	}{http.StatusText(httpStatusCode), data, errStr})
	if err != nil {
		return err
	}
	return nil
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
