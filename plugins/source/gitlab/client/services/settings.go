package services

import "github.com/xanzy/go-gitlab"

type SettingsClient interface {
	GetSettings(options ...gitlab.RequestOptionFunc) (*gitlab.Settings, *gitlab.Response, error)
}
