package rest

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
)

type FacebookClient struct {
	AdAccountId string
	AccessToken string
	httpClient  *http.Client
}

type Paging struct {
	Cursors *Cursors `json:"cursors"`
	Next    string   `json:"next"`
}

type Cursors struct {
	Before string `json:"before"`
	After  string `json:"after"`
}

type FacebookErrorResponse struct {
	FacebookError FacebookError `json:"error"`
}

type FacebookError struct {
	Type    string `json:"type"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func NewFacebookClient(httpClient *http.Client, adAccountId string, accessToken string) *FacebookClient {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	return &FacebookClient{
		AdAccountId: adAccountId,
		AccessToken: accessToken,
		httpClient:  httpClient,
	}
}

// Returns the json tag of all fields of any struct.
func getAllFieldJsonTags(s any) []string {
	reflectType := reflect.TypeOf(s)

	jsonTags := make([]string, reflectType.NumField())

	for i := 0; i < reflectType.NumField(); i++ {
		jsonTag := reflectType.Field(i).Tag.Get("json")
		if jsonTag == "" {
			continue
		}

		jsonTags[i] = jsonTag
	}

	return jsonTags
}

// Receieves an unsuccessful http response from facebook, and returns a golang error with fmt.Errorf().
func httpErrorToGolangError(response *http.Response) error {
	var facebookErrorResponse FacebookErrorResponse
	err := json.NewDecoder(response.Body).Decode(&facebookErrorResponse)
	if err != nil {
		return fmt.Errorf("http status %v", response.StatusCode)
	}

	facebookError := facebookErrorResponse.FacebookError

	return fmt.Errorf("http status %v: (%v) %v: %v",
		response.StatusCode,
		facebookError.Code,
		facebookError.Type,
		facebookError.Message)
}

// http.Do returns an error that contains the http url. Because the url contains the access token,
// we need to sanitize it.
func sanitizeUrlError(err error) error {
	if err == nil {
		return nil
	}

	var urlError *url.Error
	if errors.As(err, &urlError) {
		urlError.URL = ""
		return urlError
	}

	return err
}
