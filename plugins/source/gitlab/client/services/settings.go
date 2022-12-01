package services

import "github.com/xanzy/go-gitlab"

//go:generate mockgen -package=mocks -destination=../mocks/settings.go -source=settings.go SettingsClient
type SettingsClient interface {
	GetSettings(options ...gitlab.RequestOptionFunc) (*gitlab.Settings, *gitlab.Response, error)
}
