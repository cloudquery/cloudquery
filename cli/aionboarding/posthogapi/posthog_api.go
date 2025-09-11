package posthogapi

import (
	"context"
	"fmt"
)

type PosthogAPI interface {
	Notify(ctx context.Context, event string, properties map[string]string) error
}

type posthogAPI struct {
}

func NewPosthogAPI(apiKey string) PosthogAPI {
	return &posthogAPI{}
}

func (a *posthogAPI) Notify(ctx context.Context, event string, properties map[string]string) error {
	fmt.Println(event, properties)
	return nil
}
