package server

import (
	"golang_project/serverconst"
	"net/http"
	"strings"
)

func getErrorStatusCode(message string) int {
	if strings.HasPrefix(message, serverconst.NotFoundErr) {
		return http.StatusNotFound
	}
	if strings.HasPrefix(message, serverconst.BadRequest) || strings.HasPrefix(message, serverconst.ValidationErr) {
		return http.StatusBadRequest
	}
	return http.StatusInternalServerError
}
