package crm

import (
	"context"
	"fmt"
	"github.com/cloudquery/cloudquery/plugins/source/hubspot/client"
	"time"
)

func setLastModifiedDate(ctx context.Context, c *client.Client, tableName string, t time.Time) error {
	if err := c.Backend.SetKey(ctx, generateKey(c, tableName), t.Format(time.RFC3339Nano)); err != nil {
		return fmt.Errorf("failed to store cursor to backend: %w", err)
	}
	return nil
}

func getLastModifiedDate(ctx context.Context, c *client.Client, tableName string) (time.Time, error) {
	value, err := c.Backend.GetKey(ctx, generateKey(c, tableName))
	if err != nil {
		return time.Time{}, fmt.Errorf("failed to retrieve state from backend: %w", err)
	}
	if value == "" {
		return time.Time{}, nil
	}
	t, err := time.Parse(time.RFC3339Nano, value)
	if err != nil {
		return time.Time{}, fmt.Errorf("retrieved invalid state value: %q %w", value, err)
	}
	return t, nil
}

func generateKey(c *client.Client, tableName string) string {
	return fmt.Sprintf("%s-%s-last-modified-date", c.ID(), tableName)
}

func sortAscByField(fieldName string) []string {
	return []string{
		`{"propertyName":"` + fieldName + `","direction":"ASCENDING}"`,
	}
}
