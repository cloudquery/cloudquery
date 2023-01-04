package services

import "context"

//go:generate mockgen -package=mocks -destination=../mocks/backend.go -source=backend.go BackendClient
type Backend interface {
	Set(ctx context.Context, table, key, value string) error
	Get(ctx context.Context, table, key string) (string, error)
	Close(ctx context.Context) error
}
