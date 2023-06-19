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

	// If no projects are included then include all projects
	if project.isIncludeNull() {
		c.logger.Info().Msg("no projects specified in filter or include_list so assuming all projects")
		project.IncludeFilter = []string{"lifecycleState=ACTIVE"}
	}
	for _, includeFilter := range project.IncludeFilter {
		projects, err := listProjectsFilter(ctx, service, includeFilter)
		if err != nil {
			return fmt.Errorf("failed to list projects with filter (%s): %w", includeFilter, err)
		}
		for _, project := range projects {
			if !addProject(c.graph, project, &boolTrue) {
				c.logger.Warn().Msgf("project %s is included but could not be added to the dependency graph", project.Name)
			}
		}
	}

	for _, includeId := range project.IncludeListId {
		project, err := service.Projects.Get(includeId).Context(ctx).Do()
		if err != nil {
			return fmt.Errorf("failed to get project with id %s: %w", includeId, err)
		}
		if !addProject(c.graph, project, &boolTrue) {
			c.logger.Warn().Msgf("project %s is included but could not be added to the dependency graph", project.Name)
		}
	}
	for _, excludeId := range project.ExcludeListId {
		project, err := service.Projects.Get(excludeId).Context(ctx).Do()
		if err != nil {
			return fmt.Errorf("failed to get project with id %s: %w", excludeId, err)
		}
		if !addProject(c.graph, project, &boolFalse) {
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
