package fsx

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/fsx"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func DataRepoAssociations() *schema.Table {
	return &schema.Table{
		Name:        "aws_fsx_data_repo_associations",
		Description: "The configuration of a data repository association that links an Amazon FSx for Lustre file system to an Amazon S3 bucket",
		Resolver:    fetchFsxDataRepoAssociations,
		Multiplex:   client.ServiceAccountRegionMultiplexer("fsx"),
		Columns: []schema.Column{
			{
				Name:        "account_id",
				Description: "The AWS Account ID of the resource",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSAccount,
			},
			{
				Name:        "region",
				Description: "The AWS Region of the resource",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSRegion,
			},
			{
				Name:        "association_id",
				Description: "The system-generated, unique ID of the data repository association",
				Type:        schema.TypeString,
			},
			{
				Name:        "batch_import_meta_data_on_create",
				Description: "A boolean flag indicating whether an import data repository task to import metadata should run after the data repository association is created",
				Type:        schema.TypeBool,
			},
			{
				Name:        "creation_time",
				Description: "The time that the resource was created, in seconds (since 1970-01-01T00:00:00Z), also known as Unix time",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "data_repository_path",
				Description: "The path to the Amazon S3 data repository that will be linked to the file system",
				Type:        schema.TypeString,
			},
			{
				Name:     "failure_details",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("FailureDetails"),
			},
			{
				Name:        "file_system_id",
				Description: "The globally unique ID of the file system, assigned by Amazon FSx",
				Type:        schema.TypeString,
			},
			{
				Name:        "file_system_path",
				Description: "A path on the file system that points to a high-level directory (such as /ns1/) or subdirectory (such as /ns1/subdir/) that will be mapped 1-1 with DataRepositoryPath",
				Type:        schema.TypeString,
			},
			{
				Name:        "imported_file_chunk_size",
				Description: "For files imported from a data repository, this value determines the stripe count and maximum amount of data per file (in MiB) stored on a single physical disk",
				Type:        schema.TypeInt,
			},
			{
				Name:        "lifecycle",
				Description: "Describes the state of a data repository association",
				Type:        schema.TypeString,
			},
			{
				Name:            "arn",
				Description:     "The Amazon Resource Name (ARN) for a given resource",
				Type:            schema.TypeString,
				Resolver:        schema.PathResolver("ResourceARN"),
				CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
			},
			{
				Name:        "s3",
				Type: 			schema.TypeJSON,
				Resolver:    schema.PathResolver("S3"),
			},
			{
				Name:        "tags",
				Description: "A list of Tag values, with a maximum of 50 elements",
				Type:        schema.TypeJSON,
				Resolver:    client.ResolveTags,
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchFsxDataRepoAssociations(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	cl := meta.(*client.Client)
	svc := cl.Services().FSX
	input := fsx.DescribeDataRepositoryAssociationsInput{MaxResults: aws.Int32(25)}
	paginator := fsx.NewDescribeDataRepositoryAssociationsPaginator(svc, &input)
	for paginator.HasMorePages() {
		result, err := paginator.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- result.Associations
	}
	return nil
}
