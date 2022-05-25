package state

import (
	"context"
	"time"

	"github.com/doug-martin/goqu/v9"
	"github.com/doug-martin/goqu/v9/exp"
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
	return c.exec(ctx, goqu.Dialect("postgres").Insert("cloudquery.check_results").Rows(cr))
}

// CreatePolicyExecution inserts a policy execution in the database
func (c *Client) CreatePolicyExecution(ctx context.Context, pe *PolicyExecution) (_ *PolicyExecution, err error) {
	if pe.Id, err = uuid.NewUUID(); err != nil {
		return nil, err
	}
	pe.Timestamp = time.Now()
	return pe, c.exec(ctx, goqu.Dialect("postgres").Insert("cloudquery.policy_executions").Rows(pe))
}

// PrunePolicyExecutions deletes old policy executions in the database
func (c *Client) PrunePolicyExecutions(ctx context.Context, policyName string, pruneBefore time.Time) error {
	var expression exp.Expression = goqu.C("timestamp").Lt(pruneBefore)
	if policyName != "*" {
		expression = goqu.And(expression, goqu.C("policy_name").Eq(policyName))
	}
	return c.exec(ctx, goqu.Dialect("postgres").Delete("cloudquery.policy_executions").Where(expression))
}

// exec runs a goqu sql expression
func (c *Client) exec(ctx context.Context, sqlExpression exp.SQLExpression) error {
	sql, args, err := sqlExpression.ToSQL()
	if err != nil {
		return err
	}
	return c.db.Exec(ctx, sql, args...)
}
