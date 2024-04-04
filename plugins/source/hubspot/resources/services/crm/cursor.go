package crm

import (
	"context"
	"fmt"
	"github.com/cloudquery/cloudquery/plugins/source/hubspot/client"
)

func setCursor(ctx context.Context, cqClient *client.Client, tableName string, cursor string) error {
	if len(cursor) > 0 {
		if err := cqClient.Backend.SetKey(ctx, generateKey(cqClient, tableName), cursor); err != nil {
			return fmt.Errorf("failed to store cursor to backend: %w", err)
		}
	}
	return nil
}

func getCursor(ctx context.Context, cqClient *client.Client, key string) (string, error) {
	cursor, err := cqClient.Backend.GetKey(ctx, generateKey(cqClient, key))
	if err != nil {
		return "", fmt.Errorf("failed to retrieve cursor from backend: %w", err)
	}
	return cursor, nil
}

func generateKey(cqClient *client.Client, tableName string) string {
	return fmt.Sprintf("%s-%s", cqClient.ID(), tableName)
}
