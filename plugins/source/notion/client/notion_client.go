package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/cloudquery/cloudquery/plugins/source/notion/internal/databases"
	"github.com/cloudquery/cloudquery/plugins/source/notion/internal/pages"
	"github.com/cloudquery/cloudquery/plugins/source/notion/internal/users"
)

const defaultURL = "https://api.notion.com/v1"

type NotionClient struct {
	BaseURL       string
	Client        *http.Client
	AuthToken     string
	NotionVersion string
}

type Error struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func NewNotionClient(authToken string, notionVersion string) (*NotionClient, error) {
	return &NotionClient{
		BaseURL:       defaultURL,
		Client:        http.DefaultClient,
		AuthToken:     authToken,
		NotionVersion: notionVersion,
	}, nil
}

func (c *NotionClient) GetUsers(nextCursor string, hasMore bool) (*users.Users, error) {
	// Get request takes start_cursor in query string
	queryParameter := ""
	if hasMore {
		queryParameter = "?start_cursor=" + nextCursor
	}

	// Create an HTTP request to path /users
	req, err := http.NewRequest("GET", c.BaseURL+"/users"+queryParameter, nil)
	if err != nil {
		return nil, err
	}

	// Set the Bearer Token and Notion Version in the request headers
	token := "Bearer " + c.AuthToken
	notionVersion := c.NotionVersion

	req.Header.Set("Authorization", token)
	req.Header.Set("Notion-Version", notionVersion)

	// Send the HTTP request
	r, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}

	defer r.Body.Close()

	// Check the HTTP status code for errors
	if r.StatusCode == http.StatusOK {
		u := &users.Users{}
		if err := json.NewDecoder(r.Body).Decode(u); err != nil {
			return nil, err
		}
		return u, nil
	}

	e := &Error{}
	json.NewDecoder(r.Body).Decode(e)
	return nil, fmt.Errorf("Status :"+strconv.Itoa(e.Status)+", Message :"+e.Message, r.StatusCode)
}

func (c *NotionClient) GetPages(nextCursor string, hasMore bool) (*pages.Pages, error) {
	// Post request takes the queryParameter in reqest body
	queryParameter := ""
	if hasMore {
		queryParameter = fmt.Sprintf(`, "start_cursor" : "%s"`, nextCursor)
	}

	// Define the request body as a []byte. Max paze_size is 100 according to notion api docs.
	var reqBody = []byte(fmt.Sprintf(`{
		"filter": {
			"value": "page",
			"property": "object"
		},
		"sort": {
			"direction": "ascending",
			"timestamp": "last_edited_time"
		},
		"page_size": 100%s
	}`, queryParameter))

	// Create an HTTP request to path /search
	req, err := http.NewRequest("POST", c.BaseURL+"/search", bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}

	// Set the Bearer Token, Notion Version and Content Type in the request headers
	token := "Bearer " + c.AuthToken
	notionVersion := c.NotionVersion
	contentType := "application/json"

	req.Header.Set("Authorization", token)
	req.Header.Set("Notion-Version", notionVersion)
	req.Header.Set("Content-Type", contentType)

	// Send the HTTP request
	r, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}

	defer r.Body.Close()

	// Check the HTTP status code for errors
	if r.StatusCode == http.StatusOK {
		p := &pages.Pages{}
		if err := json.NewDecoder(r.Body).Decode(p); err != nil {
			return nil, err
		}

		return p, nil
	}

	e := &Error{}
	json.NewDecoder(r.Body).Decode(e)
	return nil, fmt.Errorf("Status :"+strconv.Itoa(e.Status)+", Message :"+e.Message, r.StatusCode)
}

func (c *NotionClient) GetDatabases(nextCursor string, hasMore bool) (*databases.Databases, error) {
	// Post request takes the queryParameter in reqest body
	queryParameter := ""
	if hasMore {
		queryParameter = fmt.Sprintf(`, "start_cursor" : "%s"`, nextCursor)
	}

	// Define the request body as a []byte. Max paze_size is 100 according to notion api docs.
	var reqBody = []byte(fmt.Sprintf(`{
		"filter": {
			"value": "database",
			"property": "object"
		},
		"sort": {
			"direction": "ascending",
			"timestamp": "last_edited_time"
		},
		"page_size": 100%s
	}`, queryParameter))

	// Create an HTTP request to path /search
	req, err := http.NewRequest("POST", c.BaseURL+"/search", bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}

	// Set the Bearer Token, Notion Version and Content Type in the request headers
	token := "Bearer " + c.AuthToken
	notionVersion := c.NotionVersion
	contentType := "application/json"

	req.Header.Set("Authorization", token)
	req.Header.Set("Notion-Version", notionVersion)
	req.Header.Set("Content-Type", contentType)

	// Send the HTTP request
	r, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()

	// Check the HTTP status code for errors
	if r.StatusCode == http.StatusOK {
		d := &databases.Databases{}
		if err := json.NewDecoder(r.Body).Decode(d); err != nil {
			return nil, err
		}

		return d, nil
	}

	e := &Error{}
	json.NewDecoder(r.Body).Decode(e)
	return nil, fmt.Errorf("Status :"+strconv.Itoa(e.Status)+", Message :"+e.Message, r.StatusCode)
}
