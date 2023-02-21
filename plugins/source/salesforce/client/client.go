package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/cloudquery/plugin-sdk/plugins/source"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/specs"
	"github.com/rs/zerolog"
)

const ApiVersion = "v56.0"

const defaultHTTPTimeout = 30 * time.Second

type Client struct {
	logger              zerolog.Logger
	pluginSpec          Spec
	LoginResponse       LoginResponse
	ListObjectsResponse ListObjectsResponse
	Object              string
	Client              *http.Client
	HTTPDataEndpoint    string
}

type LoginRequest struct {
	GrantType    string `json:"grant_type"`
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	Username     string `json:"username"`
	Password     string `json:"password"`
}

type LoginResponse struct {
	AccessToken string `json:"access_token"`
	InstanceUrl string `json:"instance_url"`
	Id          string `json:"id"`
	TokenType   string `json:"token_type"`
	IssuedAt    string `json:"issued_at"`
	Signature   string `json:"signature"`
}

type Sobject struct {
	Name string `json:"name"`
}

type ListObjectsResponse struct {
	Sobject []Sobject `json:"sobjects"`
}

func Configure(ctx context.Context, logger zerolog.Logger, spec specs.Source, _ source.Options) (schema.ClientMeta, error) {
	cqClient := Client{
		logger: logger,
	}
	var sfSpec Spec
	if err := spec.UnmarshalSpec(&sfSpec); err != nil {
		return nil, err
	}
	sfSpec.SetDefaults()
	if err := sfSpec.Validate(); err != nil {
		return nil, err
	}
	cqClient.Client = &http.Client{
		Timeout: time.Duration(30) * time.Second,
	}
	cqClient.pluginSpec = sfSpec
	if err := cqClient.login(ctx); err != nil {
		return nil, err
	}
	cqClient.HTTPDataEndpoint = cqClient.LoginResponse.InstanceUrl + "/services/data/" + ApiVersion

	if err := cqClient.listObjects(ctx); err != nil {
		return nil, err
	}
	return &cqClient, nil
}

func (c *Client) Logger() *zerolog.Logger {
	return &c.logger
}

func (c *Client) ID() string {
	return c.Object
}

func (c *Client) listObjects(ctx context.Context) error {
	request, err := http.NewRequest("GET", c.LoginResponse.InstanceUrl+"/services/data/v56.0/sobjects", nil)
	if err != nil {
		return err
	}
	request = request.WithContext(ctx)

	// Set the content type header
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer "+c.LoginResponse.AccessToken)

	// Create an http.Client and POST the request
	client := &http.Client{
		// login timeout
		Timeout: defaultHTTPTimeout,
	}
	response, err := client.Do(request)
	if err != nil {
		return err
	}

	// Read the response body
	defer response.Body.Close()
	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}
	var listObjectsResponse ListObjectsResponse
	err = json.Unmarshal(responseBody, &listObjectsResponse)
	if err != nil {
		return err
	}
	c.ListObjectsResponse = listObjectsResponse

	return nil
}

func (c *Client) login(ctx context.Context) error {
	data := url.Values{
		"grant_type":    {"password"},
		"client_id":     {c.pluginSpec.ClientId},
		"client_secret": {c.pluginSpec.ClientSecret},
		"username":      {c.pluginSpec.Username},
		"password":      {c.pluginSpec.Password},
	}

	request, err := http.NewRequestWithContext(ctx, "POST", "https://login.salesforce.com/services/oauth2/token", bytes.NewBufferString(data.Encode()))
	if err != nil {
		return err
	}

	// Set the content type header
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// Create an http.Client and POST the request
	client := &http.Client{
		// login timeout
		Timeout: 20 * time.Second,
	}
	response, err := client.Do(request)
	if err != nil {
		return err
	}

	// Read the response body
	defer response.Body.Close()
	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to login: %s", string(responseBody))
	}

	err = json.Unmarshal(responseBody, &c.LoginResponse)
	if err != nil {
		return err
	}

	return nil
}
