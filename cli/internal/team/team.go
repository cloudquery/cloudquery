package team

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	cloudquery_api "github.com/cloudquery/cloudquery-api-go"
	"github.com/cloudquery/cloudquery/cli/v6/internal/api"
	"github.com/cloudquery/cloudquery/cli/v6/internal/hub"
)

type Team = cloudquery_api.Team

type Client struct {
	api *cloudquery_api.ClientWithResponses
}

func NewClient(token string) (*Client, error) {
	cl, err := api.NewClient(token)
	if err != nil {
		return nil, err
	}
	return &Client{
		api: cl,
	}, nil
}

func NewClientFromAPI(apiClient *cloudquery_api.ClientWithResponses) *Client {
	return &Client{
		api: apiClient,
	}
}

func (c *Client) ValidateTeam(ctx context.Context, name string) error {
	teams, err := c.ListAllTeams(ctx)
	if err != nil {
		return fmt.Errorf("failed to list teams: %w", err)
	}
	_, err = c.FindTeam(teams, name)
	return err
}

func (*Client) FindTeam(teams []Team, name string) (*Team, error) {
	for _, t := range teams {
		if t.Name == name {
			return &t, nil
		}
	}
	teamNames := make([]string, 0, len(teams))
	for _, t := range teams {
		teamNames = append(teamNames, t.Name)
	}
	return nil, fmt.Errorf("team %q not found. Teams available to you: %v", name, strings.Join(teamNames, ", "))
}

func (c *Client) ListAllTeams(ctx context.Context) ([]Team, error) {
	page := cloudquery_api.Page(1)
	perPage := cloudquery_api.PerPage(100)
	teams := make([]Team, 0)
	for {
		resp, err := c.api.ListTeamsWithResponse(ctx, &cloudquery_api.ListTeamsParams{
			PerPage: &perPage,
			Page:    &page,
		})
		if err != nil {
			return nil, err
		}
		if resp.StatusCode() != http.StatusOK || resp.JSON200 == nil {
			return nil, hub.ErrorFromHTTPResponse(resp.HTTPResponse, resp)
		}
		teams = append(teams, resp.JSON200.Items...)
		if resp.JSON200.Metadata.LastPage == nil || *resp.JSON200.Metadata.LastPage <= int(page) {
			break
		}
		page++
	}
	return teams, nil
}

func (c *Client) ListAllTeamNames(ctx context.Context) ([]string, error) {
	teams, err := c.ListAllTeams(ctx)
	if err != nil {
		return nil, err
	}
	names := make([]string, 0, len(teams))
	for _, t := range teams {
		names = append(names, t.Name)
	}
	return names, nil
}

func (c *Client) GetTeam(ctx context.Context, team string) (*Team, error) {
	resp, err := c.api.GetTeamByNameWithResponse(ctx, team)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode() != http.StatusOK || resp.JSON200 == nil {
		return nil, fmt.Errorf("failed to get team %q: %w", team, hub.ErrorFromHTTPResponse(resp.HTTPResponse, resp))
	}

	return resp.JSON200, nil
}

func Names(teams []Team) []string {
	names := make([]string, len(teams))
	for i, t := range teams {
		names[i] = t.Name
	}
	return names
}
