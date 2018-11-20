package nibestats

import (
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"net/url"
)

const redirectURI = "https://nibe.jl-media.se/oauth/callback"
const clientID = ""
const clientSecret = ""

func RedirectToAuthenticationProviderHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, fmt.Sprintf("%s/oauth/authorize?response_type=code&client_id=%s&scope=%s&redirect_uri=%s&state=%s",
		NIBE_UPLINK_API,
		clientID,
		"READSYSTEM",
		redirectURI,
		"?",
	), 302)
}

// OAuthCallbackHandler - Method for reading oauth callback
func OAuthCallbackHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	resp, _ := http.PostForm(fmt.Sprintf("%s/oauth/token", NIBE_UPLINK_API),
		url.Values{
			"grant_type": {"authorization_code"},
			"client_id": {clientID},
			"client_secret": {clientSecret},
			"code": {params["code"]},
			"scope": {"READSYSTEM"},
			"redirect_uri": {redirectURI},
		})
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
}
