package client

import (
	"context"
	"fmt"
	"strings"

	"github.com/cloudquery/plugin-sdk/plugins/source"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/specs"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/identity"
	"github.com/oracle/oci-go-sdk/v65/objectstorage"
	"github.com/rs/zerolog"
)

type Client struct {
	// A map of region->`OracleClients` struct.
	// Every OracleClients struct contains all the clients we need for a single regionXcompartment
	OracleClients       map[string]*OracleClients
	AllCompartmentOcids []string

	TenancyOcid string // Tenancy == RootCompartment
	HomeRegion  string

	ObjectStorageNamespace string // A global value, used for object-storage (i.e. buckets)

	// All availibility domains in the tenancy.
	RegionAvaililbilityDomainMap map[string][]string

	// These are different per "cq-client", i.e. per multiplexed-cq-client.
	// By default (if no multiplexer is defined), Region is set to the home region, and CompartmentOcid is set to the tenancy ocid.
	Region             string
	CompartmentOcid    string
	AvailibilityDomain string // For fetches that are multiplexed by availibility domain (and not region)

	logger zerolog.Logger
}

func Configure(ctx context.Context, logger zerolog.Logger, s specs.Source, _ source.Options) (schema.ClientMeta, error) {
	configProvider := common.DefaultConfigProvider()

	tenancyOcid, err := configProvider.TenancyOCID()
	if err != nil {
		return nil, err
	}

	// The identity-client in the home region, that we use for initialization.
	homeIdentityClient, err := identity.NewIdentityClientWithConfigurationProvider(configProvider)
	if err != nil {
		return nil, err
	}

	allCompartmentOcids, err := getAllCompartmentIdsInTenancy(ctx, homeIdentityClient, tenancyOcid)
	if err != nil {
		return nil, err
	}
	logger.Info().Int("num_compartments", len(allCompartmentOcids)).Msg("syncing from all compartments in tenancy")

	allRegions, homeRegion, err := getAllSubscribedRegions(ctx, homeIdentityClient, tenancyOcid)
	if err != nil {
		return nil, err
	}
	logger.Info().Int("num_regions", len(allRegions)).Msg("syncing from all subscribed regions")

	oracleClients, err := initOracleClientsInAllRegions(configProvider, allRegions)
	if err != nil {
		return nil, err
	}

	objectStorageNamespace, err := getObjectStorageNamespace(ctx, oracleClients[homeRegion].ObjectstorageObjectstorageClient)

	if err != nil {
		return nil, err
	}

	regionAvailibilityDomainMap, err := getRegionAvailibilityDomainMap(ctx, oracleClients, tenancyOcid)
	if err != nil {
		return nil, err
	}

	logger = logger.With().Str("region", homeRegion).Str("compartment_ocid", tenancyOcid).Logger()

	return &Client{
		OracleClients:                oracleClients,
		AllCompartmentOcids:          allCompartmentOcids,
		HomeRegion:                   homeRegion,
		TenancyOcid:                  tenancyOcid,
		RegionAvaililbilityDomainMap: regionAvailibilityDomainMap,
		ObjectStorageNamespace:       objectStorageNamespace,
		Region:                       homeRegion,  // Default value if no multiplexer is defined
		CompartmentOcid:              tenancyOcid, // Default value if no multiplexer is defined
		logger:                       logger,
	}, nil
}

// cq-client id is comprised of region and compartment-ocid.
func (c *Client) ID() string {
	return strings.Join([]string{c.Region, c.CompartmentOcid}, ":")
}

func (c *Client) Logger() *zerolog.Logger {
	return &c.logger
}

func (c *Client) withRegion(region string) *Client {
	newClient := *c
	newClient.logger = c.logger.With().Str("region", region).Logger()
	newClient.Region = region
	return &newClient
}

func (c *Client) withCompartment(compartmentOcid string) *Client {
	newClient := *c
	newClient.logger = c.logger.With().Str("compartment_ocid", compartmentOcid).Logger()
	newClient.CompartmentOcid = compartmentOcid
	return &newClient
}

func (c *Client) withAvailibilityDomain(availabilityDomain string) *Client {
	newClient := *c
	newClient.logger = c.logger.With().Str("availability_domain", availabilityDomain).Logger()
	newClient.AvailibilityDomain = availabilityDomain
	return &newClient
}

// Returns a list of all compartment-ids in a tenancy, recursively. (including the tenancy itself, which is the 'root compartment').
func getAllCompartmentIdsInTenancy(ctx context.Context, client identity.IdentityClient, tenancyId string) ([]string, error) {
	allCompartmentIds := []string{tenancyId}

	var page *string
	for {
		request := identity.ListCompartmentsRequest{
			CompartmentId:          common.String(tenancyId),
			CompartmentIdInSubtree: common.Bool(true), // recursively traverse the compartment hierarchy
			LifecycleState:         identity.CompartmentLifecycleStateActive,
			Page:                   page,
		}

		response, err := client.ListCompartments(ctx, request)
		if err != nil {
			return nil, fmt.Errorf("failed to list compartments in tenancy: %w", err)
		}

		for i := range response.Items {
			allCompartmentIds = append(allCompartmentIds, *response.Items[i].Id)
		}

		page = response.OpcNextPage
		if response.OpcNextPage == nil {
			break
		}
	}

	return allCompartmentIds, nil
}

// Returns a list of all subscribed regions in a tenancy. Also returns the name of the home-region.
func getAllSubscribedRegions(ctx context.Context, client identity.IdentityClient, tenancyOcid string) ([]string, string, error) {
	request := identity.ListRegionSubscriptionsRequest{
		TenancyId: common.String(tenancyOcid),
	}

	response, err := client.ListRegionSubscriptions(ctx, request)
	if err != nil {
		return nil, "", err
	}

	regionNames := make([]string, 0)
	var homeRegion string

	for _, region := range response.Items {
		if region.Status != identity.RegionSubscriptionStatusReady {
			continue
		}

		regionNames = append(regionNames, *region.RegionName)
		if *region.IsHomeRegion {
			homeRegion = *region.RegionName
		}
	}

	if homeRegion == "" {
		return nil, "", fmt.Errorf("no home region found")
	}

	return regionNames, homeRegion, nil
}

func getObjectStorageNamespace(ctx context.Context, client *objectstorage.ObjectStorageClient) (string, error) {
	request := objectstorage.GetNamespaceRequest{}

	response, err := client.GetNamespace(ctx, request)
	if err != nil {
		return "", err
	}

	return *response.Value, nil
}

func getRegionAvailibilityDomainMap(ctx context.Context, oracleClients map[string]*OracleClients, tenancyOcid string) (map[string][]string, error) {
	regionAvailibilityDomainMap := make(map[string][]string)

	for region, clients := range oracleClients {
		request := identity.ListAvailabilityDomainsRequest{
			CompartmentId: common.String(tenancyOcid),
		}

		response, err := clients.IdentityIdentityClient.ListAvailabilityDomains(ctx, request)
		if err != nil {
			return nil, err
		}

		availabilityDomains := make([]string, 0)

		for _, domain := range response.Items {
			availabilityDomains = append(availabilityDomains, *domain.Name)
		}

		regionAvailibilityDomainMap[region] = availabilityDomains
	}

	return regionAvailibilityDomainMap, nil
}
