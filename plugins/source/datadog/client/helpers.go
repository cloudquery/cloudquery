package client

import (
	"context"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

func ConsumePaginatedResponse[T any](response <-chan datadog.PaginationResult[T], cancel context.CancelFunc, res chan<- any) error {
	defer cancel()
	for item := range response {
		if item.Error != nil {
			return item.Error
		}
		res <- item.Item
	}
	return nil
}
