package client

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/cloudquery/plugin-sdk/v4/scheduler"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/golang/mock/gomock"
	"github.com/google/go-github/v49/github"
	"github.com/rs/zerolog"
)

type TestOptions struct{}

func GithubMockTestHelper(t *testing.T, table *schema.Table, builder func(*testing.T, *gomock.Controller) GithubServices, _ TestOptions) {
	table.IgnoreInTests = false
	t.Helper()
	ctrl := gomock.NewController(t)
	l := zerolog.New(zerolog.NewTestWriter(t)).Output(
		zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.StampMicro},
	).Level(zerolog.DebugLevel).With().Timestamp().Logger()
	sched := scheduler.NewScheduler(scheduler.WithLogger(l))

	var cs github.Repository
	if err := faker.FakeObject(&cs); err != nil {
		t.Fatal(err)
	}
	someId := int64(5555555)
	cs.Parent = &github.Repository{ID: &someId}
	cs.TemplateRepository = &github.Repository{ID: &someId}
	cs.Source = &github.Repository{ID: &someId}

	b := builder(t, ctrl)

	c := &Client{
		logger: l,
		Github: b,
		orgServices: map[string]GithubServices{
			"":        b,
			"testorg": b,
		},
		orgs:            []string{"testorg"},
		orgRepositories: map[string][]*github.Repository{"testorg": {&cs}},
	}

	messages, err := sched.SyncAll(context.Background(), c, schema.Tables{table})
	if err != nil {
		t.Fatalf("failed to sync: %v", err)
	}
	records := messages.GetInserts().GetRecordsForTable(table)
	emptyColumns := schema.FindEmptyColumns(table, records)
	if len(emptyColumns) > 0 {
		t.Fatalf("empty columns: %v", emptyColumns)
	}
}
