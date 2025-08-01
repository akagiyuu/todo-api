package util

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func ReadRequest[T any](v *T, w http.ResponseWriter, r *http.Request) bool {
	b, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("error reading JSON body: %v", err)
		WriteError(&ApiError{
			Message: "Invalid body",
		}, w)

		return false
	}

	err = json.Unmarshal(b, v)
	if err != nil {
		WriteError(&ApiError{
			Message: "Invalid body",
			Detail: map[string]any{
				"detail": err.Error(),
			},
		}, w)

		return false
	}

	return true
}

func WriteResponse(response any, w http.ResponseWriter) {
	jsonResp, err := json.Marshal(response)
	if err != nil {
		log.Printf("Failed to marshal response: %v", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if _, err := w.Write(jsonResp); err != nil {
		log.Printf("Failed to write response: %v", err)
	}
}
