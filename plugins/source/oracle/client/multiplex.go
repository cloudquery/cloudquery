package client

import "github.com/cloudquery/plugin-sdk/schema"

func RegionCompartmentMultiplex(meta schema.ClientMeta) []schema.ClientMeta {
	cqClient := meta.(*Client)

	multiplexedClients := make([]schema.ClientMeta, 0, len(cqClient.OracleClients)*len(cqClient.AllCompartmentOcids))

	// The `OracleClients` map is keyed by region.
	for region := range cqClient.OracleClients {
		for _, compartmentOcid := range cqClient.AllCompartmentOcids {
			multiplexedClients = append(multiplexedClients, cqClient.withRegion(region).withCompartment(compartmentOcid))
		}
	}

	return multiplexedClients
}

// Returns a new cq-client for the home-region and root-compartment.
// root-compartment - tenancy ocid.
func TenancyMultiplex(meta schema.ClientMeta) []schema.ClientMeta {
	cqClient := meta.(*Client)

	multiplexedClient := cqClient.withRegion(cqClient.HomeRegion).withCompartment(cqClient.TenancyOcid)

	return []schema.ClientMeta{
		multiplexedClient,
	}
}

func AvailibilityDomainCompartmentMultiplex(meta schema.ClientMeta) []schema.ClientMeta {
	cqClient := meta.(*Client)

	multiplexedClients := make([]schema.ClientMeta, 0)

	for region, availibilityDomains := range cqClient.RegionAvaililbilityDomainMap {
		for _, availibilityDomain := range availibilityDomains {
			for _, compartmentOcid := range cqClient.AllCompartmentOcids {
				multiplexedClients = append(multiplexedClients, cqClient.
					withRegion(region).
					withCompartment(compartmentOcid).
					withAvailibilityDomain(availibilityDomain))
			}
		}
	}

	return multiplexedClients
}
