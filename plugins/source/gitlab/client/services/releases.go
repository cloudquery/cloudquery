package services

import "github.com/xanzy/go-gitlab"

//go:generate mockgen -package=mocks -destination=../mocks/releases.go -source=releases.go ReleasesClient
type ReleasesClient interface {
	ListReleases(pid interface{}, opt *gitlab.ListReleasesOptions, options ...gitlab.RequestOptionFunc) ([]*gitlab.Release, *gitlab.Response, error)
}
