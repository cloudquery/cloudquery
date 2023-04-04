package plugin

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/cloudquery/cloudquery/plugins/source/firestore/client"
	"github.com/cloudquery/plugin-sdk/specs"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/require"
)

func getTestConnectionString() string {
	testConn := os.Getenv("CQ_SOURCE_FIRESTORE_TEST_CONNECTION_STRING")
	if testConn == "" {
		return `{}`
	}
	return testConn
}

func TestPlugin(t *testing.T) {
	p := Plugin()
	ctx := context.Background()
	l := zerolog.New(zerolog.NewTestWriter(t)).Output(
		zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.StampMicro},
	).Level(zerolog.DebugLevel).With().Timestamp().Logger()
	p.SetLogger(l)
	spec := specs.Source{
		Name:         "test_firestore_source",
		Path:         "cloudquery/firestore",
		Version:      "vDevelopment",
		Destinations: []string{"test"},
		Tables:       []string{"test_firestore_source"},
		Spec: client.Spec{
			ProjectID: "test-project",
		},
	}
	err := p.Init(ctx, spec)
	require.NoError(t, err)
}
