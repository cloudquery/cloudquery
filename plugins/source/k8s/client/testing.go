package client

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/cloudquery/plugin-sdk/plugins/source"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/specs"
	"github.com/golang/mock/gomock"
	"github.com/rs/zerolog"
	"k8s.io/client-go/kubernetes"
)

func K8sMockTestHelper(t *testing.T, table *schema.Table, builder func(*testing.T, *gomock.Controller) kubernetes.Interface) {
	version := "vDev"

	t.Helper()

	table.IgnoreInTests = false

	mockController := gomock.NewController(t)
	l := zerolog.New(zerolog.NewTestWriter(t)).Output(
		zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.StampMicro},
	).Level(zerolog.DebugLevel).With().Timestamp().Logger()
	configureFunc := func(ctx context.Context, logger zerolog.Logger, s specs.Source, _ source.Options) (schema.ClientMeta, error) {
		var k8sSpec Spec
		if err := s.UnmarshalSpec(&k8sSpec); err != nil {
			return nil, fmt.Errorf("failed to unmarshal k8s spec: %w", err)
		}

		c := &Client{
			logger:   logger,
			Context:  "testContext",
			spec:     &k8sSpec,
			contexts: []string{"testContext"},
		}
		c.clients = map[string]kubernetes.Interface{"testContext": builder(t, mockController)}
		return c, nil
	}

	plugin := source.NewPlugin(
		table.Name,
		version,
		[]*schema.Table{
			table,
		},
		configureFunc,
	)
	plugin.SetLogger(l)
	source.TestPluginSync(t, plugin, specs.Source{
		Name:         "dev",
		Path:         "cloudquery/dev",
		Version:      version,
		Tables:       []string{table.Name},
		Destinations: []string{"mock-destination"},
	})
}
