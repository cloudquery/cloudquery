package client

import (
	"context"
	"fmt"

	crmv1 "google.golang.org/api/cloudresourcemanager/v1"
)

func (c *Client) resolveOrgs(ctx context.Context, org ResourceDiscovery) error {
	var err error
	service, err := crmv1.NewService(ctx, c.ClientOptions...)
	if err != nil {
		return fmt.Errorf("failed to create cloudresourcemanager service: %w", err)
	}
	if !org.isNull() {
		for _, includeFilter := range org.IncludeFilter {
			orgs, err := getOrganizationsFilter(ctx, service, includeFilter)
			if err != nil {
				return fmt.Errorf("failed to get organizations with filter: %w", err)
			}
			c.included_orgs = append(c.included_orgs, orgs...)
		}
		for _, excludeFilter := range org.ExcludeFilter {
			orgs, err := getOrganizationsFilter(ctx, service, excludeFilter)
			if err != nil {
				return fmt.Errorf("failed to get organizations with filter: %w", err)
			}
			c.excluded_orgs = append(c.excluded_orgs, orgs...)

		}
		// Resolve organization from gcpSpec.Projects.Organizations.id_include_list and add to c.included_orgs
		for _, orgId := range org.IncludeListId {
			org, err := getOrganizationFromId(ctx, service, orgId)
			if err != nil {
				return fmt.Errorf("failed to get organization with id %s: %w", orgId, err)
			}
			c.included_orgs = append(c.included_orgs, org)
		}
		// Resolve organization from gcpSpec.Projects.Organizations.id_exclude_list and add to c.excluded_orgs
		for _, orgId := range org.ExcludeListId {
			org, err := getOrganizationFromId(ctx, service, orgId)
			if err != nil {
				return fmt.Errorf("failed to get organization with id %s: %w", orgId, err)
			}
			c.excluded_orgs = append(c.excluded_orgs, org)
		}
	}
	return nil
}

func getOrganizationsFilter(ctx context.Context, service *crmv1.Service, filter string) ([]*crmv1.Organization, error) {
	organizationsWithFilter := make([]*crmv1.Organization, 0)

	input := &crmv1.SearchOrganizationsRequest{}
	if filter == "" {
		input.Filter = filter
	}

	if err := service.Organizations.Search(input).Context(ctx).Pages(ctx, func(page *crmv1.SearchOrganizationsResponse) error {
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
