package client

import (
	"context"
	"fmt"
	"strings"

	"cloud.google.com/go/resourcemanager/apiv3/resourcemanagerpb"
	"github.com/thoas/go-funk"
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
	c.projects = findAllIncludedProjects(c.graph, []string{})
	c.folderIds = findAllIncludedFolders(c.graph, []string{})
	c.orgs = findAllIncludedOrganizations(c.graph, []*crmv1.Organization{})

	return nil
}

type node struct {
	parentID  *string
	included  *bool
	relations nodes
	project   *crmv1.Project
	folder    *resourcemanagerpb.Folder
	org       *crmv1.Organization
}

type nodes []*node

func findNodeByParentID(node *node, id string) *node {
	if node.folder != nil && node.folder.Name == id {
		return node
	}
	if node.org != nil && node.org.Name == id {
		return node
	}

	if len(node.relations) > 0 {
		for _, child := range node.relations {
			cNode := findNodeByParentID(child, id)
			if cNode != nil {
				return cNode
			}
		}
	}
	return nil
}

func addFolder(topNode *node, folder *resourcemanagerpb.Folder, toBeIncluded *bool) bool {
	existingNode := findNodeByParentID(topNode, folder.Parent)
	if existingNode == nil {
		return false
	}
	for _, existingFolder := range existingNode.relations {
		if existingFolder.folder.Name == folder.Name {
			return false
		}
	}
	newNode := node{
		parentID: &folder.Parent,
		included: toBeIncluded,
		folder:   folder,
	}

	existingNode.relations = append(existingNode.relations, &newNode)
	return true
}

func updateFolder(topNode *node, folder *resourcemanagerpb.Folder, toBeIncluded *bool) bool {
	existingNode := findNodeByParentID(topNode, folder.Name)
	if existingNode == nil {
		return false
	}
	for _, existingFolder := range existingNode.relations {
		if existingFolder.folder.Name == folder.Name {
			existingFolder.included = toBeIncluded
			return true
		}
	}
	return true
}

func addProject(topNode *node, project *crmv1.Project, toBeIncluded *bool) bool {
	parentID := strings.ToLower(project.Parent.Type) + "s/" + strings.ToLower(project.Parent.Id)
	existingNode := findNodeByParentID(topNode, parentID)
	if existingNode == nil {
		return false
	}
	for _, existingProject := range existingNode.relations {
		if existingProject.project != nil && existingProject.project.ProjectId == project.ProjectId {
			return false
		}
	}
	newNode := node{
		parentID: &project.Parent.Id,
		included: toBeIncluded,
		project:  project,
	}

	existingNode.relations = append(existingNode.relations, &newNode)
	return true
}

func findAllIncludedProjects(topNode *node, results []string) []string {
	if topNode.project != nil && *topNode.included {
		results = funk.UniqString(append(results, topNode.project.ProjectId))
	}
	for _, child := range topNode.relations {
		// Only check relations if the parent is included
		if child.included == nil || *child.included {
			results = funk.UniqString(append(results, findAllIncludedProjects(child, results)...))
		}
	}
	return funk.UniqString(results)
}

func findAllIncludedFolders(topNode *node, results []string) []string {
	if topNode.folder != nil && topNode.included != nil && *topNode.included {
		results = funk.UniqString(append(results, topNode.folder.Name))
	}
	for _, child := range topNode.relations {
		// Only check relations if the parent is included
		if child.included == nil || *child.included {
			results = funk.UniqString(append(results, findAllIncludedFolders(child, results)...))
		}
	}
	return funk.UniqString(results)
}

func findAllIncludedOrganizations(topNode *node, results []*crmv1.Organization) []*crmv1.Organization {
	if topNode.org != nil && *topNode.included {
		results = uniqOrg(append(results, topNode.org))
	}
	for _, child := range topNode.relations {
		// Only check relations if the parent is included
		if child.included == nil || *child.included {
			results = uniqOrg(append(results, findAllIncludedOrganizations(child, results)...))
		}
	}
	return uniqOrg(results)
}

func uniqOrg(orgs []*crmv1.Organization) []*crmv1.Organization {
	uniq := []*crmv1.Organization{}
	for _, org := range orgs {
		if !funk.Contains(uniq, org) {
			uniq = append(uniq, org)
		}
	}
	return uniq
}
