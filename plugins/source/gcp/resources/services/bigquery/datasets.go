// Code generated by codegen; DO NOT EDIT.

package bigquery

import (
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugins/source/gcp/client"
)

func Datasets() *schema.Table {
	return &schema.Table{
		Name:                "gcp_bigquery_datasets",
		Resolver:            fetchDatasets,
		PreResourceResolver: datasetGet,
		Multiplex:           client.ProjectMultiplex,
		Columns: []schema.Column{
			{
				Name:     "project_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveProject,
			},
			{
				Name:     "access",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Access"),
			},
			{
				Name:     "creation_time",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("CreationTime"),
			},
			{
				Name:     "dataset_reference",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("DatasetReference"),
			},
			{
				Name:     "default_collation",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DefaultCollation"),
			},
			{
				Name:     "default_encryption_configuration",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("DefaultEncryptionConfiguration"),
			},
			{
				Name:     "default_partition_expiration_ms",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("DefaultPartitionExpirationMs"),
			},
			{
				Name:     "default_table_expiration_ms",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("DefaultTableExpirationMs"),
			},
			{
				Name:     "description",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Description"),
			},
			{
				Name:     "etag",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Etag"),
			},
			{
				Name:     "friendly_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("FriendlyName"),
			},
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Id"),
			},
			{
				Name:     "is_case_insensitive",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("IsCaseInsensitive"),
			},
			{
				Name:     "kind",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Kind"),
			},
			{
				Name:     "labels",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Labels"),
			},
			{
				Name:     "last_modified_time",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("LastModifiedTime"),
			},
			{
				Name:     "location",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Location"),
			},
			{
				Name:     "max_time_travel_hours",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("MaxTimeTravelHours"),
			},
			{
				Name:     "satisfies_pzs",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("SatisfiesPzs"),
			},
			{
				Name:     "self_link",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SelfLink"),
			},
			{
				Name:     "storage_billing_model",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("StorageBillingModel"),
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Tags"),
			},
		},

		Relations: []*schema.Table{
			Tables(),
		},
	}
}
