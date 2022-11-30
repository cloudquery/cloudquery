package services

import "github.com/xanzy/go-gitlab"

//go:generate mockgen -package=mocks -destination=../mocks/groups.go -source=groups.go GroupsClient
type GroupsClient interface {
	ListGroupMembers(gid interface{}, opt *gitlab.ListGroupMembersOptions, options ...gitlab.RequestOptionFunc) ([]*gitlab.GroupMember, *gitlab.Response, error)
	ListGroups(opt *gitlab.ListGroupsOptions, options ...gitlab.RequestOptionFunc) ([]*gitlab.Group, *gitlab.Response, error)
}
