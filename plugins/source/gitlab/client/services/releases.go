package services

import "github.com/xanzy/go-gitlab"

type ReleasesClient interface {
	ListReleases(pid any, opt *gitlab.ListReleasesOptions, options ...gitlab.RequestOptionFunc) ([]*gitlab.Release, *gitlab.Response, error)
}
