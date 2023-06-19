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
	// var includedOrgs, excludedOrgs []*crmv1.Organization
	if c.graph == nil {
		c.graph = &node{}
	}

	// If no orgs are included then include all orgs
	if organization.isIncludeNull() {
		c.logger.Info().Msg("no organizations specified in filter or include_list so assuming all orgs")
		organization.IncludeFilter = []string{"name:*"}
	}

	for _, includeFilter := range organization.IncludeFilter {
		orgs, err := getOrganizationsFilter(ctx, service, includeFilter)
		if err != nil {
			return fmt.Errorf("failed to get organizations with filter: %w", err)
		}
		for _, org := range orgs {
			if !addOrg(c.graph, org, &boolTrue) {
				c.logger.Warn().Msgf("organization %s is excluded but could not be added to the dependency graph", org.Name)
			}
		}
	}

	for _, orgId := range organization.IncludeListId {
		org, err := getOrganizationFromId(ctx, service, orgId)
		if err != nil {
			return fmt.Errorf("failed to get organization with id %s: %w", orgId, err)
		}
		if !addOrg(c.graph, org, &boolTrue) {
			c.logger.Warn().Msgf("organization %s is excluded but could not be added to the dependency graph", org.Name)
		}
	}
	for _, orgId := range organization.ExcludeListId {
		org, err := getOrganizationFromId(ctx, service, orgId)
		if err != nil {
			return fmt.Errorf("failed to get organization with id %s: %w", orgId, err)
		}
		if !addOrg(c.graph, org, &boolFalse) {
			c.logger.Warn().Msgf("organization %s is excluded but could not be added to the dependency graph", org.Name)
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
