// Code generated by codegen; DO NOT EDIT.

package docdb

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func OrderableDbInstanceOptions() *schema.Table {
	return &schema.Table{
		Name:        "aws_docdb_orderable_db_instance_options",
		Description: "https://docs.aws.amazon.com/documentdb/latest/developerguide/API_OrderableDBInstanceOption.html",
		Resolver:    fetchDocdbOrderableDbInstanceOptions,
		Multiplex:   client.ServiceAccountRegionMultiplexer("docdb"),
		Columns: []schema.Column{
			{
				Name:     "availability_zones",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("AvailabilityZones"),
			},
			{
				Name:     "db_instance_class",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DBInstanceClass"),
			},
			{
				Name:     "engine",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Engine"),
			},
			{
				Name:     "engine_version",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("EngineVersion"),
			},
			{
				Name:     "license_model",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("LicenseModel"),
			},
			{
				Name:     "vpc",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("Vpc"),
			},
		},
	}
}
