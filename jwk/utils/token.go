package utils

import (
	"encoding/json"
	"fmt"
	"github.com/auth0-community/go-auth0"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

const (
	//free trier account:  https://auth0.com/signup?&signUpData=%7B%22category%22%3A%22docs%22%7D
	AUTH0_DOMAIN   = "https://newbmiao.auth0.com"
	AUTH0_AUDIENCE = "https://newbmiao.auth0.com/api/v2/"
	CLIENT_ID      = "2QMsBU6bru7gtSR3k3S39LVVZLLh12o2"
	CLIENT_SECRET  = "k4DTlMmBKgau0X6Rge6VbwMQKC2Ewq-RSURPTd59ivLFT5zA-jXGLxQxLsOIDNN4"
	KID            = "QjY4MzJBMzc5NkQxODU0OTJGOERFRENBMjEyNjA2RjFERjNGQTQ1RA"
)

type tokenType struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	Scope       string `json:"scope"`
	TokenType   string `json:"token_type"`
}

func GetToken() string {

	var token tokenType
	url := AUTH0_DOMAIN + "/oauth/token"

	payload := strings.NewReader("{\"client_id\":\"" + CLIENT_ID + "\",\"client_secret\":\"" + CLIENT_SECRET +
		"\",\"audience\":\"" + AUTH0_AUDIENCE + "\",\"grant_type\":\"client_credentials\"}")

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("content-type", "application/json")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	json.Unmarshal(body, &token)
	fmt.Println("get token:", token.AccessToken)
	return token.AccessToken
}

//the way different from server.go -> getPemCert()
func GetPublicKey() interface{} {
	fmt.Println("get public key from JWK_URI")
	opts := auth0.JWKClientOptions{URI: AUTH0_DOMAIN + "/.well-known/jwks.json"}
	// Creating key cacher with max age of 100sec and max size of 5 entries.
	// Defaults to persistent key cacher if not specified when creating a client.
	keyCacher := auth0.NewMemoryKeyCacher(time.Duration(100)*time.Second, 5)
	client := auth0.NewJWKClientWithCache(opts, nil, keyCacher)

	searchedKey, err := client.GetKey(KID)

	if err != nil {
		fmt.Println("Cannot get key because of", err)
		return nil
	}
	fmt.Println("key:", searchedKey)
	return searchedKey.Certificates[0].PublicKey
}
