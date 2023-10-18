package team

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	cloudquery_api "github.com/cloudquery/cloudquery-api-go"
)

type Client struct {
	url string
	api *cloudquery_api.ClientWithResponses
}

func NewClient(url string, token string) (*Client, error) {
	cl, err := cloudquery_api.NewClientWithResponses(url, cloudquery_api.WithRequestEditorFn(func(ctx context.Context, req *http.Request) error {
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
		return nil
	}))
	if err != nil {
		return nil, err
	}
	return &Client{
		url: url,
		api: cl,
	}, nil
}

func (c *Client) ValidateTeam(ctx context.Context, teamName string) error {
	teams, err := c.ListAllTeams(ctx)
	if err != nil {
		return err
	}
	for _, team := range teams {
		if team == teamName {
			return nil
		}
	}
	return fmt.Errorf("team %q not found. Teams available to you: %v", teamName, strings.Join(teams, ", "))
}

func (c *Client) ListAllTeams(ctx context.Context) ([]string, error) {
	page := cloudquery_api.Page(1)
	perPage := cloudquery_api.PerPage(100)
	teams := make([]string, 0)
	for {
		resp, err := c.api.ListTeamsWithResponse(ctx, &cloudquery_api.ListTeamsParams{
			PerPage: &perPage,
			Page:    &page,
		})
		if err != nil {
			return nil, err
		}
		if resp.StatusCode() != http.StatusOK || resp.JSON200 == nil {
			return nil, fmt.Errorf("failed to list teams: %s", resp.Status())
		}
		for _, team := range resp.JSON200.Items {
			teams = append(teams, team.Name)
		}
		if resp.JSON200.Metadata.LastPage == nil || *resp.JSON200.Metadata.LastPage <= int(page) {
			break
		}
		page++
	}
	return teams, nil
}
