package client

import (
	"context"
	"time"

	"github.com/cloudquery/cq-provider-sdk/cqproto"
	"github.com/cloudquery/cq-provider-sdk/provider/schema/diag"
	"github.com/google/uuid"
	"github.com/huandu/go-sqlbuilder"
	"github.com/jackc/pgx/v4/pgxpool"
)

// FetchSummary includes a summarized report of fetch, such as fetch id, fetch start and finish,
// resources fetch results
type FetchSummary struct {
	СqId uuid.UUID `db:"id"`
	//  Unique Id of fetch session
	FetchId            uuid.UUID              `db:"fetch_id"`
	Start              time.Time              `db:"start"`
	Finish             time.Time              `db:"finish"`
	IsSuccess          bool                   `db:"is_success"`
	TotalResourceCount uint64                 `db:"total_resource_count"`
	TotalErrorsCount   uint64                 `db:"total_errors_count"`
	ProviderName       string                 `db:"provider_name"`
	ProviderVersion    string                 `db:"provider_version"`
	ProviderMeta       []byte                 `db:"provider_meta"` // reserved field to store provider's metadata such as
	FetchedResources   []ResourceFetchSummary `db:"results"`
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

// SaveFetchSummary saves fetch summary into fetches database
func SaveFetchSummary(ctx context.Context, pool *pgxpool.Pool, fs *FetchSummary) error {
	conn, err := pool.Acquire(ctx)
	if err != nil {
		return err
	}
	defer conn.Release()

	id, err := uuid.NewUUID()
	if err != nil {
		return err
	}
	fs.СqId = id
	dbStruct := sqlbuilder.NewStruct(new(FetchSummary))
	sql, args := dbStruct.InsertInto("cloudquery.fetches", fs).BuildWithFlavor(sqlbuilder.PostgreSQL)

	_, err = conn.Exec(ctx, sql, args...)
	return err
}
