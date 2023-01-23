package client

import (
	"context"
	"strconv"

	"github.com/cloudquery/plugin-sdk/backend"
)

type BackendWrapper struct {
	lastValuePerID map[string]int64

	backend.Backend
}

func NewBackendWrapper(b backend.Backend) *BackendWrapper {
	return &BackendWrapper{
		lastValuePerID: map[string]int64{},
		Backend:        b,
	}
}

func (b *BackendWrapper) SetIfMaximum(ctx context.Context, table, clientID string, value int64) error {
	key := table + "|" + clientID

	current, ok := b.lastValuePerID[key]
	if !ok {
		b.lastValuePerID[key] = value
		return b.Backend.Set(ctx, table, clientID, strconv.FormatInt(value, 10))
	}

	if current >= value {
		return nil
	}

	b.lastValuePerID[key] = value
	return b.Backend.Set(ctx, table, clientID, strconv.FormatInt(value, 10))
}
