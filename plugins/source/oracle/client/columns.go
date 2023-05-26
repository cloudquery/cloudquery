package client

import (
	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/plugin-sdk/v3/schema"
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
)
