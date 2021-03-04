package client

import (
	"context"
	"errors"
	"github.com/cloudquery/cloudquery/database"
	"github.com/cloudquery/cq-provider-okta/resources"
	"github.com/hashicorp/go-hclog"
	"github.com/okta/okta-sdk-golang/v2/okta"
	"github.com/okta/okta-sdk-golang/v2/okta/query"
)

type Client struct {
	domain string
	okta   *okta.Client
	db     *database.Database
	logger hclog.Logger
}

func New(ctx context.Context, db *database.Database, logger hclog.Logger, domain string, token string) (*Client, error) {
	_, c, err := okta.NewClient(ctx, okta.WithOrgUrl(domain), okta.WithToken(token), okta.WithCache(true))
	if err != nil {
		return nil, err
	}
	return &Client{
		domain: domain,
		okta:   c,
		logger: logger,
		db:     db,
	}, nil
}

func (c Client) FetchApplications(ctx context.Context) error {
	totalCount := 0
	apps, resp, err := c.okta.Application.ListApplications(ctx, query.NewQueryParams(query.WithLimit(200), query.WithAfter("")))
	if err != nil {
		return err
	}
	totalCount += len(apps)
	c.db.ChunkedUpsert(resources.TransformApplications(apps))
	// fetch all app users
	if err := c.fetchAppsAssignedResources(ctx, apps); err != nil {
		return err
	}
	for resp != nil && resp.HasNextPage() {
		var nextAppSet []okta.App
		resp, err = resp.Next(ctx, &nextAppSet)
		if err != nil {
			return err
		}
		totalCount += len(nextAppSet)
		c.db.ChunkedUpsert(resources.TransformApplications(nextAppSet))
		// fetch all app users
		if err := c.fetchAppsAssignedResources(ctx, apps); err != nil {
			return err
		}
	}
	c.logger.Info("Fetched okta users resources", "count", totalCount)
	return nil
}

func (c Client) FetchUsers(ctx context.Context) error {
	totalCount := 0
	users, resp, err := c.okta.User.ListUsers(ctx, query.NewQueryParams(query.WithLimit(200), query.WithAfter("")))
	if err != nil {
		return err
	}
	totalCount += len(users)
	c.db.ChunkedUpsert(resources.TransformUsers(users))
	for resp != nil && resp.HasNextPage() {
		var nextUserSet []*okta.User
		resp, err = resp.Next(ctx, &nextUserSet)
		if err != nil {
			return err
		}
		totalCount += len(nextUserSet)
		c.db.ChunkedUpsert(resources.TransformUsers(nextUserSet))
	}
	c.logger.Info("Fetched okta users resources", "count", totalCount)
	return nil
}

func (c Client) FetchUserTypes(ctx context.Context) error {
	totalCount := 0
	ut, resp, err := c.okta.UserType.ListUserTypes(ctx)
	if err != nil {
		return err
	}
	totalCount += len(ut)
	c.db.ChunkedUpsert(resources.TransformUserTypes(ut))
	for resp != nil && resp.HasNextPage() {
		var nextUserTypeSet []*okta.UserType
		resp, err = resp.Next(ctx, &nextUserTypeSet)
		if err != nil {
			return err
		}
		totalCount += len(nextUserTypeSet)
		c.db.ChunkedUpsert(resources.TransformUserTypes(nextUserTypeSet))
	}
	c.logger.Info("Fetched okta user types resource", "count", totalCount)
	return nil
}

func (c Client) FetchGroups(ctx context.Context) error {
	totalCount := 0
	groups, resp, err := c.okta.Group.ListGroups(ctx, query.NewQueryParams(query.WithLimit(200), query.WithAfter(""), query.WithExpand("app")))
	if err != nil {
		return err
	}
	totalCount += len(groups)
	fetchedGroups := resources.TransformGroups(groups)
	c.db.ChunkedUpsert(fetchedGroups)
	if err := c.queryGroupUsers(ctx, fetchedGroups); err != nil {
		return err
	}

	for resp != nil && resp.HasNextPage() {
		var nextGroupSet []*okta.Group
		resp, err = resp.Next(ctx, &nextGroupSet)
		if err != nil {
			return err
		}
		totalCount += len(nextGroupSet)
		fetchedGroups := resources.TransformGroups(nextGroupSet)
		if err := c.queryGroupUsers(ctx, fetchedGroups); err != nil {
			return err
		}
		c.db.ChunkedUpsert(fetchedGroups)
	}

	c.logger.Info("Fetched okta groups resource", "count", totalCount)
	return nil
}

func (c Client) queryGroupUsers(ctx context.Context, groups []*resources.Group) error {
	for _, g := range groups {
		u, err := c.fetchGroupUsers(ctx, g.Id)
		if err != nil {
			return err
		}
		g.Users = u
	}
	return nil
}

func (c Client) fetchGroupUsers(ctx context.Context, groupId string) ([]*resources.User, error) {
	var fetchedUsers []*resources.User
	var users []*okta.User
	users, resp, err := c.okta.Group.ListGroupUsers(ctx, groupId, query.NewQueryParams(query.WithLimit(200), query.WithAfter("")))
	if err != nil {
		return nil, err
	}
	fetchedUsers = append(fetchedUsers, resources.TransformUsers(users)...)
	for resp != nil && resp.HasNextPage() {
		var nextUserSet []*okta.User
		resp, err = resp.Next(ctx, &nextUserSet)
		if err != nil {
			return nil, err
		}
		fetchedUsers = append(fetchedUsers, resources.TransformUsers(nextUserSet)...)
	}
	c.logger.Info("Fetched okta group's users", "count", len(fetchedUsers), "group", groupId)
	return fetchedUsers, nil
}

func (c Client) fetchAppsAssignedResources(ctx context.Context, aa []okta.App) error {
	for _, a := range aa {
		if !a.IsApplicationInstance() {
			c.logger.Warn("received non application instance")
			return errors.New("unknown app instance type")
		}
		app := a.(*okta.Application)
		if err := c.fetchAppUsers(ctx, app.Id); err != nil {
			return err
		}
		if err := c.fetchAppGroups(ctx, app.Id); err != nil {
			return err
		}
	}
	return nil
}

func (c Client) fetchAppUsers(ctx context.Context, appId string) error {
	users, resp, err := c.okta.Application.ListApplicationUsers(ctx, appId, query.NewQueryParams(query.WithLimit(200), query.WithAfter("")))
	if err != nil {
		return err
	}
	totalAppUsers := len(users)
	c.db.ChunkedUpsert(resources.TransformAppUsers(appId, users))

	for resp != nil && resp.HasNextPage() {
		var nextUserSet []*okta.AppUser
		resp, err = resp.Next(ctx, &nextUserSet)
		if err != nil {
			return err
		}
		totalAppUsers += len(nextUserSet)
		c.db.ChunkedUpsert(resources.TransformAppUsers(appId, nextUserSet))
	}
	c.logger.Info("Fetched okta application's assigned users", "count", totalAppUsers, "app", appId)
	return nil
}

func (c Client) fetchAppGroups(ctx context.Context, appId string) error {
	groups, resp, err := c.okta.Application.ListApplicationGroupAssignments(ctx, appId, query.NewQueryParams(query.WithLimit(200), query.WithAfter("")))
	if err != nil {
		return err
	}
	totalAppGroups := len(groups)
	c.db.ChunkedUpsert(resources.TransformAppGroups(appId, groups))

	for resp != nil && resp.HasNextPage() {
		var nextGroupSet []*okta.ApplicationGroupAssignment
		resp, err = resp.Next(ctx, &nextGroupSet)
		if err != nil {
			return err
		}
		totalAppGroups += len(nextGroupSet)
		c.db.ChunkedUpsert(resources.TransformAppGroups(appId, nextGroupSet))
	}
	c.logger.Info("Fetched okta application's assigned groups", "count", totalAppGroups, "app", appId)
	return nil
}
