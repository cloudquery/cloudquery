package fsx

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/fsx"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func StorageVms() *schema.Table {
	return &schema.Table{
		Name:        "aws_fsx_storage_vms",
		Description: "Describes the Amazon FSx for NetApp ONTAP storage virtual machine (SVM) configuration",
		Resolver:    fetchFsxStorageVms,
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
				Name:     "active_directory_configuration",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ActiveDirectoryConfiguration"),
			},
			{
				Name:        "creation_time",
				Description: "The time that the resource was created, in seconds (since 1970-01-01T00:00:00Z), also known as Unix time",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "endpoints",
				Type: 			schema.TypeJSON,
			},
			{
				Name:        "file_system_id",
				Description: "The globally unique ID of the file system, assigned by Amazon FSx",
				Type:        schema.TypeString,
			},
			{
				Name:        "lifecycle",
				Description: "Describes the SVM's lifecycle status",
				Type:        schema.TypeString,
			},
			{
				Name:     "lifecycle_transition_reason",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("LifecycleTransitionReason"),
			},
			{
				Name:        "name",
				Description: "The name of the SVM, if provisioned",
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
				Name:        "root_volume_security_style",
				Description: "The security style of the root volume of the SVM",
				Type:        schema.TypeString,
			},
			{
				Name:        "id",
				Description: "The SVM's system generated unique ID",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("StorageVirtualMachineId"),
			},
			{
				Name:        "subtype",
				Description: "Describes the SVM's subtype",
				Type:        schema.TypeString,
			},
			{
				Name:        "tags",
				Description: "A list of Tag values, with a maximum of 50 elements",
				Type:        schema.TypeJSON,
				Resolver:    client.ResolveTags,
			},
			{
				Name:        "uuid",
				Description: "The SVM's UUID (universally unique identifier)",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("UUID"),
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchFsxStorageVms(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	cl := meta.(*client.Client)
	svc := cl.Services().FSX
	input := fsx.DescribeStorageVirtualMachinesInput{MaxResults: aws.Int32(1000)}
	paginator := fsx.NewDescribeStorageVirtualMachinesPaginator(svc, &input)
	for paginator.HasMorePages() {
		result, err := paginator.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- result.StorageVirtualMachines
	}
	return nil
}
