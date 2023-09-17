package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

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

	u := &users.Users{}
	if err := json.NewDecoder(r.Body).Decode(u); err != nil {
		return nil, err
	}
	return u, nil
}

func (c *NotionClient) GetPages(nextCursor string, hasMore bool) (*pages.Pages, error) {
	// Post request takes the queryParameter in reqest body
	queryParameter := ""
	if hasMore {
		queryParameter = fmt.Sprintf(`, "start_cursor" : "%s"`, nextCursor)
	}

	// Define the request body as a []byte
	var reqBody = []byte(fmt.Sprintf(`{
		"filter": {
			"value": "page",
			"property": "object"
		},
		"sort": {
			"direction": "ascending",
			"timestamp": "last_edited_time"
		},
		"page_size": 1%s
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

	p := &pages.Pages{}
	if err := json.NewDecoder(r.Body).Decode(p); err != nil {
		return nil, err
	}

	return p, nil
}

func (c *NotionClient) GetDatabases(nextCursor string, hasMore bool) (*databases.Databases, error) {
	// Post request takes the queryParameter in reqest body
	queryParameter := ""
	if hasMore {
		queryParameter = fmt.Sprintf(`, "start_cursor" : "%s"`, nextCursor)
	}

	// Define the request body as a []byte
	var reqBody = []byte(fmt.Sprintf(`{
		"filter": {
			"value": "database",
			"property": "object"
		},
		"sort": {
			"direction": "ascending",
			"timestamp": "last_edited_time"
		},
		"page_size": 1%s
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

	d := &databases.Databases{}
	if err := json.NewDecoder(r.Body).Decode(d); err != nil {
		return nil, err
	}

	return d, nil
}
