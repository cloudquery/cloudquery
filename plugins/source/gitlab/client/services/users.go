package services

import "github.com/xanzy/go-gitlab"

//go:generate mockgen -package=mocks -destination=../mocks/users.go -source=users.go UsersClient
type UsersClient interface {
	GetUser(user int, opt gitlab.GetUsersOptions, options ...gitlab.RequestOptionFunc) (*gitlab.User, *gitlab.Response, error)
}
