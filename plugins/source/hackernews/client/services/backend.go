package services

import "github.com/cloudquery/plugin-sdk/v4/state"

//go:generate mockgen -package=mocks -destination=../mocks/backend.go -source=backend.go BackendClient
type BackendClient interface {
	state.Client
}
