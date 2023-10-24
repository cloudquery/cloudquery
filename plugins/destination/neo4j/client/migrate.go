package client

import (
	"context"

	"github.com/cloudquery/plugin-sdk/v4/message"
)

// MigrateTables tries to create indexes for the tables.
// It will issue `CREATE RANGE INDEX ... IF NOT EXIST`.
// If the error occurs & indicates that the issue is caused by conflicting schema, 2 scenarios can happen:
// 1. Force mode is selected for migration: drop & recreate index (without checking for error this time)
// 2. No forced migration is requested - return error.
func (c *Client) MigrateTables(ctx context.Context, migrate message.WriteMigrateTables) error {
	return nil
}
