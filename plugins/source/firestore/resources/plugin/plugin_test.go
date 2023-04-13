package plugin_test

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/cloudquery/cloudquery/plugins/source/firestore/client"
	"github.com/cloudquery/cloudquery/plugins/source/firestore/resources/plugin"
	"github.com/cloudquery/plugin-sdk/specs"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/require"
)

func TestPlugin(t *testing.T) {
	p := plugin.Plugin()
	ctx := context.Background()
	l := zerolog.New(zerolog.NewTestWriter(t)).Output(
		zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.StampMicro},
	).Level(zerolog.DebugLevel).With().Timestamp().Logger()
	p.SetLogger(l)

	tableSetup(ctx, t, "test_firestore_source")

	spec := specs.Source{
		Name:         "test_firestore_source",
		Path:         "cloudquery/firestore",
		Version:      "vDevelopment",
		Destinations: []string{"test"},
		Tables:       []string{"test_firestore_source"},
		Spec: client.Spec{
			ProjectID: "cqtest-project",
		},
	}
	require.NoError(t, p.Init(ctx, spec))
}

func TestPlugin_OrderBy(t *testing.T) {
	p := plugin.Plugin()
	ctx := context.Background()
	l := zerolog.New(zerolog.NewTestWriter(t)).Output(
		zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.StampMicro},
	).Level(zerolog.DebugLevel).With().Timestamp().Logger()
	p.SetLogger(l)

	tableSetup(ctx, t, "test_firestore_source")

	spec := specs.Source{
		Name:         "test_firestore_source",
		Path:         "cloudquery/firestore",
		Version:      "vDevelopment",
		Destinations: []string{"test"},
		Tables:       []string{"test_firestore_source"},
		Spec: client.Spec{
			ProjectID:        "cqtest-project",
			OrderByField:     "test",
			OrderByDirection: "DESC",
		},
	}
	require.NoError(t, p.Init(ctx, spec))
}
