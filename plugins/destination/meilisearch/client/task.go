package client

import (
	"context"
	"fmt"

	"github.com/meilisearch/meilisearch-go"
)

func (c *Client) waitTask(ctx context.Context, info *meilisearch.TaskInfo) error {
	task, err := c.Meilisearch.WaitForTask(info.TaskUID, meilisearch.WaitParams{Context: ctx})
	if err != nil {
		return err
	}
	if task.Status == meilisearch.TaskStatusSucceeded {
		return nil
	}
	return fmt.Errorf("wait for task %q finished with status %q", info.TaskUID, task.Status)
}
