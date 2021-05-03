package gcr

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	_ "github.com/aditmeno/registry-credential-helper/interface"
)

type GCRCredentialHelper struct {
	Token string
}

type responseJSON struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	TokenType   string `json:"token_type"`
}

func GetGCRCredentialHelper() *GCRCredentialHelper {
	return &GCRCredentialHelper{}
}

func (credHelper *GCRCredentialHelper) Login() {
	// Dummy function does nothing
	login()

}

func (credHelper *GCRCredentialHelper) GetToken() string {
	credHelper.Token = getToken()
	return credHelper.Token
}

func login() {
	// Dummy function kept to be complaint in code format with ECR
}

func getToken() string {
	metadataURL := "http://metadata.google.internal/computeMetadata/v1/instance/service-accounts/default/token"
	client := &http.Client{}
	req, _ := http.NewRequest("GET", metadataURL, nil)
	req.Header.Set("Metadata-Flavor", "Google")
	res, err := client.Do(req)
	if err != nil {
		log.Fatalf("There was an error getting the respose: %s", err.Error())
	}
	defer res.Body.Close()
	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatalf("Error getting auth token: %s", readErr.Error())
	}
	responseJSONContent := responseJSON{}
	jsonErr := json.Unmarshal(body, &responseJSONContent)
	if jsonErr != nil {
		log.Fatalf("Error unmarshalling json: %s", jsonErr.Error())
	}

	return responseJSONContent.AccessToken
}
