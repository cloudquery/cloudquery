package client

import "context"

type NopBackend struct{}

func (*NopBackend) Set(ctx context.Context, table, clientID, value string) error {
	return nil
}

func (*NopBackend) Get(ctx context.Context, table, clientID string) (string, error) {
	return "", nil
}

func (*NopBackend) Close(ctx context.Context) error {
	return nil
}
