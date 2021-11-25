package resources

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker"
	sagemakertypes "github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func SagemakerNotebookInstances() *schema.Table {
	return &schema.Table{
		Name:         "aws_sagemaker_notebook_instances",
		Description:  "Provides summary information for an Amazon SageMaker notebook instance.",
		Resolver:     fetchSagemakerNotebookInstances,
		Multiplex:    client.AccountRegionMultiplex,
		IgnoreError:  client.IgnoreAccessDeniedServiceDisabled,
		DeleteFilter: client.DeleteAccountRegionFilter,
		Options:      schema.TableCreationOptions{PrimaryKeys: []string{"arn"}},
		Columns: []schema.Column{
			{
				Name:        "account_id",
				Description: "The AWS Account ID of the resource.",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSAccount,
			},
			{
				Name:        "region",
				Description: "The AWS Region of the resource.",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSRegion,
			},
			{
				Name:        "network_interface_id",
				Description: "The network interface IDs that Amazon SageMaker created at the time of creating the instance.",
				Type:        schema.TypeString,
			},
			{
				Name:        "kms_key_id",
				Description: "The Amazon Web Services KMS key ID Amazon SageMaker uses to encrypt data when storing it on the ML storage volume attached to the instance.",
				Type:        schema.TypeString,
			},
			{
				Name:        "subnet_id",
				Description: "The ID of the VPC subnet.",
				Type:        schema.TypeString,
			},
			{
				Name:        "volume_size_in_gb",
				Description: "The size, in GB, of the ML storage volume attached to the notebook instance.",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("VolumeSizeInGB"),
			},
			{
				Name:        "accelerator_types",
				Description: "A list of the Elastic Inference (EI) instance types associated with this notebook instance.",
				Type:        schema.TypeStringArray,
			},
			{
				Name:        "security_groups",
				Description: "The IDs of the VPC security groups.",
				Type:        schema.TypeJSON,
			},
			{
				Name:        "direct_internet_access",
				Description: "Describes whether Amazon SageMaker provides internet access to the notebook instance.",
				Type:        schema.TypeBool,
				Resolver:    resolveSagemakerNotebookInstanceDirectInternetAccess,
			},
			{
				Name:        "tags",
				Description: "The tags associated with the notebook instance.",
				Type:        schema.TypeJSON,
				Resolver:    resolveSagemakerNotebookInstanceTags,
			},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) of the notebook instance.  This member is required.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("NotebookInstanceArn"),
			},
			{
				Name:        "name",
				Description: "The name of the notebook instance that you want a summary for.  This member is required.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("NotebookInstanceName"),
			},
			{
				Name:        "additional_code_repositories",
				Description: "An array of up to three Git repositories associated with the notebook instance. These can be either the names of Git repositories stored as resources in your account, or the URL of Git repositories in Amazon Web Services CodeCommit (https://docs.aws.amazon.com/codecommit/latest/userguide/welcome.html) or in any other Git repository",
				Type:        schema.TypeStringArray,
			},
			{
				Name:        "creation_time",
				Description: "A timestamp that shows when the notebook instance was created.",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "default_code_repository",
				Description: "The Git repository associated with the notebook instance as its default code repository",
				Type:        schema.TypeString,
			},
			{
				Name:        "instance_type",
				Description: "The type of ML compute instance that the notebook instance is running on.",
				Type:        schema.TypeString,
			},
			{
				Name:        "last_modified_time",
				Description: "A timestamp that shows when the notebook instance was last modified.",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "notebook_instance_lifecycle_config_name",
				Description: "The name of a notebook instance lifecycle configuration associated with this notebook instance",
				Type:        schema.TypeString,
			},
			{
				Name:        "notebook_instance_status",
				Description: "The status of the notebook instance.",
				Type:        schema.TypeString,
			},
			{
				Name:        "url",
				Description: "The URL that you use to connect to the Jupyter instance running in your notebook instance.",
				Type:        schema.TypeString,
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchSagemakerNotebookInstances(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan interface{}) error {
	c := meta.(*client.Client)
	svc := c.Services().SageMaker
	config := sagemaker.ListNotebookInstancesInput{}
	for {
		response, err := svc.ListNotebookInstances(ctx, &config, func(options *sagemaker.Options) {
			options.Region = c.Region
		})
		if err != nil {
			return err
		}

		// get more details about the notebook instance
		for _, n := range response.NotebookInstances {

			config := sagemaker.DescribeNotebookInstanceInput{
				NotebookInstanceName: n.NotebookInstanceName,
			}
			response, err := svc.DescribeNotebookInstance(ctx, &config, func(options *sagemaker.Options) {
				options.Region = c.Region
			})
			if err != nil {
				return err
			}

			notebook := WrappedSageMakerNotebookInstance{
				DescribeNotebookInstanceOutput: response,
				NotebookInstanceArn:            *n.NotebookInstanceArn,
				NotebookInstanceName:           *n.NotebookInstanceName,
			}

			res <- notebook
		}

		if aws.ToString(response.NextToken) == "" {
			break
		}
		config.NextToken = response.NextToken
	}
	return nil
}

func resolveSagemakerNotebookInstanceTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, _ schema.Column) error {
	r, ok := resource.Item.(WrappedSageMakerNotebookInstance)

	if !ok {
		return fmt.Errorf("expected WrappedNotebookInstance but got %T", r)
	}

	c := meta.(*client.Client)
	svc := c.Services().SageMaker
	config := sagemaker.ListTagsInput{
		ResourceArn: &r.NotebookInstanceArn,
	}
	response, err := svc.ListTags(ctx, &config, func(options *sagemaker.Options) {
		options.Region = c.Region
	})
	if err != nil {
		return err
	}

	tags := make(map[string]*string, len(response.Tags))
	for _, t := range response.Tags {
		tags[*t.Key] = t.Value
	}

	return resource.Set("tags", tags)
}

func resolveSagemakerNotebookInstanceDirectInternetAccess(_ context.Context, _ schema.ClientMeta, resource *schema.Resource, _ schema.Column) error {
	r, ok := resource.Item.(WrappedSageMakerNotebookInstance)

	if !ok {
		return fmt.Errorf("expected WrappedNotebookInstance but got %T", r)
	}

	if r.DirectInternetAccess == sagemakertypes.DirectInternetAccessEnabled {
		return resource.Set("direct_internet_access", true)
	}

	return resource.Set("direct_internet_access", false)

}

type WrappedSageMakerNotebookInstance struct {
	*sagemaker.DescribeNotebookInstanceOutput
	NotebookInstanceArn  string
	NotebookInstanceName string
}
