package main

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// App struct
type App struct {
	ctx context.Context
}

// User login type for doing the requests on poe trade
type UserAuth struct {
	Username string // username for trade site
	AccessToken string // token used to auth on trade
	RefreshToken string // used to refresh the token
}

type User struct {
	Username string
	Password string
}

type AuthorizeResponse struct {
	State string
	Code string

}

type TokenRequest struct {
	ClientID string
	ClientSecret string
	GrantType string // always authorization_code
	Code string
	RedirectURL string
	Scope string
	CodeVerifier string

}

// Item query type for storing the item data
func SignIn() bool {
	result := false
	// get a random 32 byte value 
	secret := make([]byte, 32)
	if _, err := rand.Read(secret); err != nil{
		panic(err)
	}
	
	// encode the random 32 byte value by base 64
	codeVerifier := base64.RawURLEncoding.EncodeToString(secret)

	// encode the sha256 hash of the code verifier	
	hash := sha256.Sum256([]byte(codeVerifier))
	codeChallenge := base64.RawURLEncoding.EncodeToString(hash[:])

	var clientId = "PriceChecker"
	var responseType = "code"
	var scope = "account:profile"
	var state = "awdadawd" /// i need to generate this
	var redirectUri = "wdawdawd" // i need to have some redirect to know i go to the next step
	var codeChallengeMethod = "S256"
	var authorizeURL = "https://www.pathofexile.com/oauth/authorize" + 
	"?client_id=" + clientId +
	"&response_type=" + responseType +
	"&scope" + scope + 
	"&state=" + state + 
	"&redirect_uri=" + redirectUri + 
	"&code_challenge=" + codeChallenge +
	"&code_challendge_method" + codeChallengeMethod

	res, err := http.Get(authorizeURL); if err != nil {
		log.Fatal("Couldn't complete authorization", err)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	res.Body.Close()
	if res.StatusCode > 299 {
		log.Fatalf("Reponse failed with status code: %d")
	}

	var authResponse AuthorizeResponse
	json.Unmarshal(body, &authResponse)


	var tokenURL = "http://pathofexile.com/oauth/token"
	var contentType = "application/x-www-form-urlencoded"

	if authResponse.State == state {
		res, err := http.Post(tokenURL, contentType, ); if err != nil {
			log.Fatal("Couldn't complete token exchange", err)
		}
		// continue
	} else {
		log.Fatalf("state returned doesn't match generated state: %s", state)
	}













	return result
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}
