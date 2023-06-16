package client

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/cloudquery/plugin-pb-go/specs"
	"github.com/cloudquery/plugin-sdk/v3/plugins/source"
	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/golang/mock/gomock"
	"github.com/rs/zerolog"
	v1 "k8s.io/api/core/v1"
	apiextensionsclientset "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	"k8s.io/client-go/kubernetes"
)

type TestOption func(*Client)

func WithTestNamespaces(namespaces ...v1.Namespace) TestOption {
	return func(c *Client) {
		c.namespaces[c.Context] = namespaces
	}
}

func K8sMockTestHelper(t *testing.T, table *schema.Table, builder func(*testing.T, *gomock.Controller) kubernetes.Interface, opts ...TestOption) {
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
			logger:     logger,
			Context:    "testContext",
			spec:       &k8sSpec,
			contexts:   []string{"testContext"},
			namespaces: map[string][]v1.Namespace{},
		}
		c.clients = map[string]kubernetes.Interface{"testContext": builder(t, mockController)}
		for _, opt := range opts {
			opt(c)
		}
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

func APIExtensionsMockTestHelper(t *testing.T, table *schema.Table, builder func(*testing.T, *gomock.Controller) apiextensionsclientset.Interface, opts ...TestOption) {
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
			logger:     logger,
			Context:    "testContext",
			spec:       &k8sSpec,
			contexts:   []string{"testContext"},
			namespaces: map[string][]v1.Namespace{},
		}
		c.apiExtensions = map[string]apiextensionsclientset.Interface{"testContext": builder(t, mockController)}
		for _, opt := range opts {
			opt(c)
		}
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
