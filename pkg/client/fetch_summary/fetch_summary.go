package fetch_summary

import (
	"context"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"

	"github.com/jackc/pgx/v4"

	"github.com/cloudquery/cq-provider-sdk/cqproto"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/cloudquery/cq-provider-sdk/provider/schema/diag"
	"github.com/doug-martin/goqu/v9"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/google/uuid"
)

type FetchSummaryClient struct {
	db schema.QueryExecer
}

func NewFetchSummaryClient(db schema.QueryExecer) *FetchSummaryClient {
	return &FetchSummaryClient{
		db: db,
	}
}

// FetchSummary includes a summarized report of fetch, such as fetch id, fetch start and finish,
// resources fetch results
type FetchSummary struct {
	CqId uuid.UUID `db:"id"`
	//  Unique Id of fetch session
	FetchId            uuid.UUID              `db:"fetch_id"`
	CreatedAt          *time.Time             `db:"created_at"`
	Start              *time.Time             `db:"start"`
	Finish             *time.Time             `db:"finish"`
	IsSuccess          bool                   `db:"is_success"`
	TotalResourceCount uint64                 `db:"total_resource_count"`
	TotalErrorsCount   uint64                 `db:"total_errors_count"`
	ProviderName       string                 `db:"provider_name"`
	ProviderAlias      string                 `db:"provider_alias"`
	ProviderVersion    string                 `db:"provider_version"`
	CoreVersion        string                 `db:"core_version"`
	Resources          ResourceFetchSummaries `db:"results"`
}

type ResourceFetchSummaries []ResourceFetchSummary

// Value implements Valuer interface required by goqu
func (r ResourceFetchSummaries) Value() (driver.Value, error) {
	if len(r) == 0 {
		return nil, nil
	}
	return json.Marshal(r)
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
func (c *FetchSummaryClient) SaveFetchSummary(ctx context.Context, fs *FetchSummary) error {
	id, err := uuid.NewUUID()
	if err != nil {
		return err
	}
	fs.CqId = id
	q := goqu.Dialect("postgres").Insert("cloudquery.fetches").Rows(fs)
	sql, args, err := q.ToSQL()
	if err != nil {
		return err
	}
	return c.db.Exec(ctx, sql, args...)
}

// GetLatestFetchSummaryForProvider gets latest fetch summary for specific provider
func (c *FetchSummaryClient) GetLatestFetchSummaryForProvider(ctx context.Context, provider string) (*FetchSummary, error) {
	q := goqu.Dialect("postgres").
		Select("provider_version", "is_success").
		From("cloudquery.fetches").
		Where(goqu.Ex{"provider_name": provider, "finish": goqu.Op{"isNot": nil}}).
		Limit(1).
		Order(goqu.I("finish").Desc())
	sql, _, err := q.ToSQL()
	if err != nil {
		return nil, err
	}
	var data FetchSummary
	err = pgxscan.Get(ctx, c.db, &data, sql)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, errors.New("there is no successful fetch for requested provider")
		}
		return nil, err
	}
	return &data, nil
}
