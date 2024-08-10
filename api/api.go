package api

import (
	"encoding/json"
	"net/http"
)

type CoinBalanceParams struct {
	Username string
}

type CoinBalanceResponse struct {
	//	Success Code
	Code int

	//	AccountBalance
	Balance int64
}

type Error struct {
	//	Success Code
	Code int

	//	Error Message
	Message string
}

func writeError(w http.ResponseWriter, code int, message string) {
	resp := Error{
		Code:    code,
		Message: message,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	json.NewEncoder(w).Encode(resp)
}

var (
	RequestErrorHandler = func(w http.ResponseWriter, err error) {
		writeError(w, http.StatusBadRequest, err.Error())
	}

	InternalErrorHandler = func(w http.ResponseWriter) {
		writeError(w, http.StatusInternalServerError, "An internal error occurred.")
	}
)