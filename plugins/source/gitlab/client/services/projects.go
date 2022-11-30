package services

import "github.com/xanzy/go-gitlab"

//go:generate mockgen -package=mocks -destination=../mocks/projects.go -source=projects.go ProjectsClient
type ProjectsClient interface {
	ListProjects(opt *gitlab.ListProjectsOptions, options ...gitlab.RequestOptionFunc) ([]*gitlab.Project, *gitlab.Response, error)
	GetProject(pid interface{}, opt *gitlab.GetProjectOptions, options ...gitlab.RequestOptionFunc) (*gitlab.Project, *gitlab.Response, error)
}
