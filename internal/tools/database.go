package tools

import (
	log "github.com/sirupsen/logrus"
)

// Database Collections
type LoginDetails struct {
	AuthToken string //Authentication token for the user
	Username  string //Username of the user
}

type CoinDetails struct {
	Coins    int64  //Balance of coins for the user
	Username string //Username of the user
}

type DatabaseInterface interface {
	GetUserLoginDetails(username string) *LoginDetails //Get the login details for a user by username
	GetUserCoins(username string) *CoinDetails         //Get the coin details for a user by username
	SetupDatabase() error                              //Setup the database connection
}

func NewDatabase() (*DatabaseInterface, error) {
	var database DatabaseInterface = &mockDB{} //Create a new instance of the Database struct

	var err error = database.SetupDatabase()
	if err != nil {
		log.Error(err)
		return nil, err //Return nil and the error if there is an error creating the database
	}

	return &database, nil //Return a pointer to the database and nil error if successful
}
