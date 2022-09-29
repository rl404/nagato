package main

import (
	"context"
	"encoding/json"
	"os"

	"github.com/rl404/nagato/mal"
)

// Credential from myanimelist.
// You can get it by registering here.
// https://myanimelist.net/apiconfig.
//
// Client secret may be empty depending on
// your app type when registering.
const clientID = "mal-client-id"
const clientSecret = "mal-client-secret"
const state = ""

var malClient *mal.Client

func main() {

	// Use this if you want to retrieve public info only.
	clientPublic()

	// Use this if you want to retrieve public info
	// and update/retrieve personal data.
	clientWithOauth2()

	// Run example.
	example()
}

func clientPublic() {
	malClient = mal.NewPublic(clientID)
}

func clientWithOauth2() {
	// Generating token should be done outside your project flow
	// and should be done one time only.
	if err := generateToken(context.Background(), clientID, clientSecret, state); err != nil {
		panic(err)
	}

	cachedToken, err := loadCachedToken()
	if err != nil {
		panic(err)
	}

	malClient, err = mal.NewWithOauth2(mal.Oauth2Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		State:        state,
		AccessToken:  cachedToken.AccessToken,
		RefreshToken: cachedToken.RefreshToken,
		Expiry:       cachedToken.Expiry,
	})
	if err != nil {
		panic(err)
	}
}

func toJson(name string, d interface{}) {
	jsonData, _ := json.MarshalIndent(d, "", " ")

	f, _ := os.Create(name + ".json")
	defer f.Close()

	f.Write(jsonData)
}
