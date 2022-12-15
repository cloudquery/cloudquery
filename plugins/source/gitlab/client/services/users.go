package services

import "github.com/xanzy/go-gitlab"

type UsersClient interface {
	GetUser(user int, opt gitlab.GetUsersOptions, options ...gitlab.RequestOptionFunc) (*gitlab.User, *gitlab.Response, error)
	ListUsers(opt *gitlab.ListUsersOptions, options ...gitlab.RequestOptionFunc) ([]*gitlab.User, *gitlab.Response, error)
}
