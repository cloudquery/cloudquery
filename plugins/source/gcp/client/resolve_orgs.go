package client

import (
	"context"
	"fmt"

	crmv1 "google.golang.org/api/cloudresourcemanager/v1"
)

func (c *Client) resolveOrgs(ctx context.Context, organization ResourceDiscovery) error {
	var err error
	service, err := crmv1.NewService(ctx, c.ClientOptions...)
	if err != nil {
		return fmt.Errorf("failed to create cloudresourcemanager service: %w", err)
	}
	var includedOrgs, excludedOrgs []*crmv1.Organization
	if !organization.isNull() {
		for _, includeFilter := range organization.IncludeFilter {
			orgs, err := getOrganizationsFilter(ctx, service, includeFilter)
			if err != nil {
				return fmt.Errorf("failed to get organizations with filter: %w", err)
			}
			includedOrgs = append(includedOrgs, orgs...)
		}
		for _, excludeFilter := range organization.ExcludeFilter {
			orgs, err := getOrganizationsFilter(ctx, service, excludeFilter)
			if err != nil {
				return fmt.Errorf("failed to get organizations with filter: %w", err)
			}
			excludedOrgs = append(excludedOrgs, orgs...)
		}
		// Resolve organization from gcpSpec.Projects.Organizations.id_include_list and add to includedOrgs
		for _, orgId := range organization.IncludeListId {
			org, err := getOrganizationFromId(ctx, service, orgId)
			if err != nil {
				return fmt.Errorf("failed to get organization with id %s: %w", orgId, err)
			}
			includedOrgs = append(includedOrgs, org)
		}
		// Resolve organization from gcpSpec.Projects.Organizations.id_exclude_list and add to excludedOrgs
		for _, orgId := range organization.ExcludeListId {
			org, err := getOrganizationFromId(ctx, service, orgId)
			if err != nil {
				return fmt.Errorf("failed to get organization with id %s: %w", orgId, err)
			}
			excludedOrgs = append(excludedOrgs, org)
		}
	}
	if organization.isIncludeNull() {
		orgs, err := getOrganizationsFilter(ctx, service, "name:*")
		if err != nil {
			return fmt.Errorf("failed to get organizations with filter: %w", err)
		}
		includedOrgs = append(includedOrgs, orgs...)
	}
	if c.graph == nil {
		c.graph = &node{}
	}
	trueBool := true
	for _, orgId := range includedOrgs {
		c.graph.relations = append(c.graph.relations, &node{
			org:      orgId,
			included: &trueBool,
		})
	}
	falseBool := false
	for _, orgId := range excludedOrgs {
		c.graph.relations = append(c.graph.relations, &node{
			org:      orgId,
			included: &falseBool,
		})
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
