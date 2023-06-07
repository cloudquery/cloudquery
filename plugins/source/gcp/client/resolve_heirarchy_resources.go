package client

import (
	"context"
	"fmt"

	crmv1 "google.golang.org/api/cloudresourcemanager/v1"
)

func (c *Client) resolveDiscovery(ctx context.Context, gcpSpec Spec) error {
	if err := c.resolveOrgs(ctx, gcpSpec.HierarchyDiscovery.Organizations); err != nil {
		return fmt.Errorf("failed to resolve organizations: %w", err)
	}
	if err := c.resolveFolders(ctx, gcpSpec.HierarchyDiscovery.Folders); err != nil {
		return fmt.Errorf("failed to resolve organizations: %w", err)
	}
	if err := c.resolveProjects(ctx, gcpSpec.HierarchyDiscovery.Projects); err != nil {
		return fmt.Errorf("failed to resolve organizations: %w", err)
	}

	return c.filterResources()
}

func (c *Client) filterResources() error {
	c.projects = findAllIncludedProjects(c.graph, []string{})
	c.folderIds = findAllIncludedFolders(c.graph, []string{})
	c.orgs = findAllIncludedOrganizations(c.graph, []*crmv1.Organization{})

	return nil
}
