package team

import (
	"context"
	"fmt"
	"net/http"
	"slices"
	"strings"

	cloudquery_api "github.com/cloudquery/cloudquery-api-go"
	"github.com/cloudquery/cloudquery/cli/v6/internal/api"
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
		return err
	}
	return c.ValidateTeamAgainstTeams(name, teams)
}

func (*Client) ValidateTeamAgainstTeams(name string, teams []string) error {
	if slices.Contains(teams, name) {
		return nil
	}
	return fmt.Errorf("team %q not found. Teams available to you: %v", name, strings.Join(teams, ", "))
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

func (c *Client) GetTeam(ctx context.Context, team string) (*Team, error) {
	resp, err := c.api.GetTeamByNameWithResponse(ctx, team)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode() != http.StatusOK || resp.JSON200 == nil {
		return nil, fmt.Errorf("failed to get team %q: %s", team, resp.Status())
	}

	return resp.JSON200, nil
}
