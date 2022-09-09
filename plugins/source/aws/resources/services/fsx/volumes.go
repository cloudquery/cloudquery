package fsx

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/fsx"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Volumes() *schema.Table {
	return &schema.Table{
		Name:        "aws_fsx_volumes",
		Description: "Describes an Amazon FSx for NetApp ONTAP or Amazon FSx for OpenZFS volume",
		Resolver:    fetchFsxVolumes,
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
				Name:        "creation_time",
				Description: "The time that the resource was created, in seconds (since 1970-01-01T00:00:00Z), also known as Unix time",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "file_system_id",
				Description: "The globally unique ID of the file system, assigned by Amazon FSx",
				Type:        schema.TypeString,
			},
			{
				Name:        "lifecycle",
				Description: "The lifecycle status of the volume",
				Type:        schema.TypeString,
			},
			{
				Name:        "lifecycle_transition_reason_message",
				Description: "A detailed error message",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("LifecycleTransitionReason.Message"),
			},
			{
				Name:        "name",
				Description: "The name of the volume",
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
				Name:        "tags",
				Description: "A list of Tag values, with a maximum of 50 elements",
				Type:        schema.TypeJSON,
				Resolver:    client.ResolveTags,
			},
			{
				Name:        "id",
				Description: "The system-generated, unique ID of the volume",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("VolumeId"),
			},
			{
				Name:        "volume_type",
				Description: "The type of the volume",
				Type:        schema.TypeString,
			},
			{
				Name:        "ontap_configuration",
				Description: "The configuration of an Amazon FSx for NetApp ONTAP volume",
				Type:        schema.TypeJSON,
				Resolver:    schema.PathResolver("OntapConfiguration"),
			},
			{
				Name:        "open_zfs_configuration",
				Description: "The configuration of an Amazon FSx for OpenZFS volume",
				Type:        schema.TypeJSON,
				Resolver:    schema.PathResolver("OpenZFSConfiguration"),
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchFsxVolumes(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	cl := meta.(*client.Client)
	svc := cl.Services().FSX
	input := fsx.DescribeVolumesInput{MaxResults: aws.Int32(1000)}
	paginator := fsx.NewDescribeVolumesPaginator(svc, &input)
	for paginator.HasMorePages() {
		result, err := paginator.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- result.Volumes
	}
	return nil
}
