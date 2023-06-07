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
	var includedProjects, excludedProjects []*crmv1.Project

	if project.isIncludeNull() {
		includedProjects, err = listProjectsFilter(ctx, service, "lifecycleState=ACTIVE")
		if err != nil {
			return fmt.Errorf("failed to list active projects: %w", err)
		}
	}
	for _, includeFilter := range project.IncludeFilter {
		projects, err := listProjectsFilter(ctx, service, includeFilter)
		if err != nil {
			return fmt.Errorf("failed to list projects with filter (%s): %w", includeFilter, err)
		}
		includedProjects = append(includedProjects, projects...)
	}
	for _, excludeFilter := range project.ExcludeFilter {
		projects, err := listProjectsFilter(ctx, service, excludeFilter)
		if err != nil {
			return fmt.Errorf("failed to list projects with filter (%s): %w", excludeFilter, err)
		}
		excludedProjects = append(excludedProjects, projects...)
	}
	for _, includeId := range project.IncludeListId {
		project, err := service.Projects.Get(includeId).Context(ctx).Do()
		if err != nil {
			return fmt.Errorf("failed to get project with id %s: %w", includeId, err)
		}
		includedProjects = append(includedProjects, project)
	}
	for _, excludeId := range project.ExcludeListId {
		project, err := service.Projects.Get(excludeId).Context(ctx).Do()
		if err != nil {
			return fmt.Errorf("failed to get project with id %s: %w", excludeId, err)
		}
		excludedProjects = append(excludedProjects, project)
	}
	trueBool := true
	for _, project := range includedProjects {
		addProject(c.graph, project, &trueBool)
		if !addProject(c.graph, project, &trueBool) {
			c.logger.Warn().Msgf("project %s is included but could not be added to the dependency graph", project.Name)
		}
	}
	falseBool := false
	for _, project := range excludedProjects {
		if !addProject(c.graph, project, &falseBool) {
			c.logger.Warn().Msgf("project %s is excluded but could not be added to the dependency graph", project.Name)
		}
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
