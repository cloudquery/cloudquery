// Code generated by codegen; DO NOT EDIT.

package fsx

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func DataRepositoryAssociations() *schema.Table {
	return &schema.Table{
		Name:      "aws_fsx_data_repository_associations",
		Resolver:  fetchFsxDataRepositoryAssociations,
		Multiplex: client.ServiceAccountRegionMultiplexer("fsx"),
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSAccount,
			},
			{
				Name:     "region",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSRegion,
			},
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ResourceARN"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: client.ResolveTags,
			},
			{
				Name:     "association_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AssociationId"),
			},
			{
				Name:     "batch_import_meta_data_on_create",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("BatchImportMetaDataOnCreate"),
			},
			{
				Name:     "creation_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("CreationTime"),
			},
			{
				Name:     "data_repository_path",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DataRepositoryPath"),
			},
			{
				Name:     "failure_details",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("FailureDetails"),
			},
			{
				Name:     "file_system_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("FileSystemId"),
			},
			{
				Name:     "file_system_path",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("FileSystemPath"),
			},
			{
				Name:     "imported_file_chunk_size",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("ImportedFileChunkSize"),
			},
			{
				Name:     "lifecycle",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Lifecycle"),
			},
			{
				Name:     "s_3",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("S3"),
			},
		},
	}
}
