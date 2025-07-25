package tools

import (
	"time"
)

type mockDB struct{}

// Simulated in-memory storage for user login details and coin balances
var mockLoginDetails = map[string]LoginDetails{
	"alex": {
		AuthToken: "123ABC",
		Username:  "alex",
	},
	"jane": {
		AuthToken: "456DEF",
		Username:  "jane",
	},
	"marie": {
		AuthToken: "789GHI",
		Username:  "marie",
	},
}

var mockCoinDetails = map[string]CoinDetails{
	"alex": {
		Coins:    1000,
		Username: "alex",
	},
	"jane": {
		Coins:    2000,
		Username: "jane",
	},
	"marie": {
		Coins:    3000,
		Username: "marie",
	},
}

func (d *mockDB) GetUserLoginDetails(username string) *LoginDetails {
	//Simulate DB call
	time.Sleep(time.Second * 1)

	var clientData = LoginDetails{}
	clientData, ok := mockLoginDetails[username]
	if !ok {
		return nil
	}

	return &clientData
}

func (d *mockDB) GetUserCoins(username string) *CoinDetails {
	//Simulate DB call
	time.Sleep(time.Second * 1)

	var clientData = CoinDetails{}
	clientData, ok := mockCoinDetails[username]
	if !ok {
		return nil
	}

	return &clientData
}

func (d *mockDB) SetupDatabase() error {
	return nil
}
