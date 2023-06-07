package client

import (
	"context"
	"fmt"

	resourcemanager "cloud.google.com/go/resourcemanager/apiv3"
	crmv1 "google.golang.org/api/cloudresourcemanager/v1"
)

func (c *Client) resolveLegacy(ctx context.Context, gcpSpec Spec) error {
	projects := gcpSpec.ProjectIDs
	organizations := make([]*crmv1.Organization, 0)

	projectsClient, err := resourcemanager.NewProjectsClient(ctx, c.ClientOptions...)
	if err != nil {
		return fmt.Errorf("failed to create projects client: %w", err)
	}
	foldersClient, err := resourcemanager.NewFoldersClient(ctx, c.ClientOptions...)
	if err != nil {
		return fmt.Errorf("failed to create folders client: %w", err)
	}

	switch {
	case len(projects) == 0 && len(gcpSpec.FolderIDs) == 0 && len(gcpSpec.ProjectFilter) == 0 && len(gcpSpec.FolderFilter) == 0:
		c.logger.Info().Msg("No project_ids, folder_ids, or project_filter specified - assuming all active projects")
		projects, err = getProjectsV1(ctx, c.ClientOptions...)
		if err != nil {
			return fmt.Errorf("failed to get projects: %w", err)
		}
	case len(gcpSpec.FolderIDs) > 0:
		var folderIds []string

		for _, parentFolder := range gcpSpec.FolderIDs {
			c.logger.Info().Msg("Listing folders...")
			childFolders, err := listFolders(ctx, foldersClient, parentFolder, *gcpSpec.FolderRecursionDepth)
			if err != nil {
				return fmt.Errorf("failed to list folders: %w", err)
			}
			folderIds = append(folderIds, childFolders...)
		}

		logFolderIds(&c.logger, folderIds)

		c.logger.Info().Msg("listing folder projects...")
		folderProjects, err := listProjectsInFolders(ctx, projectsClient, folderIds)
		projects = setUnion(projects, folderProjects)
		if err != nil {
			return fmt.Errorf("failed to list projects: %w", err)
		}

	case len(gcpSpec.ProjectFilter) > 0:
		c.logger.Info().Msg("Listing projects with filter...")
		projectsWithFilter, err := getProjectsV1WithFilter(ctx, gcpSpec.ProjectFilter, c.ClientOptions...)
		if err != nil {
			return fmt.Errorf("failed to get projects with filter: %w", err)
		}

		projects = setUnion(projects, projectsWithFilter)
	}

	if len(gcpSpec.OrganizationIDs) == 0 && len(gcpSpec.OrganizationFilter) == 0 {
		c.logger.Info().Msg("No organization_ids or organization_filter specified - assuming all organizations")
		c.logger.Info().Msg("Listing organizations...")

		organizations, err = getOrganizations(ctx, "", c.ClientOptions...)
		if err != nil {
			c.logger.Err(err).Msg("failed to get organizations")
		}
	} else {
		if len(gcpSpec.OrganizationIDs) > 0 {
			for _, orgID := range gcpSpec.OrganizationIDs {
				c.logger.Info().Msgf("Getting spec organization %q...", orgID)
				org, err := getOrganization(ctx, orgID, c.ClientOptions...)
				if err != nil {
					return fmt.Errorf("failed to get spec organization: %w", err)
				}
				organizations = append(organizations, org)
			}
		}
		if len(gcpSpec.OrganizationFilter) > 0 {
			c.logger.Info().Msg("Listing organizations with filter...")
			organizationsWithFilter, err := getOrganizations(ctx, gcpSpec.OrganizationFilter, c.ClientOptions...)
			if err != nil {
				return fmt.Errorf("failed to get organizations with filter: %w", err)
			}
			for i := range organizationsWithFilter {
				found := false
				for _, org := range organizations {
					if organizationsWithFilter[i].Name == org.Name {
						found = true
						break
					}
				}
				if !found {
					organizations = append(organizations, organizationsWithFilter[i])
				}
			}
		}
	}

	logProjectIds(&c.logger, projects)
	logOrganizationIds(&c.logger, organizations)

	if len(projects) == 0 {
		return fmt.Errorf("no active projects")
	}

	c.projects = projects
	c.folderIds = gcpSpec.FolderIDs
	c.orgs = organizations
	if err != nil {
		c.logger.Err(err).Msg("failed to get organizations")
	}
	c.logger.Info().Interface("orgs", c.orgs).Msg("Retrieved organizations")
	return nil
}
