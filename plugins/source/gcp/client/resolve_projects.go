package client

import (
	"context"
	"fmt"

	crmv1 "google.golang.org/api/cloudresourcemanager/v1"
)

func (c *Client) resolveProjects(ctx context.Context, project ResourceDiscovery) error {
	var err error
	service, err := crmv1.NewService(ctx, c.ClientOptions...)
	if err != nil {
		return fmt.Errorf("failed to create cloudresourcemanager service: %w", err)
	}

	if project.isIncludeNull() {
		c.include_projects, err = listProjectsFilter(ctx, service, "lifecycleState=ACTIVE")
		if err != nil {
			return fmt.Errorf("failed to list active projects: %w", err)
		}
	}
	for _, includeFilter := range project.IncludeFilter {
		projects, err := listProjectsFilter(ctx, service, includeFilter)
		if err != nil {
			return fmt.Errorf("failed to list projects with filter (%s): %w", includeFilter, err)
		}
		c.include_projects = append(c.include_projects, projects...)
	}
	for _, excludeFilter := range project.ExcludeFilter {
		projects, err := listProjectsFilter(ctx, service, excludeFilter)
		if err != nil {
			return fmt.Errorf("failed to list projects with filter (%s): %w", excludeFilter, err)
		}
		c.exclude_projects = append(c.exclude_projects, projects...)
	}
	for _, includeId := range project.IncludeListId {
		project, err := service.Projects.Get(includeId).Context(ctx).Do()
		if err != nil {
			return fmt.Errorf("failed to get project with id %s: %w", includeId, err)
		}
		c.include_projects = append(c.include_projects, project)
	}
	for _, excludeId := range project.ExcludeListId {
		project, err := service.Projects.Get(excludeId).Context(ctx).Do()
		if err != nil {
			return fmt.Errorf("failed to get project with id %s: %w", excludeId, err)
		}
		c.exclude_projects = append(c.exclude_projects, project)
	}

	return nil
}

func listProjectsFilter(ctx context.Context, service *crmv1.Service, filter string) ([]*crmv1.Project, error) {
	var projects []*crmv1.Project
	call := service.Projects.List().Filter(filter).Context(ctx)
	for {
		output, err := call.Do()
		if err != nil {
			return nil, err
		}
		projects = append(projects, output.Projects...)
		if output.NextPageToken == "" {
			break
		}
		call.PageToken(output.NextPageToken)
	}
	return projects, nil
}
