package services

import (
	"bytes"
	"context"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/cloudquery/cloudquery/plugins/source/salesforce/client"
)

const (
	// job states
	StateJobComplete    = "JobComplete"
	StateInProgress     = "InProgress"
	StateUploadComplete = "UploadComplete"
	maxWaitForJob       = 5 * time.Minute
)

type createQueryJobRequest struct {
	Operation       string `json:"operation"`
	Query           string `json:"query"`
	ContentType     string `json:"contentType"`
	ColumnDelimiter string `json:"columnDelimiter"`
	LineEnding      string `json:"lineEnding"`
}

type createQueryJobResponse struct {
	Id    string `json:"id"`
	State string `json:"state"`
}

type getQueryJobStatusResponse struct {
	Id    string `json:"id"`
	State string `json:"state"`
}

func fetchJobResults(ctx context.Context, c *client.Client, jobId string, fields []string, res chan<- any) error {
	var queryJobStatusRes getQueryJobStatusResponse
	sleepInterval := time.Second
	totalWait := time.Duration(0)
	url := c.HTTPDataEndpoint + "/jobs/query/" + jobId
	for {
		if totalWait > maxWaitForJob {
			return fmt.Errorf("job %s timed out", jobId)
		}
		if err := c.Get(ctx, url, &queryJobStatusRes); err != nil {
			return err
		}
		if queryJobStatusRes.State == StateJobComplete {
			break
		} else if queryJobStatusRes.State == StateInProgress || queryJobStatusRes.State == StateUploadComplete {
			time.Sleep(sleepInterval)
			sleepInterval *= 2
			totalWait += sleepInterval
			continue
		}
		return fmt.Errorf("job state is %s", queryJobStatusRes.State)
	}

	url = c.HTTPDataEndpoint + "/jobs/query/" + jobId + "/results?maxRecords=10000"
	for {
		request, err := http.NewRequestWithContext(ctx, "GET", url, nil)
		if err != nil {
			return err
		}

		request.Header.Set("Accept", "text/csv")
		request.Header.Set("Authorization", "Bearer "+c.LoginResponse.AccessToken)

		response, err := c.Client.Do(request)
		if err != nil {
			return err
		}
		body, err := io.ReadAll(response.Body)
		if err != nil {
			response.Body.Close()
			return err
		}
		response.Body.Close()

		if response.StatusCode != http.StatusOK {
			return fmt.Errorf("url: %s. returned: %d with body: %s", url, response.StatusCode, string(body))
		}
		jsonresults := make([]map[string]any, 0)
		reader := csv.NewReader(bytes.NewReader(body))
		// read header
		_, err = reader.Read()
		if err != nil {
			return err
		}

		for {
			record, err := reader.Read()
			if err == io.EOF {
				break
			}
			if err != nil {
				return err
			}
			obj := make(map[string]any)
			for i, field := range fields {
				obj[field] = record[i]
			}
			jsonresults = append(jsonresults, obj)
		}
		res <- jsonresults
		locator := response.Header.Get("Sforce-Locator")
		if locator == "" || locator == "null" {
			break
		}
		url = c.HTTPDataEndpoint + "/jobs/query/" + jobId + "/results?maxRecords=10000&locator+" + locator
	}

	return nil
}

func createQueryJob(ctx context.Context, c *client.Client, fields []string) (string, error) {
	var query strings.Builder
	query.WriteString("SELECT ")
	query.WriteString(strings.Join(fields, ","))
	query.WriteString(" FROM ")
	query.WriteString(c.Object)
	req := createQueryJobRequest{
		Operation:       "query",
		Query:           query.String(),
		ContentType:     "CSV",
		ColumnDelimiter: "COMMA",
		LineEnding:      "LF",
	}
	reqBytes, err := json.Marshal(req)
	if err != nil {
		return "", err
	}

	url := c.HTTPDataEndpoint + "/jobs/query"
	request, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewReader(reqBytes))
	if err != nil {
		return "", err
	}
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Authorization", "Bearer "+c.LoginResponse.AccessToken)

	response, err := c.Client.Do(request)
	if err != nil {
		return "", err
	}

	// Read the response body
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	if response.StatusCode != http.StatusOK {
		return "", fmt.Errorf("url: %s. returned: %d with body: %s", url, response.StatusCode, string(body))
	}
	var createQueryJobRes createQueryJobResponse
	if err := json.Unmarshal(body, &createQueryJobRes); err != nil {
		return "", err
	}
	return createQueryJobRes.Id, nil
}
