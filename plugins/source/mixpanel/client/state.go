package client

import (
	"context"
	"strconv"
	"sync"

	"github.com/cloudquery/plugin-sdk/backend"
)

type BackendWrapper struct {
	backend.Backend

	lastValuePerID map[string]int64
	mu             *sync.RWMutex
}

func NewBackendWrapper(b backend.Backend) *BackendWrapper {
	return &BackendWrapper{
		Backend: b,

		lastValuePerID: map[string]int64{},
		mu:             &sync.RWMutex{},
	}
}

// SetHWM sets the high watermark for the table and clientID pair.
func (b *BackendWrapper) SetHWM(ctx context.Context, table, clientID string, value int64) error {
	key := table + "|" + clientID

	b.mu.RLock()
	current, ok := b.lastValuePerID[key]
	b.mu.RUnlock()
	if !ok {
		b.mu.Lock()
		b.lastValuePerID[key] = value
		b.mu.Unlock()
		return b.Backend.Set(ctx, table, clientID, strconv.FormatInt(value, 10))
	}

	if current >= value {
		return nil
	}

	b.mu.Lock()
	b.lastValuePerID[key] = value
	b.mu.Unlock()
	return b.Backend.Set(ctx, table, clientID, strconv.FormatInt(value, 10))
}
