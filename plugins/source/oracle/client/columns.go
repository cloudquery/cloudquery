package client

import (
	"github.com/apache/arrow/go/v14/arrow"
	"github.com/cloudquery/plugin-sdk/v4/schema"
)

var (
	RegionColumn = schema.Column{
		Name:       "region",
		Type:       arrow.BinaryTypes.String,
		Resolver:   ResolveOracleRegion,
		PrimaryKey: true,
	}
	CompartmentIDColumn = schema.Column{
		Name:       "compartment_id",
		Type:       arrow.BinaryTypes.String,
		Resolver:   ResolveCompartmentID,
		PrimaryKey: true,
	}
	AvailabilityDomainColumn = schema.Column{
		Name:       "availability_domain",
		Type:       arrow.BinaryTypes.String,
		Resolver:   ResolveAvailabilityDomain,
		PrimaryKey: true,
	}
)
