package fetch_summary

import (
	"context"
	"time"

	"github.com/huandu/go-sqlbuilder"
	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/cloudquery/cq-provider-sdk/cqproto"
	"github.com/cloudquery/cq-provider-sdk/provider/schema/diag"
	"github.com/google/uuid"
)

const createCqFetchesTable = `
CREATE TABLE IF NOT EXISTS cq_fetches
(
    cq_id                UUID NOT NULL,
    fetch_id             UUID NOT NULL,
    START                TIMESTAMP,
    finish               TIMESTAMP,
    total_resource_count BIGINT,
    provider_name        TEXT,
    provider_version     TEXT,
    provider_meta        jsonb,
    fetch_results        jsonb,
    CONSTRAINT "cq_fetches_id" PRIMARY KEY (cq_id),
    CONSTRAINT "cq_fetches_pk" UNIQUE (fetch_id, provider_name),
	CONSTRAINT "non_nil_fetch_id" CHECK (fetch_id != '00000000-0000-0000-0000-000000000000')
);
`

type FetchSummarizer struct {
	pool     *pgxpool.Pool
	dbStruct *sqlbuilder.Struct
}

// NewFetchSummarizer creates cq_fetches table and returns a summarizer that saves fetch summary to cq_fetches table
func NewFetchSummarizer(ctx context.Context, pool *pgxpool.Pool) (*FetchSummarizer, error) {
	conn, err := pool.Acquire(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Release()

	// create table if not exists
	if _, err := conn.Exec(ctx, createCqFetchesTable); err != nil {
		return nil, err
	}

	return &FetchSummarizer{
		pool:     pool,
		dbStruct: sqlbuilder.NewStruct(new(FetchSummary)),
	}, nil
}

// FetchSummary includes a summarized report of fetch, such as fetch id, fetch start and finish,
// resources fetch results
type FetchSummary struct {
	СqId uuid.UUID `db:"cq_id"`
	//  Unique Id of fetch session
	FetchId            uuid.UUID              `db:"fetch_id"`
	Start              time.Time              `db:"start"`
	Finish             time.Time              `db:"finish"`
	TotalResourceCount uint64                 `db:"total_resource_count"`
	ProviderName       string                 `db:"provider_name"`
	ProviderVersion    string                 `db:"provider_version"`
	ProviderMeta       []byte                 `db:"provider_meta"` // reserved field to store providers metadata such as
	FetchedResources   []ResourceFetchSummary `db:"fetch_results"`
}

// ResourceFetchSummary includes a data about fetching specific resource
type ResourceFetchSummary struct {
	ResourceName string `json:"resource_name"`
	// map of resources that have finished fetching
	FinishedResources map[string]bool `json:"finished_resources"`
	// Amount of resources collected so far
	// Error value if any, if returned the stream will be canceled
	Error string `json:"error"`
	// list of resources where the fetching failed
	PartialFetchFailedResources []*cqproto.FailedResourceFetch `json:"partial_fetch_failed_resources"`
	// Execution status of resource
	Status string `json:"status"`
	// Total Amount of resources collected by this resource
	ResourceCount uint64 `json:"resource_count"`
	// Diagnostics of failed resource fetch, the diagnostic provides insights such as severity, summary and
	// details on how to solve this issue
	Diagnostics diag.Diagnostics `json:"diagnostics"`
}

// SaveFetchSummary saves fetch summary into cq_fetches database
func (c *FetchSummarizer) SaveFetchSummary(ctx context.Context, fs FetchSummary) error {
	conn, err := c.pool.Acquire(ctx)
	if err != nil {
		return err
	}
	defer conn.Release()

	id, err := uuid.NewUUID()
	if err != nil {
		return err
	}
	fs.СqId = id
	sql, args := c.dbStruct.InsertInto("cq_fetches", fs).BuildWithFlavor(sqlbuilder.PostgreSQL)

	_, err = conn.Exec(ctx, sql, args...)
	return err
}
