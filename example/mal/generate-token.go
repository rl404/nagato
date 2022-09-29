package main

import (
	"bufio"
	"context"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"runtime"

	"golang.org/x/oauth2"
)

func generateToken(ctx context.Context, clientID, clientSecret, state string) error {
	oauth2Cfg := &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Endpoint: oauth2.Endpoint{
			AuthURL:   "https://myanimelist.net/v1/oauth2/authorize",
			TokenURL:  "https://myanimelist.net/v1/oauth2/token",
			AuthStyle: oauth2.AuthStyleInParams,
		},
	}

	oauth2Token, err := loadCachedToken()
	if err == nil {
		refreshedToken, err := oauth2Cfg.TokenSource(ctx, oauth2Token).Token()
		if err == nil && (oauth2Token != refreshedToken) {
			if err := cacheToken(*refreshedToken); err != nil {
				return err
			}
			return nil
		}
		return nil
	}

	// Generate a code verifier, a high-entropy cryptographic random string. It
	// will be set as the code_challenge in the authentication URL. It should
	// have a minimum length of 43 characters and a maximum length of 128
	// characters.
	const codeVerifierLength = 128
	codeVerifier, err := generateCodeVerifier(codeVerifierLength)
	if err != nil {
		return fmt.Errorf("generating code verifier: %v", err)
	}

	// Produce the authentication URL where the user needs to be redirected and
	// allow your application to access their MyAnimeList data.
	authURL := oauth2Cfg.AuthCodeURL(state, oauth2.SetAuthURLParam("code_challenge", codeVerifier))
	if err := openBrowser(authURL); err != nil {
		fmt.Println("Could not open browser.")
	}

	fmt.Printf("Your browser should open: %v\n", authURL)
	fmt.Print("After authenticating, copy the code from the browser URL and paste it here: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	code := scanner.Text()
	if scanner.Err() != nil {
		return fmt.Errorf("reading code from terminal: %v", err)
	}

	// Exchange the authentication code for a token. MyAnimeList currently only
	// supports the plain code_challenge_method so to verify the string, just
	// make sure it is the same as the one you entered in the code_challenge.
	token, err := oauth2Cfg.Exchange(ctx, code, oauth2.SetAuthURLParam("code_verifier", codeVerifier))
	if err != nil {
		return fmt.Errorf("exchanging code for token: %v", err)
	}

	fmt.Println("Authentication was successful. Caching oauth2 token...")
	if err := cacheToken(*token); err != nil {
		return fmt.Errorf("caching oauth2 token: %s", err)
	}

	return nil
}

const cacheName = "auth-token-cache.json"

func cacheToken(token oauth2.Token) error {
	b, err := json.MarshalIndent(token, "", "   ")
	if err != nil {
		return err
	}
	if err := os.WriteFile(cacheName, b, 0644); err != nil {
		return err
	}
	return nil
}

func loadCachedToken() (*oauth2.Token, error) {
	b, err := os.ReadFile(cacheName)
	if err != nil {
		return nil, err
	}
	var token oauth2.Token
	if err := json.Unmarshal(b, &token); err != nil {
		return nil, err
	}
	return &token, nil
}

func generateCodeVerifier(length int) (string, error) {
	const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstvuwxyz0123456789-._~"
	bytes := make([]byte, length)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	for i, b := range bytes {
		bytes[i] = charset[b%byte(len(charset))]
	}
	return string(bytes), nil
}

func openBrowser(url string) error {
	switch runtime.GOOS {
	case "linux":
		return exec.Command("xdg-open", url).Start()
	case "windows":
		return exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		return exec.Command("open", url).Start()
	default:
		return fmt.Errorf("openBrowser: unsupported operating system: %v", runtime.GOOS)
	}
}
