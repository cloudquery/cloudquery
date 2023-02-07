package client

import (
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/cloudquery/plugin-sdk/plugins/source"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/specs"
	"github.com/julienschmidt/httprouter"
	"github.com/rs/zerolog"
)

func MockTestHelper(t *testing.T, table *schema.Table, createService func(*httprouter.Router) error) {
	version := "vDev"
	t.Helper()

	table.IgnoreInTests = false
	mux := httprouter.New()
	ts := httptest.NewUnstartedServer(mux)
	defer ts.Close()
	if err := createService(mux); err != nil {
		t.Fatalf("failed to createService: %v", err)
	}
	ts.Start()

	l := zerolog.New(zerolog.NewTestWriter(t)).Output(
		zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.StampMicro},
	).Level(zerolog.DebugLevel).With().Timestamp().Logger()

	p := source.NewPlugin(
		table.Name,
		version,
		[]*schema.Table{
			table,
		},
		Configure)
	p.SetLogger(l)

	source.TestPluginSync(t, p, specs.Source{
		Name:         "dev",
		Path:         "cloudquery/dev",
		Version:      version,
		Tables:       []string{table.Name},
		Spec:         &Spec{APIKey: "test", Tailnet: "test", EndpointURL: ts.URL},
		Destinations: []string{"mock-destination"},
	})
}
