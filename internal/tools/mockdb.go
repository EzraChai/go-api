package tools

import "time"

type mockDB struct{}

var mockLoginDetails = map[string]LoginDetails{
	"ezra": {
		Username:  "ezra",
		AuthToken: "1234567890",
	},
	"jason": {
		Username:  "jason",
		AuthToken: "1234567890jason",
	},
	"chloe": {
		Username:  "chloe",
		AuthToken: "chloe1234567890",
	},
}

var mockCoinDetails = map[string]CoinDetails{
	"ezra": {
		Coins:    1000,
		Username: "ezra",
	},
	"jason": {
		Coins:    200,
		Username: "jason",
	},
	"chloe": {
		Coins:    500,
		Username: "chloe",
	},
}

func (d *mockDB) GetUserCoins(username string) *CoinDetails {
	//	Simulate DB call
	time.Sleep(time.Millisecond * 100)

	var clientData = CoinDetails{}
	clientData, ok := mockCoinDetails[username]
	if !ok {
		return nil
	}

	return &clientData
}

func (d *mockDB) GetLoginDetails(username string) *LoginDetails {
	//	Simulate DB call
	time.Sleep(time.Millisecond * 100)

	var clientData = LoginDetails{}
	clientData, ok := mockLoginDetails[username]
	if !ok {
		return nil
	}

	return &clientData
}

func (d *mockDB) SetupDatabase() error {
	return nil
}
