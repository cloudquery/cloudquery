package meta_storage

import (
	"context"
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/cloudquery/cloudquery/pkg/client/database"
	"github.com/cloudquery/cloudquery/pkg/client/history"
	sdkdb "github.com/cloudquery/cq-provider-sdk/database"
	"github.com/google/uuid"
	"github.com/hashicorp/go-hclog"
	"github.com/stretchr/testify/assert"
)

const testDBConnection = "postgres://postgres:pass@localhost:5432/postgres?sslmode=disable"

type fetchSummaryTest struct {
	summary     FetchSummary
	err         error
	skipFetchId bool
}

var fetchSummaryTests = []fetchSummaryTest{
	{
		summary: FetchSummary{
			ProviderName:    "test",
			ProviderVersion: "v0.0.0",
		},
	},
	{
		summary: FetchSummary{
			ProviderName: "test1",
			Resources: []ResourceFetchSummary{
				{
					ResourceName:  "test",
					ResourceCount: 99,
				},
			},
		},
	},
	{
		summary: FetchSummary{
			ProviderName:    "test2",
			ProviderVersion: "v0.0.1",
		},
	},
	{
		summary: FetchSummary{
			ProviderName: "test4",
			Resources: []ResourceFetchSummary{
				{
					ResourceName:  "test",
					ResourceCount: 99,
				},
				{
					ResourceName:  "test2",
					ResourceCount: 99,
				},
			},
		},
	},
	{
		summary: FetchSummary{
			ProviderName:    "test2",
			ProviderVersion: "v0.0.1",
		},
		err: errors.New("ERROR: duplicate key value violates unique constraint \"fetches_pk\" (SQLSTATE 23505)"),
	},
	{
		summary: FetchSummary{
			ProviderName:    "test3",
			ProviderVersion: "v0.0.1",
		},
		skipFetchId: true,
		err:         errors.New("ERROR: new row for relation \"fetches\" violates check constraint \"non_nil_fetch_id\" (SQLSTATE 23514)"),
	},
}

func TestFetchSaveSummary(t *testing.T) {
	// create database connection
	db, err := sdkdb.New(context.Background(), hclog.NewNullLogger(), testDBConnection)
	assert.NoError(t, err)

	fetchSummaryClient := NewClient(db, hclog.NewNullLogger())

	_, de, err := database.GetExecutor(hclog.NewNullLogger(), testDBConnection, &history.Config{})
	if err != nil {
		t.Fatal(fmt.Errorf("getExecutor: %w", err))
	}

	err = fetchSummaryClient.MigrateCore(context.Background(), de)
	assert.NoError(t, err)

	fetchId := uuid.New()
	for _, f := range fetchSummaryTests {
		if !f.skipFetchId {
			f.summary.FetchId = fetchId
		}
		start := time.Now()
		f.summary.Start = &start
		err := fetchSummaryClient.SaveFetchSummary(context.Background(), &f.summary)
		if f.err != nil {
			assert.EqualError(t, err, f.err.Error())
		} else {
			assert.NoError(t, err)
		}
	}
}
