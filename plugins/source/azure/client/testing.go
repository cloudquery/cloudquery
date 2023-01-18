package client

import (
	"context"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"testing"
	"time"

	"github.com/gorilla/mux"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armresources"
	"github.com/cloudquery/plugin-sdk/plugins/source"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/specs"
	"github.com/rs/zerolog"
)

const TestSubscription = "12345678-1234-1234-1234-123456789000"

var testResourceGroup = "test-resource-group"

type MockCreds struct {
}

func (*MockCreds) GetToken(ctx context.Context, options policy.TokenRequestOptions) (azcore.AccessToken, error) {
	return azcore.AccessToken{
		Token:     "SomeToken",
		ExpiresOn: time.Now().Add(time.Hour * 24),
	}, nil
}

type MockHttpClient struct {
	rootURL string
	scheme  string
	host    string
	client  *http.Client
}

func NewMockHttpClient(cl *http.Client, rootURL string) *MockHttpClient {
	u, err := url.Parse(rootURL)
	if err != nil {
		panic(err)
	}
	return &MockHttpClient{
		client:  cl,
		rootURL: rootURL,
		scheme:  u.Scheme,
		host:    u.Host,
	}
}

func (c *MockHttpClient) Do(req *http.Request) (*http.Response, error) {
	req.URL.Host = c.host
	req.URL.Scheme = c.scheme
	return c.client.Do(req)
}

func MockTestHelper(t *testing.T, table *schema.Table, createServices func(*mux.Router) error) {
	version := "vDev"
	t.Helper()
	debug = true
	table.IgnoreInTests = false
	creds := &MockCreds{}
	router := mux.NewRouter()
	h := httptest.NewServer(router)
	defer h.Close()
	mockClient := NewMockHttpClient(h.Client(), h.URL)

	l := zerolog.New(zerolog.NewTestWriter(t)).Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.StampMicro}).Level(zerolog.DebugLevel).With().Timestamp().Logger()
	newTestExecutionClient := func(ctx context.Context, logger zerolog.Logger, spec specs.Source, _ source.Options) (schema.ClientMeta, error) {
		err := createServices(router)
		if err != nil {
			return nil, err
		}
		registeredNamespaces := make(map[string]map[string]bool)
		registeredNamespaces[TestSubscription] = make(map[string]bool)
		for _, namespace := range namespaces {
			registeredNamespaces[TestSubscription][namespace] = true
		}
		c := &Client{
			logger: l,
			Options: &arm.ClientOptions{
				ClientOptions: policy.ClientOptions{
					Transport: mockClient,
				},
			},
			registeredNamespaces: registeredNamespaces,
			Creds:                creds,
			subscriptions:        []string{TestSubscription},
			resourceGroups: map[string][]*armresources.GenericResourceExpanded{
				TestSubscription: {
					{
						Name: &testResourceGroup,
					},
				},
			},
		}

		return c, nil
	}

	p := source.NewPlugin(table.Name, version, []*schema.Table{table}, newTestExecutionClient)
	p.SetLogger(l)
	source.TestPluginSync(t, p, specs.Source{
		Name:         "dev",
		Path:         "cloudquery/dev",
		Version:      version,
		Tables:       []string{table.Name},
		Destinations: []string{"mock-destination"},
	})
}
