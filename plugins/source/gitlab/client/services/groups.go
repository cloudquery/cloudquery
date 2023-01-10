package services

import "github.com/xanzy/go-gitlab"

type GroupsClient interface {
	ListGroupMembers(gid any, opt *gitlab.ListGroupMembersOptions, options ...gitlab.RequestOptionFunc) ([]*gitlab.GroupMember, *gitlab.Response, error)
	ListGroups(opt *gitlab.ListGroupsOptions, options ...gitlab.RequestOptionFunc) ([]*gitlab.Group, *gitlab.Response, error)
	ListGroupProjects(gid any, opt *gitlab.ListGroupProjectsOptions, options ...gitlab.RequestOptionFunc) ([]*gitlab.Project, *gitlab.Response, error)
}
