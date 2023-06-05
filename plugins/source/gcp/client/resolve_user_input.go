package client

import (
	"context"
	"fmt"

	resourcemanager "cloud.google.com/go/resourcemanager/apiv3"
	"github.com/thoas/go-funk"
	crmv1 "google.golang.org/api/cloudresourcemanager/v1"
)

func (c *Client) getProjects(ctx context.Context, gcpSpec Spec) error {

	if !gcpSpec.Projects.Organizations.isNull() {
		var organizationsInclude, organizationsExclude []*crmv1.Organization
		var err error
		service, err := crmv1.NewService(ctx, c.ClientOptions...)
		if err != nil {
			return fmt.Errorf("failed to create cloudresourcemanager service: %w", err)
		}

		if gcpSpec.Projects.Organizations.include_filter != "" {
			organizationsInclude, err = getOrganizationsFilter(ctx, service, gcpSpec.Projects.Organizations.include_filter)
			if err != nil {
				return fmt.Errorf("failed to get organizations with filter: %w", err)
			}
		}
		if gcpSpec.Projects.Organizations.exclude_filter != "" {
			organizationsExclude, err = getOrganizationsFilter(ctx, service, gcpSpec.Projects.Organizations.exclude_filter)
			if err != nil {
				return fmt.Errorf("failed to get organizations with filter: %w", err)
			}
		}
		// Resolve organization from gcpSpec.Projects.Organizations.id_include_list and add to organizationsInclude
		for _, orgId := range gcpSpec.Projects.Organizations.id_include_list {
			org, err := getOrganizationFromId(ctx, service, orgId)
			if err != nil {
				return fmt.Errorf("failed to get organization with id %s: %w", orgId, err)
			}
			organizationsInclude = append(organizationsInclude, org)
		}
		// Resolve organization from gcpSpec.Projects.Organizations.id_exclude_list and add to organizationsExclude

		for _, orgId := range gcpSpec.Projects.Organizations.id_exclude_list {
			org, err := getOrganizationFromId(ctx, service, orgId)
			if err != nil {
				return fmt.Errorf("failed to get organization with id %s: %w", orgId, err)
			}
			organizationsExclude = append(organizationsExclude, org)
		}

		// remove all organizations in organizationsExclude from organizationsInclude
		for _, orgInclude := range organizationsInclude {
			match := false
			for _, orgExclude := range organizationsExclude {
				if compareOrgs(orgInclude, orgExclude) {
					match = true
					break
				}
			}
			if !match {
				c.orgs = append(c.orgs, orgInclude)
			}

		}

	}
	if !gcpSpec.Projects.Folders.isNull() {
		var foldersInclude, foldersExclude []string
		var err error
		foldersClient, err := resourcemanager.NewFoldersClient(ctx, c.ClientOptions...)

		if err != nil {
			return fmt.Errorf("failed to create folders client: %w", err)
		}
		if gcpSpec.Projects.Folders.include_filter != "" {
			foldersInclude, err = searchFolders(ctx, foldersClient, gcpSpec.Projects.Organizations.include_filter)
			if err != nil {
				return fmt.Errorf("failed to get organizations with filter: %w", err)
			}
		}
		if gcpSpec.Projects.Organizations.exclude_filter != "" {
			foldersExclude, err = searchFolders(ctx, foldersClient, gcpSpec.Projects.Organizations.exclude_filter)
			if err != nil {
				return fmt.Errorf("failed to get organizations with filter: %w", err)
			}
		}
		// combine all folders found with a filter and the hard coded ids supplied by user
		foldersInclude = append(foldersInclude, gcpSpec.Projects.Folders.id_include_list...)
		foldersExclude = append(foldersExclude, gcpSpec.Projects.Folders.id_exclude_list...)

		// Ensure all folder ids in each list are unique
		foldersInclude = funk.UniqString(foldersInclude)
		foldersExclude = funk.UniqString(foldersExclude)

		// Subtract the exclude from the include lists
		foldersInclude = funk.SubtractString(foldersInclude, foldersExclude)
		c.folderIds = foldersInclude
	}

	// if !gcpSpec.Projects.Projects.isNull() {}

	return nil
}

func getOrganizationsFilter(ctx context.Context, service *crmv1.Service, filter string) ([]*crmv1.Organization, error) {
	organizationsWithFilter := make([]*crmv1.Organization, 0)
	if err := service.Organizations.Search(&crmv1.SearchOrganizationsRequest{Filter: filter}).Context(ctx).Pages(ctx, func(page *crmv1.SearchOrganizationsResponse) error {
		organizationsWithFilter = append(organizationsWithFilter, page.Organizations...)
		return nil
	}); err != nil {
		return nil, err
	}
	return organizationsWithFilter, nil
}

func getOrganizationFromId(ctx context.Context, service *crmv1.Service, id string) (*crmv1.Organization, error) {
	return service.Organizations.Get("organizations/" + id).Context(ctx).Do()
}

func compareOrgs(org1, org2 *crmv1.Organization) bool {
	if org1.Name != org2.Name {
		return false
	}
	if org1.CreationTime != org2.CreationTime {
		return false
	}
	return true
}
