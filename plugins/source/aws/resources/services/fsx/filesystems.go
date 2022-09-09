package fsx

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/fsx"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Filesystems() *schema.Table {
	return &schema.Table{
		Name:        "aws_fsx_filesystems",
		Description: "A description of a specific Amazon FSx file system",
		Resolver:    fetchFsxFilesystems,
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
				Description: "The time that the file system was created, in seconds (since 1970-01-01T00:00:00Z), also known as Unix time",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "dns_name",
				Description: "The Domain Name System (DNS) name for the file system",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DNSName"),
			},
			{
				Name:     "failure_details",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("FailureDetails"),
			},
			{
				Name:        "id",
				Description: "The system-generated, unique 17-digit ID of the file system",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("FileSystemId"),
			},
			{
				Name:        "type",
				Description: "The type of Amazon FSx file system, which can be LUSTRE, WINDOWS, ONTAP, or OPENZFS",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("FileSystemType"),
			},
			{
				Name:        "version",
				Description: "The Lustre version of the Amazon FSx for Lustre file system, either 2.10 or 2.12",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("FileSystemTypeVersion"),
			},
			{
				Name:        "kms_key_id",
				Description: "The ID of the Key Management Service (KMS) key used to encrypt Amazon FSx file system data",
				Type:        schema.TypeString,
			},
			{
				Name:        "lifecycle",
				Description: "The lifecycle status of the file system",
				Type:        schema.TypeString,
			},
			{
				Name:        "network_interface_ids",
				Description: "The IDs of the elastic network interfaces from which a specific file system is accessible",
				Type:        schema.TypeStringArray,
			},
			{
				Name:        "owner_id",
				Description: "The Amazon Web Services account that created the file system",
				Type:        schema.TypeString,
			},
			{
				Name:            "arn",
				Description:     "The Amazon Resource Name (ARN) of the file system resource",
				Type:            schema.TypeString,
				Resolver:        schema.PathResolver("ResourceARN"),
				CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
			},
			{
				Name:        "storage_capacity",
				Description: "The storage capacity of the file system in gibibytes (GiB)",
				Type:        schema.TypeInt,
			},
			{
				Name:        "storage_type",
				Description: "The type of storage the file system is using",
				Type:        schema.TypeString,
			},
			{
				Name:        "subnet_ids",
				Description: "Specifies the IDs of the subnets that the file system is accessible from",
				Type:        schema.TypeStringArray,
			},
			{
				Name:        "tags",
				Description: "The tags to associate with the file system",
				Type:        schema.TypeJSON,
				Resolver:    client.ResolveTags,
			},
			{
				Name:        "vpc_id",
				Description: "The ID of the primary virtual private cloud (VPC) for the file system",
				Type:        schema.TypeString,
			},
			{
				Name:        "lustre_configuration",
				Description: "The configuration for the Amazon FSx for Lustre file system",
				Type:        schema.TypeJSON,
				Resolver:    schema.PathResolver("LustreConfiguration"),
			},
			{
				Name:        "ontap_configuration",
				Description: "Configuration for the FSx for NetApp ONTAP file system",
				Type:        schema.TypeJSON,
				Resolver:    schema.PathResolver("OntapConfiguration"),
			},
			{
				Name:        "open_zfs_configuration",
				Description: "The configuration for the Amazon FSx for OpenZFS file system",
				Type:        schema.TypeJSON,
				Resolver:    schema.PathResolver("OpenZFSConfiguration"),
			},
			{
				Name:        "windows_configuration",
				Description: "The configuration for this Microsoft Windows file system",
				Type:        schema.TypeJSON,
				Resolver:    schema.PathResolver("WindowsConfiguration"),
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchFsxFilesystems(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	cl := meta.(*client.Client)
	svc := cl.Services().FSX
	input := fsx.DescribeFileSystemsInput{MaxResults: aws.Int32(1000)}
	paginator := fsx.NewDescribeFileSystemsPaginator(svc, &input)
	for paginator.HasMorePages() {
		result, err := paginator.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- result.FileSystems
	}
	return nil
}
