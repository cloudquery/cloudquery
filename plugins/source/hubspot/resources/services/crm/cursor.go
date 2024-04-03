package crm

import (
	"context"
	"fmt"
	"github.com/cloudquery/cloudquery/plugins/source/hubspot/client"
)

func setCursor(ctx context.Context, cqClient *client.Client, key string, cursor string) error {
	if cqClient.Backend != nil && len(cursor) > 0 {
		if err := cqClient.Backend.SetKey(ctx, key, cursor); err != nil {
			return fmt.Errorf("failed to store cursor to backend: %w", err)
		}
	}
	return nil
}

func getCursor(ctx context.Context, cqClient *client.Client, key string) (string, error) {
	if cqClient.Backend == nil {
		return "", nil
	}
	cursor, err := cqClient.Backend.GetKey(ctx, key)
	if err != nil {
		return "", fmt.Errorf("failed to retrieve cursor from backend: %w", err)
	}
	return cursor, nil
}
