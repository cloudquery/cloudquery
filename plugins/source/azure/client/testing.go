package client

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/cloudquery/plugin-sdk/plugins"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/specs"
	"github.com/rs/zerolog"
)

const TestSubscription = "12345678-1234-1234-1234-123456789000"

type MockCreds struct {

}

func (*MockCreds) GetToken(ctx context.Context, options policy.TokenRequestOptions) (azcore.AccessToken, error) {
	return azcore.AccessToken{
		Token: "SomeToken",
		ExpiresOn: time.Now().Add(time.Hour*24),
	}, nil
}

func MockTestHelper(t *testing.T, table *schema.Table, createServices func() (*arm.ClientOptions, error)) {
	version := "vDev"
	t.Helper()
	table.IgnoreInTests = false
	creds := &MockCreds{}
	l := zerolog.New(zerolog.NewTestWriter(t)).Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.StampMicro}).Level(zerolog.DebugLevel).With().Timestamp().Logger()
	newTestExecutionClient := func(ctx context.Context, logger zerolog.Logger, spec specs.Source) (schema.ClientMeta, error) {
		options, err := createServices()
		if err != nil {
			return nil, err
		}
		c := &Client{
			logger:        l,
			Options: options,
			Creds: creds,
			subscriptions: []string{TestSubscription},
		}

		return c, nil
	}

	p := plugins.NewSourcePlugin(table.Name, version, []*schema.Table{table}, newTestExecutionClient)
	p.SetLogger(l)
	plugins.TestSourcePluginSync(t, p, specs.Source{
		Name:         "dev",
		Path:         "cloudquery/dev",
		Version:      version,
		Tables:       []string{table.Name},
		Destinations: []string{"mock-destination"},
	})
}
