package util

import "net/http"

type ApiError struct {
	Message string
	Detail  map[string]any
}

func WriteError(e *ApiError, w http.ResponseWriter) {
	 	http.Error(w, e.Message, http.StatusBadRequest)
	if e.Detail != nil {
		WriteResponse(e.Detail, w)
	}
}
