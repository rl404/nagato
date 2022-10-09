package main

import (
	"context"
	"encoding/json"
	"os"

	"github.com/rl404/nagato"
	"github.com/rl404/nagato/mal"
)

const clientID = "mal-client-id"
const clientSecret = "mal-client-secret"
const state = ""

var nagatoClient *nagato.Client

func main() {
	clientPublic()

	clientOauth2()

	example()
}

func clientPublic() {
	nagatoClient = nagato.New(clientID)
}

func clientOauth2() {
	nagatoClient = nagato.New(clientID)

	// Generating token should be done outside your project flow
	// and should be done one time only.
	if err := generateToken(context.Background(), clientID, clientSecret, state); err != nil {
		panic(err)
	}

	cachedToken, err := loadCachedToken()
	if err != nil {
		panic(err)
	}

	malWithOauth2, err := mal.NewWithOauth2(mal.Oauth2Config{
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

	nagatoClient.SetMalClient(malWithOauth2)
}

func toJson(name string, d interface{}) {
	jsonData, _ := json.MarshalIndent(d, "", " ")

	f, _ := os.Create(name + ".json")
	defer f.Close()

	f.Write(jsonData)
}
