package cloudapi

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	cloudquery_api "github.com/cloudquery/cloudquery-api-go"
	"github.com/rs/zerolog"
	"github.com/samber/lo"
)

type Client struct {
	client *cloudquery_api.ClientWithResponses
}

func MustNewClient(logger zerolog.Logger) *Client {
	client, err := cloudquery_api.NewClientWithResponses("https://api.cloudquery.io")
	if err != nil {
		logger.Fatal().Err(err).Msg("Failed to create CloudQuery API client")
	}
	return &Client{client: client}
}

func (c *Client) ListAllPlugins(ctx context.Context) ([]cloudquery_api.ListPlugin, error) {
	page := cloudquery_api.Page(1)
	perPage := cloudquery_api.PerPage(100)
	plugins := make([]cloudquery_api.ListPlugin, 0)
	for {
		resp, err := c.client.ListPluginsWithResponse(ctx, &cloudquery_api.ListPluginsParams{
			PerPage: &perPage,
			Page:    &page,
		})
		if err != nil {
			return nil, fmt.Errorf("failed to list plugins: %w", err)
		}
		if resp.StatusCode() != http.StatusOK || resp.JSON200 == nil {
			return nil, fmt.Errorf("failed to list plugins: %s", resp.Status())
		}
		plugins = append(plugins, resp.JSON200.Items...)
		if resp.JSON200.Metadata.LastPage == nil || *resp.JSON200.Metadata.LastPage <= int(page) {
			break
		}
		page++
	}
	return plugins, nil
}

func (c *Client) GetLatestVersion(ctx context.Context, plugin string, kind cloudquery_api.PluginKind) (string, error) {
	// fmt.Printf("[DEBUG] Getting latest version for plugin: %s (kind: %s)\n", plugin, kind)

	resp, err := c.client.GetPluginWithResponse(ctx, "cloudquery", kind, plugin)
	if err != nil {
		return "", fmt.Errorf("failed to get plugin: %w", err)
	}

	// fmt.Printf("[DEBUG] GetPlugin API Response Status: %s\n", resp.Status())

	if resp.StatusCode() != http.StatusOK {
		// fmt.Printf("[DEBUG] Error response body: %s\n", string(resp.Body))
		return "", fmt.Errorf("failed to get plugin: %s (status: %d)", resp.Status(), resp.StatusCode())
	}

	if resp.JSON200 == nil {
		return "", errors.New("failed to get plugin: empty response body")
	}

	version := *resp.JSON200.LatestVersion
	// fmt.Printf("[DEBUG] Latest version for %s (%s): %s\n", plugin, kind, version)
	return version, nil
}

func (c *Client) GetPluginTables(ctx context.Context, plugin, version string) ([]string, error) {
	page := cloudquery_api.Page(1)
	perPage := cloudquery_api.PerPage(100)

	var tables []string
	for {
		opts := &cloudquery_api.ListPluginVersionTablesParams{
			Page:    &page,
			PerPage: &perPage,
		}
		resp, err := c.client.ListPluginVersionTablesWithResponse(ctx, "cloudquery", cloudquery_api.PluginKindSource, plugin, version, opts)
		if err != nil {
			return nil, fmt.Errorf("failed to list plugin tables: %w", err)
		}
		if resp.StatusCode() != http.StatusOK || resp.JSON200 == nil {
			return nil, fmt.Errorf("failed to list plugin tables: %s", resp.Status())
		}

		for i := range resp.JSON200.Items {
			tables = append(tables, resp.JSON200.Items[i].Name)
		}
		if resp.JSON200.Metadata.LastPage == nil || *resp.JSON200.Metadata.LastPage <= int(page) {
			break
		}
		page++
	}

	return tables, nil
}

func (c *Client) GetPluginDocs(ctx context.Context, plugin, kind, version string) (map[string]string, error) {
	// fmt.Printf("[DEBUG] Getting docs for plugin: %s, kind: %s, version: %s\n", plugin, kind, version)

	resp, err := c.client.ListPluginVersionDocsWithResponse(ctx, "cloudquery", cloudquery_api.PluginKind(kind), plugin, version, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get docs: %w", err)
	}

	// fmt.Printf("[DEBUG] API Response Status: %s\n", resp.Status())

	if resp.StatusCode() != http.StatusOK {
		// fmt.Printf("[DEBUG] Error response body: %s\n", string(resp.Body))
		return nil, fmt.Errorf("failed to get docs: %s (status: %d)", resp.Status(), resp.StatusCode())
	}

	if resp.JSON200 == nil {
		return nil, errors.New("failed to get docs: empty response body")
	}

	// fmt.Printf("[DEBUG] Successfully retrieved %d documentation pages\n", len(resp.JSON200.Items))

	docs := lo.Reduce(resp.JSON200.Items, func(agg map[string]string, item cloudquery_api.PluginDocsPage, _ int) map[string]string {
		agg[item.Name] = item.Content
		return agg
	}, make(map[string]string))

	return docs, nil
}

func (c *Client) GetPluginTableSchema(ctx context.Context, plugin, version, table string) (cloudquery_api.PluginTableDetails, error) {
	resp, err := c.client.GetPluginVersionTableWithResponse(ctx, "cloudquery", cloudquery_api.PluginKindSource, plugin, version, table)
	if err != nil {
		return cloudquery_api.PluginTableDetails{}, fmt.Errorf("failed to get table schema: %w", err)
	}
	if resp.StatusCode() != http.StatusOK || resp.JSON200 == nil {
		return cloudquery_api.PluginTableDetails{}, fmt.Errorf("failed to get table schema: %s", resp.Status())
	}
	return *resp.JSON200, nil
}
