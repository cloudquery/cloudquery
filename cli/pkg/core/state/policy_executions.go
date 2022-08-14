package state

import (
	"context"
	"time"

	"github.com/doug-martin/goqu/v9"
	"github.com/google/uuid"
)

// CheckResult represents the output of a single check run
type CheckResult struct {
	ExecutionId        uuid.UUID `db:"execution_id"`
	ExecutionTimestamp time.Time `db:"execution_timestamp"`
	Name               string    `db:"name"`
	Selector           string    `db:"selector"`
	Description        string    `db:"description"`
	Status             string    `db:"status"`
	RawResults         string    `db:"raw_results"`
	Error              string    `db:"error"`
}

// PolicyExecution represents a single policy execution and their summary
type PolicyExecution struct {
	Id           uuid.UUID `db:"id"`
	Timestamp    time.Time `db:"timestamp"`
	Scheme       string    `db:"scheme"`
	Location     string    `db:"location"`
	PolicyName   string    `db:"policy_name"`
	Selector     string    `db:"selector"`
	Sha256Hash   string    `db:"sha256_hash"`
	Version      string    `db:"version"`
	ChecksTotal  int       `db:"checks_total"`
	ChecksFailed int       `db:"checks_failed"`
	ChecksPassed int       `db:"checks_passed"`
}

// CreateCheckResult inserts a check result in the database
func (c *Client) CreateCheckResult(ctx context.Context, cr *CheckResult) error {
	q := goqu.Dialect("postgres").Insert("cloudquery.check_results").Rows(cr)
	sql, args, err := q.ToSQL()
	if err != nil {
		return err
	}
	return c.db.Exec(ctx, sql, args...)
}

// CreatePolicyExecution inserts a policy execution in the database
func (c *Client) CreatePolicyExecution(ctx context.Context, pe *PolicyExecution) (*PolicyExecution, error) {
	var err error
	if pe.Id, err = uuid.NewUUID(); err != nil {
		return nil, err
	}
	pe.Timestamp = time.Now()
	q := goqu.Dialect("postgres").Insert("cloudquery.policy_executions").Rows(pe)
	sql, args, err := q.ToSQL()
	if err != nil {
		return nil, err
	}
	return pe, c.db.Exec(ctx, sql, args...)
}

// PrunePolicyExecutions deletes old policy executions in the database
func (c *Client) PrunePolicyExecutions(ctx context.Context, pruneBefore time.Time) error {
	q := goqu.Dialect("postgres").Delete("cloudquery.policy_executions").Where(goqu.Ex{
		"timestamp": goqu.Op{"lt": pruneBefore},
	})
	sql, args, err := q.ToSQL()
	if err != nil {
		return err
	}
	return c.db.Exec(ctx, sql, args...)
}
