package pendo

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"

	"github.com/rs/zerolog"
)

const (
	// BaseApiUrl is the base URL for the Pendo API.
	BaseApiUrl         = "https://app.pendo.io/api/v1"
	PagesEndpoint      = BaseApiUrl + "/page" + expand
	FeaturesEndpoint   = BaseApiUrl + "/feature" + expand
	TrackTypesEndpoint = BaseApiUrl + "/tracktype" + expand
	GuidesEndpoint     = BaseApiUrl + "/guide" + expand
	ReportsEndpoint    = BaseApiUrl + "/report"

	// to return everything from all applications within the subscription
	expand = "?expand=*"
)

type Doer interface {
	Do(req *http.Request) (*http.Response, error)
}

type HttpClient struct {
	doer   Doer
	logger *zerolog.Logger

	apiKey string

	pagesUrl      *url.URL
	featuresUrl   *url.URL
	trackTypesUrl *url.URL
	guidesUrl     *url.URL
	reportsUrl    *url.URL
}

var _ Client = &HttpClient{}

type HttpClientOption func(*HttpClient)

func WithHttpDoer(doer Doer) HttpClientOption {
	return func(c *HttpClient) {
		c.doer = doer
	}
}

func WithLogger(logger *zerolog.Logger) HttpClientOption {
	return func(c *HttpClient) {
		c.logger = logger
	}
}

func WithAPIKey(apiKey string) HttpClientOption {
	return func(c *HttpClient) {
		c.apiKey = apiKey
	}
}

func NewClient(opts ...HttpClientOption) (*HttpClient, error) {
	c := &HttpClient{}

	for _, opt := range opts {
		opt(c)
	}

	c.withDefaults()

	if c.prepareUrls() != nil {
		return nil, fmt.Errorf("failed to prepare urls: %w", c.prepareUrls())
	}

	if c.validate() != nil {
		return nil, fmt.Errorf("failed to validate client: %w", c.validate())
	}

	return c, nil
}

func (c *HttpClient) GetPages(ctx context.Context) ([]Page, error) {
	return getter[Page]{}.list(ctx, c.logger, c.apiKey, c.pagesUrl, c.doer)
}

func (c *HttpClient) GetFeatures(ctx context.Context) ([]Feature, error) {
	return getter[Feature]{}.list(ctx, c.logger, c.apiKey, c.featuresUrl, c.doer)
}

func (c *HttpClient) GetTrackTypes(ctx context.Context) ([]TrackType, error) {
	return getter[TrackType]{}.list(ctx, c.logger, c.apiKey, c.trackTypesUrl, c.doer)
}

func (c *HttpClient) GetGuides(ctx context.Context) ([]Guide, error) {
	return getter[Guide]{}.list(ctx, c.logger, c.apiKey, c.guidesUrl, c.doer)
}

func (c *HttpClient) GetReports(ctx context.Context) ([]Report, error) {
	return getter[Report]{}.list(ctx, c.logger, c.apiKey, c.reportsUrl, c.doer)
}

func (c *HttpClient) validate() error {
	if c.apiKey == "" {
		return fmt.Errorf("api key is required")
	}

	if c.doer == nil {
		return fmt.Errorf("doer client is required")
	}

	if c.logger == nil {
		return fmt.Errorf("logger is required")
	}

	return nil
}

func (c *HttpClient) withDefaults() {
	if c.doer == nil {
		c.doer = http.DefaultClient
	}

	if c.logger == nil {
		logger := zerolog.New(os.Stderr).With().Timestamp().Logger()
		c.logger = &logger
	}
}

func (c *HttpClient) prepareUrls() error {
	pagesUrl, err := url.Parse(PagesEndpoint)
	if err != nil {
		return fmt.Errorf("failed to parse pages url: %w", err)
	}
	c.pagesUrl = pagesUrl

	featuresUrl, err := url.Parse(FeaturesEndpoint)
	if err != nil {
		return fmt.Errorf("failed to parse features url: %w", err)
	}
	c.featuresUrl = featuresUrl

	trackTypesUrl, err := url.Parse(TrackTypesEndpoint)
	if err != nil {
		return fmt.Errorf("failed to parse track types url: %w", err)
	}
	c.trackTypesUrl = trackTypesUrl

	guidesUrl, err := url.Parse(GuidesEndpoint)
	if err != nil {
		return fmt.Errorf("failed to parse guides url: %w", err)
	}
	c.guidesUrl = guidesUrl

	reportsUrl, err := url.Parse(ReportsEndpoint)
	if err != nil {
		return fmt.Errorf("failed to parse reports url: %w", err)
	}
	c.reportsUrl = reportsUrl

	return nil
}

type getter[T any] struct{}

func (getter[T]) list(ctx context.Context, logger *zerolog.Logger, apiKey string, requestUrl *url.URL, doer Doer) ([]T, error) {
	req := &http.Request{
		Method: http.MethodGet,
		URL:    requestUrl,
		Header: http.Header{
			"X-Pendo-Integration-Key": []string{apiKey},
			"Content-Type":            []string{"application/json"},
		},
	}
	req = req.WithContext(ctx)

	response, err := doer.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to do doer request: %w", err)
	}
	defer func() {
		if err := response.Body.Close(); err != nil {
			logger.Err(err).Msg("failed to close response body")
		}
	}()

	if response.StatusCode != http.StatusOK {
		bytes, err := io.ReadAll(response.Body)
		if err != nil {
			return nil, fmt.Errorf("failed to read response body: %w", err)
		}
		logger.Error().
			Str("body", string(bytes)).
			Int("status_code", response.StatusCode).
			Msg("unexpected status code")
		return nil, fmt.Errorf("unexpected status code: %d", response.StatusCode)
	}

	var result []T
	if err := json.NewDecoder(response.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode pages: %w", err)
	}

	return result, nil
}
