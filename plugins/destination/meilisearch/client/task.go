package client

import (
	"context"
	"fmt"
	"time"

	"github.com/meilisearch/meilisearch-go"
)

func (c *Client) waitTask(ctx context.Context, info *meilisearch.TaskInfo) error {
	// Use a default interval of 50ms for polling
	task, err := c.Meilisearch.TaskManager().WaitForTaskWithContext(ctx, info.TaskUID, 50*time.Millisecond)
	if err != nil {
		return err
	}
	if task.Status == meilisearch.TaskStatusSucceeded {
		return nil
	}
	return fmt.Errorf("wait for task %d finished with status %q", info.TaskUID, task.Status)
}
