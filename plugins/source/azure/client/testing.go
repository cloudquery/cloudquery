package client

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/cloudquery/plugin-sdk/plugins"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/specs"
	"github.com/golang/mock/gomock"
	"github.com/rs/zerolog"
)

func MockTestHelper(t *testing.T, table *schema.Table, createServices func(t *testing.T, ctrl *gomock.Controller) *Services) {
	version := "vDev"

	t.Helper()

	table.IgnoreInTests = false
	l := zerolog.New(zerolog.NewTestWriter(t)).Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.StampMicro}).Level(zerolog.DebugLevel).With().Timestamp().Logger()
	newTestExecutionClient := func(ctx context.Context, logger zerolog.Logger, spec specs.Source) (schema.ClientMeta, error) {
		svc := createServices(t, gomock.NewController(t))
		servicesMap := make(map[string]*Services)
		servicesMap["testSubscription"] = svc
		return &Client{logger: l, services: servicesMap}, nil
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

func CreatePager[RESPONSE any](response RESPONSE) *runtime.Pager[RESPONSE] {
	return runtime.NewPager(runtime.PagingHandler[RESPONSE]{
		More: func(RESPONSE) bool {
			// it'll be called ONLY after the 1st page was processed
			return false
		},
		Fetcher: func(context.Context, *RESPONSE) (RESPONSE, error) {
			return response, nil
		},
	})
}
