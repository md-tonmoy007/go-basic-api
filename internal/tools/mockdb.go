package tools


import (
	"time"

)



type mockDB struct {}


var mockLoginDetails = map[string]LoginDetails{
	"tonmoy": {
		Username: "tonmoy",
		AuthToken: "12345",
	},
	"alice": {
		Username: "alice",
		AuthToken: "abcdef1234567890",
	},
	"bob": {
		Username: "bob",
		AuthToken: "fedcba0987654321",
	},
}


var mockCoinDetails = map[string]CoinDetails{
	"tonmoy": {
		Username: "tonmoy",
		Coins: 1000,
	},
	"alice": {
		Username: "alice",
		Coins: 500,
	},
	"bob": {
		Username: "bob",
		Coins: 750,
	},
}

func (db *mockDB) GetUserLoginDetails(username string) *LoginDetails {
	time.Sleep(time.Second * 1)

	var clientData = LoginDetails{}

	clientData, ok := mockLoginDetails[username]
	if !ok {
		return nil
	}
	return &clientData
}


func (db *mockDB) GetUserCoinDetails(username string) *CoinDetails {
	time.Sleep(time.Second * 1)

	var clientData = CoinDetails{}

	clientData, ok := mockCoinDetails[username]
	if !ok {
		return nil
	}
	return &clientData
}

func (db *mockDB) SetupDatabase() error {
	// Mock setup, no actual database connection needed
	return nil
}