package client

import (
	"context"
	"fmt"

	"cloud.google.com/go/resourcemanager/apiv3/resourcemanagerpb"
	crmv1 "google.golang.org/api/cloudresourcemanager/v1"
)

func (c *Client) resolveDiscovery(ctx context.Context, gcpSpec Spec) error {
	if err := c.resolveOrgs(ctx, gcpSpec.Projects.Organizations); err != nil {
		return fmt.Errorf("failed to resolve organizations: %w", err)
	}
	if err := c.resolveFolders(ctx, gcpSpec.Projects.Folders); err != nil {
		return fmt.Errorf("failed to resolve organizations: %w", err)
	}
	if err := c.resolveProjects(ctx, gcpSpec.Projects.Projects); err != nil {
		return fmt.Errorf("failed to resolve organizations: %w", err)
	}
	return c.filterResources()
}

func (c *Client) filterResources() error {
	// check if any included folders are a descendent of an excluded organization, if so, then move the folder to the excludedFolder list
	// check if any included folder is a descendent of any excluded folder, if so, remove the folder from the includedFolder list
	// check if any included project is a descendent of any excluded organization, if so, move the project to the excludedProject list
	// check if any included project is a descendent of any excluded folder, if so, move the project to the excludedProject list

	return nil
}

type node struct {
	parentID  *string
	included  bool
	relations nodes
	// project   *crmv1.Project
	folder *resourcemanagerpb.Folder
	org    *crmv1.Organization
}

type nodes []*node

func findNodeByID(node *node, id string) *node {
	if node.folder != nil && node.folder.Name == id {
		return node
	}
	if node.org != nil && node.org.Name == id {
		return node
	}

	if len(node.relations) > 0 {
		for _, child := range node.relations {
			findNodeByID(child, id)
		}
	}
	return nil
}
