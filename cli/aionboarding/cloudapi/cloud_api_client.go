package cloudapi

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"regexp"

	cloudquery_api "github.com/cloudquery/cloudquery-api-go"
	"github.com/rs/zerolog"
	"github.com/samber/lo"
)

type cloudAPIImpl struct {
	client       *cloudquery_api.ClientWithResponses
	logger       zerolog.Logger
	cacheManager *CacheManager
}

func NewCloudAPI(logger zerolog.Logger) CloudAPI {
	client, err := cloudquery_api.NewClientWithResponses("https://api.cloudquery.io")
	if err != nil {
		logger.Fatal().Err(err).Msg("Failed to create CloudQuery API client")
	}

	cacheManager, err := NewCacheManager(logger)
	if err != nil {
		logger.Fatal().Err(err).Msg("Failed to create cache manager")
	}

	impl := &cloudAPIImpl{
		client:       client,
		logger:       logger,
		cacheManager: cacheManager,
	}

	// Start background initialization
	go impl.initializeBackground()

	return impl
}

func (c *cloudAPIImpl) initializeBackground() {
	ctx := context.Background()

	// List all plugins
	_, err := c.ListAllPlugins(ctx)
	if err != nil {
		c.logger.Error().Err(err).Msg("Failed to initialize plugins list")
	}

	// Get plugin data into cache for AWS source
	_, awsVersion, err := c.GetPluginData(ctx, "aws", "source")
	if err != nil {
		c.logger.Error().Err(err).Msg("Failed to initialize AWS plugin data")
	}

	// Get plugin data into cache for postgres destination
	_, _, err = c.GetPluginData(ctx, "postgresql", "destination")
	if err != nil {
		c.logger.Error().Err(err).Msg("Failed to initialize PostgreSQL plugin data")
	}

	// Get source plugin tables into cache for AWS with * regex
	_, err = c.GetSourcePluginTables(ctx, "aws", awsVersion, "*")
	if err != nil {
		c.logger.Error().Err(err).Msg("Failed to initialize AWS plugin tables")
	}
}

func (c *cloudAPIImpl) ListAllPlugins(ctx context.Context) ([]string, error) {
	cacheKey := "list_all_plugins"
	if cached, exists := c.cacheManager.Get(cacheKey); exists {
		return cached.([]string), nil
	}

	plugins, err := c.listAllPlugins(ctx)
	if err != nil {
		return nil, err
	}

	pluginNames := make([]string, len(plugins))
	for i, plugin := range plugins {
		pluginNames[i] = fmt.Sprintf("%s (%s)", plugin.Name, plugin.Kind)
	}

	c.cacheManager.Set(cacheKey, pluginNames)
	return pluginNames, nil
}

func (c *cloudAPIImpl) listAllPlugins(ctx context.Context) ([]cloudquery_api.ListPlugin, error) {
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

func (c *cloudAPIImpl) GetPluginData(ctx context.Context, pluginName, pluginKind string) (string, string, error) {
	cacheKey := "plugin_data_" + pluginName + "_" + pluginKind
	if cached, exists := c.cacheManager.Get(cacheKey); exists {
		data := cached.(struct {
			docs    string
			version string
		})
		return data.docs, data.version, nil
	}

	// Get latest version
	kind := cloudquery_api.PluginKind(pluginKind)
	version, err := c.getLatestVersion(ctx, pluginName, kind)
	if err != nil {
		return "", "", err
	}

	// Get plugin docs
	docs, err := c.getPluginDocs(ctx, pluginName, pluginKind, version)
	if err != nil {
		return "", "", err
	}

	// Convert docs map to string
	docsStr := ""
	for name, content := range docs {
		docsStr += "=== " + name + " ===\n" + content + "\n\n"
	}

	result := struct {
		docs    string
		version string
	}{
		docs:    docsStr,
		version: version,
	}

	c.cacheManager.Set(cacheKey, result)
	return docsStr, version, nil
}

func (c *cloudAPIImpl) GetSourcePluginTables(ctx context.Context, pluginName, version, tablesRegex string) ([]string, error) {
	// If version is empty, get latest version
	if version == "" {
		var err error
		version, err = c.getLatestVersion(ctx, pluginName, cloudquery_api.PluginKindSource)
		if err != nil {
			return nil, err
		}
	}

	// Cache key for all tables (without regex)
	cacheKey := "source_plugin_tables_" + pluginName + "_" + version
	var allTables []string
	var err error

	// Try to get from cache first
	if cached, exists := c.cacheManager.Get(cacheKey); exists {
		allTables = cached.([]string)
	} else {
		// Get all tables for the plugin
		allTables, err = c.getPluginTables(ctx, pluginName, version)
		if err != nil {
			return nil, err
		}
		// Cache all tables
		c.cacheManager.Set(cacheKey, allTables)
	}

	// Apply regex filter to cached/fetched tables
	var filteredTables []string
	if tablesRegex == "*" {
		filteredTables = allTables
	} else {
		regex, err := regexp.Compile(tablesRegex)
		if err != nil {
			return nil, err
		}

		for _, table := range allTables {
			if regex.MatchString(table) {
				filteredTables = append(filteredTables, table)
			}
		}
	}

	return filteredTables, nil
}

func (c *cloudAPIImpl) GetConversationID(ctx context.Context, userID string) (string, error) {
	conversationID, exists := c.cacheManager.GetConversationID(userID)
	if !exists {
		return "", nil
	}
	return conversationID, nil
}

func (c *cloudAPIImpl) SetConversationID(ctx context.Context, userID string, conversationID string) error {
	c.cacheManager.SetConversationID(userID, conversationID)
	return nil
}

func (c *cloudAPIImpl) GetTokenUsage(ctx context.Context, teamName string) (int, error) {
	tokenUsage, exists := c.cacheManager.GetTokenUsage(teamName)
	if !exists {
		return 0, nil
	}
	return tokenUsage, nil
}

func (c *cloudAPIImpl) SetTokenUsage(ctx context.Context, teamName string, tokenUsage int) error {
	c.cacheManager.SetTokenUsage(teamName, tokenUsage)
	return nil
}

func (c *cloudAPIImpl) getLatestVersion(ctx context.Context, plugin string, kind cloudquery_api.PluginKind) (string, error) {
	c.logger.Debug().Str("plugin", plugin).Str("kind", string(kind)).Msg("Getting latest version for plugin")

	resp, err := c.client.GetPluginWithResponse(ctx, "cloudquery", kind, plugin)
	if err != nil {
		return "", fmt.Errorf("failed to get plugin: %w", err)
	}

	c.logger.Debug().Str("status", resp.Status()).Msg("GetPlugin API Response Status")

	if resp.StatusCode() != http.StatusOK {
		c.logger.Debug().Str("body", string(resp.Body)).Msg("Error response body")
		return "", fmt.Errorf("failed to get plugin: %s (status: %d)", resp.Status(), resp.StatusCode())
	}

	if resp.JSON200 == nil {
		return "", errors.New("failed to get plugin: empty response body")
	}

	version := *resp.JSON200.LatestVersion
	c.logger.Debug().Str("plugin", plugin).Str("kind", string(kind)).Str("version", version).Msg("Latest version retrieved")
	return version, nil
}

func (c *cloudAPIImpl) getPluginTables(ctx context.Context, plugin, version string) ([]string, error) {
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

func (c *cloudAPIImpl) getPluginDocs(ctx context.Context, plugin, kind, version string) (map[string]string, error) {
	c.logger.Debug().Str("plugin", plugin).Str("kind", kind).Str("version", version).Msg("Getting docs for plugin")

	resp, err := c.client.ListPluginVersionDocsWithResponse(ctx, "cloudquery", cloudquery_api.PluginKind(kind), plugin, version, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get docs: %w", err)
	}

	c.logger.Debug().Str("status", resp.Status()).Msg("API Response Status")

	if resp.StatusCode() != http.StatusOK {
		c.logger.Debug().Str("body", string(resp.Body)).Msg("Error response body")
		return nil, fmt.Errorf("failed to get docs: %s (status: %d)", resp.Status(), resp.StatusCode())
	}

	if resp.JSON200 == nil {
		return nil, errors.New("failed to get docs: empty response body")
	}

	c.logger.Debug().Int("count", len(resp.JSON200.Items)).Msg("Successfully retrieved documentation pages")

	docs := lo.Reduce(resp.JSON200.Items, func(agg map[string]string, item cloudquery_api.PluginDocsPage, _ int) map[string]string {
		agg[item.Name] = item.Content
		return agg
	}, make(map[string]string))

	return docs, nil
}
