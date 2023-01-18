package client

import "context"

type NopBackend struct{}

func (b *NopBackend) Set(ctx context.Context, table, clientID, value string) error {
	return nil
}

func (b *NopBackend) Get(ctx context.Context, table, clientID string) (string, error) {
	return "", nil
}

func (b *NopBackend) Close(ctx context.Context) error {
	return nil
}
