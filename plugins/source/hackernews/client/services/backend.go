package services

import "github.com/cloudquery/plugin-sdk/v3/backend"

//go:generate mockgen -package=mocks -destination=../mocks/backend.go -source=backend.go BackendClient
type Backend interface {
	backend.Backend
}
