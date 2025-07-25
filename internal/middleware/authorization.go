package middleware

import (
	"errors"
	"net/http"

	"github.com/JamesDuf/go-api-tutorial/api"
	"github.com/JamesDuf/go-api-tutorial/internal/tools"
	log "github.com/sirupsen/logrus"
)

var UnAuthorizedError = errors.New("Invalid credentials.") // UnAuthorizedError is an error that indicates the user is not authorized to access a resource

func Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { //ResponseWriter is used to send a response back to the client, Request contains information about the HTTP request (response body, header, status code)

		var username string = r.URL.Query().Get("username") //Get the Username from the request header
		var token = r.Header.Get("Authorization")           //Get the Authorization token from the request header
		var err error

		if username == "" || token == "" { //Check if the username or token is empty
			log.Error(UnAuthorizedError)                  //Log the error if the username or token is empty
			api.RequestErrorHandler(w, UnAuthorizedError) //Call the RequestErrorHandler to handle the error
			return                                        //Return early if the username or token is empty
		}
		var database *tools.DatabaseInterface
		database, err = tools.NewDatabase() //Get the database connection
		if err != nil {
			api.InternalErrorHandler(w) //Call the InternalErrorHandler to handle the error
			return                      //Return early if there is an error getting the database connection
		}

		var loginDetails *tools.LoginDetails
		loginDetails = (*database).GetUserLoginDetails(username) //Get the login details for the username from the database

		if loginDetails == nil || (token != (*loginDetails).AuthToken) { //Check if the login details are nil or if the token does not match
			log.Error(UnAuthorizedError)                  //Log the error if the login details are nil or if the token does not match
			api.RequestErrorHandler(w, UnAuthorizedError) //Call the RequestErrorHandler to handle the error
			return                                        //Return early if the login details are nil or if the token does not match
		}

		next.ServeHTTP(w, r) //If the username and token are valid, call the next handler in the chain
	})
}
