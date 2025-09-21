package main

import (
	"bytes"
	"context"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

// App struct
type App struct {
	ctx context.Context
}

// User login type for doing the requests on poe trade
type UserAuth struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int64  `json:"expires_in"`
	TokenType    string `json:"bearer"`
	Scope        string `json:"scope"`
	Username     string `json:"username"`
	Sub          string `json:"sub"`
	RefreshToken string `json:"refresh_token"`
}

type User struct {
	Username string
	Password string
}

type AuthorizeResponse struct {
	State string `json:"state"`
	Code  string `json:"code"`
}

type TokenRequest struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	GrantType    string `json:"grant_type"`
	Code         string `json:"code"`
	RedirectURL  string `json:"redirect_uri"`
	Scope        string `json:"scope"`
	CodeVerifier string `json:"code_verifier"`
}

type RefreshToken struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	GrantType    string `json:"grant_type"`
	RefreshToken string `json:"refresh_token"`
}

type Modifier struct {
}

type Item struct {
	First  Modifier `json:"stat"`
	Second Modifier
	Third  Modifier
	Fourth Modifier
	Fifth  Modifier
	Sixth  Modifier
}

var userAuth UserAuth // contains user authorization data

// Item query type for storing the item data

// does through the sign up process for the user
func SignIn() bool {
	result := false
	// get a random 32 byte value
	secret := make([]byte, 32)
	if _, err := rand.Read(secret); err != nil {
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
	var state = "awdadawd"       /// i need to generate this
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

	res, err := http.Get(authorizeURL)
	if err != nil {
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

	if authResponse.State == state {
		var tokenURL = "http://pathofexile.com/oauth/token"
		var contentType = "application/x-www-form-urlencoded"

		tokenRequest := TokenRequest{
			ClientID:     "PriceChecker",
			ClientSecret: "asdasdasd",          // i need to create this
			GrantType:    "authorization_code", // always authorization_code
			Code:         authResponse.Code,
			RedirectURL:  redirectUri, // need to figure this out
			Scope:        scope,
			CodeVerifier: codeVerifier,
		}

		jsonData, err := json.Marshal(tokenRequest)
		if err != nil {
			log.Fatal(err)
		}

		res, err := http.Post(tokenURL, contentType, bytes.NewBuffer(jsonData))
		if err != nil {
			log.Fatal("Couldn't complete token exchange", err)
		}

		body, err := io.ReadAll(res.Body)
		if err != nil {
			log.Fatal(err)
		}
		res.Body.Close()

		// at this point we will want to add this information to the user.json file but also keep it for use
		err = os.WriteFile("user.json", body, 0644)
		if err != nil {
			panic(err)
		}

		// data i can use within the application
		json.Unmarshal(body, &userAuth)

		// converting the expiration duration to an expiration time for revaliding
		expiresAt := time.Now().Add(time.Duration(userAuth.ExpiresIn) * time.Second)
		userAuth.ExpiresIn = expiresAt.Unix()

		result = true
	} else {
		log.Fatalf("state returned doesn't match generated state: %s", state)
	}

	return result
}

// starts the application by either doing a user authorization or using data from the previous user auth
func ApplicationStart() {
	// first read the user.json
	userInfo, err := os.ReadFile("user.json")
	if err != nil {
		panic(err)
	}
	// if the user.json file hasn't been initialized
	if len(userInfo) == 0 {
		result := SignIn()
		if result {
			// userAuth contains the values we need for sending requests and we can make theme
		}
	} else {
		err := json.Unmarshal(userInfo, &userAuth)
		if err != nil {
			log.Fatal("Failed to convert userInfo to struct")
		}

		//check if the tokens have expired
		var time = time.Now()
		if userAuth.ExpiresIn < time.Unix() {
			// do revalidation

			refreshRequest := RefreshToken{
				ClientID:     "PriceChecker",
				ClientSecret: "asdasdasd", // need to create this
				GrantType:    "refresh_token",
				RefreshToken: userAuth.RefreshToken,
			}

			jsonData, err := json.Marshal(refreshRequest)
			if err != nil {
				log.Fatal("Failed to convert to json")
			}

			var refreshURL = "http://www.pathofexile.com/oauth/token"
			var contentType = "application/x-www-form-urlencoded"

			res, err := http.Post(refreshURL, contentType, bytes.NewBuffer(jsonData))
			if err != nil {
				panic(err)
			}
			body, err := io.ReadAll(res.Body)
			if err != nil {
				panic(err)
			}
			res.Body.Close()

			err = os.WriteFile("user.json", body, 0644)
			if err != nil {
				panic(err)
			}
			// save user info locally
			json.Unmarshal(body, &userAuth)

		}
	}
}

func SearchItem() {

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
