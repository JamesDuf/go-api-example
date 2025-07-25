package api

import (
	"encoding/json"
	"net/http"
)

// Coin balance params
// parameters the coin balance API endpoint will take
type CoinBalanceParams struct {
	Username string
}

// Successful response from the coin balance API
type CoinBalanceResponse struct {
	//Success code, usually 200
	Code int

	//Message describing the response
	//Account Balance
	Balance int64
}

type Error struct {
	//Error code, usually 400 for bad request
	Code int

	//Error message describing the issue
	Message string
}

func writeError(w http.ResponseWriter, code int, message string) {
	response := Error{
		Code:    code,
		Message: message,
	}
	w.Header().Set("Content-Type", "application/json") //Set the content type of the response to JSON
	w.WriteHeader(code)                                //Set the HTTP status code

	json.NewEncoder(w).Encode(response)
}

var (
	RequestErrorHandler = func(w http.ResponseWriter, err error) {
		writeError(w, http.StatusBadRequest, err.Error()) //Write a 400 Bad Request error with the error message
	}
	InternalErrorHandler = func(w http.ResponseWriter) {
		writeError(w, http.StatusInternalServerError, "An unexpected error occurred.") //Write a 500 Internal Server Error
	}
)
