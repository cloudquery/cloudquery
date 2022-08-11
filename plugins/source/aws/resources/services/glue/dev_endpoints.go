package glue

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/glue"
	"github.com/aws/aws-sdk-go-v2/service/glue/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

//go:generate cq-gen --resource dev_endpoints --config dev_endpoints.hcl --output .
func DevEndpoints() *schema.Table {
	return &schema.Table{
		Name:         "aws_glue_dev_endpoints",
		Description:  "A development endpoint where a developer can remotely debug extract, transform, and load (ETL) scripts",
		Resolver:     fetchGlueDevEndpoints,
		Multiplex:    client.ServiceAccountRegionMultiplexer("glue"),
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
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) of the workflow.",
				Type:        schema.TypeString,
				Resolver:    resolveGlueDevEndpointArn,
			},
			{
				Name:        "tags",
				Description: "Resource tags.",
				Type:        schema.TypeJSON,
				Resolver:    resolveGlueDevEndpointTags,
			},
			{
				Name:        "arguments",
				Description: "A map of arguments used to configure the DevEndpoint",
				Type:        schema.TypeJSON,
			},
			{
				Name:        "availability_zone",
				Description: "The AWS Availability Zone where this DevEndpoint is located",
				Type:        schema.TypeString,
			},
			{
				Name:        "created_timestamp",
				Description: "The point in time at which this DevEndpoint was created",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "name",
				Description: "The name of the DevEndpoint",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("EndpointName"),
			},
			{
				Name:        "extra_jars_s3_path",
				Description: "The path to one or more Java jar files in an S3 bucket that should be loaded in your DevEndpoint",
				Type:        schema.TypeString,
			},
			{
				Name:        "extra_python_libs_s3_path",
				Description: "The paths to one or more Python libraries in an Amazon S3 bucket that should be loaded in your DevEndpoint",
				Type:        schema.TypeString,
			},
			{
				Name:        "failure_reason",
				Description: "The reason for a current failure in this DevEndpoint",
				Type:        schema.TypeString,
			},
			{
				Name:        "glue_version",
				Description: "Glue version determines the versions of Apache Spark and Python that Glue supports",
				Type:        schema.TypeString,
			},
			{
				Name:        "last_modified_timestamp",
				Description: "The point in time at which this DevEndpoint was last modified",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "last_update_status",
				Description: "The status of the last update",
				Type:        schema.TypeString,
			},
			{
				Name:        "number_of_nodes",
				Description: "The number of Glue Data Processing Units (DPUs) allocated to this DevEndpoint",
				Type:        schema.TypeBigInt,
			},
			{
				Name:        "number_of_workers",
				Description: "The number of workers of a defined workerType that are allocated to the development endpoint",
				Type:        schema.TypeBigInt,
			},
			{
				Name:        "private_address",
				Description: "A private IP address to access the DevEndpoint within a VPC if the DevEndpoint is created within one",
				Type:        schema.TypeString,
			},
			{
				Name:        "public_address",
				Description: "The public IP address used by this DevEndpoint",
				Type:        schema.TypeString,
			},
			{
				Name:        "public_key",
				Description: "The public key to be used by this DevEndpoint for authentication",
				Type:        schema.TypeString,
			},
			{
				Name:        "public_keys",
				Description: "A list of public keys to be used by the DevEndpoints for authentication",
				Type:        schema.TypeStringArray,
			},
			{
				Name:        "role_arn",
				Description: "The Amazon Resource Name (ARN) of the IAM role used in this DevEndpoint",
				Type:        schema.TypeString,
			},
			{
				Name:        "security_configuration",
				Description: "The name of the SecurityConfiguration structure to be used with this DevEndpoint",
				Type:        schema.TypeString,
			},
			{
				Name:        "security_group_ids",
				Description: "A list of security group identifiers used in this DevEndpoint",
				Type:        schema.TypeStringArray,
			},
			{
				Name:        "status",
				Description: "The current status of this DevEndpoint",
				Type:        schema.TypeString,
			},
			{
				Name:        "subnet_id",
				Description: "The subnet ID for this DevEndpoint",
				Type:        schema.TypeString,
			},
			{
				Name:        "vpc_id",
				Description: "The ID of the virtual private cloud (VPC) used by this DevEndpoint",
				Type:        schema.TypeString,
			},
			{
				Name:        "worker_type",
				Description: "The type of predefined worker that is allocated to the development endpoint Accepts a value of Standard, G1X, or G2X",
				Type:        schema.TypeString,
			},
			{
				Name:        "yarn_endpoint_address",
				Description: "The YARN endpoint address used by this DevEndpoint",
				Type:        schema.TypeString,
			},
			{
				Name:        "zeppelin_remote_spark_interpreter_port",
				Description: "The Apache Zeppelin port for the remote Apache Spark interpreter",
				Type:        schema.TypeBigInt,
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchGlueDevEndpoints(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Glue
	input := glue.GetDevEndpointsInput{}
	for {
		result, err := svc.GetDevEndpoints(ctx, &input)
		if err != nil {
			if cl.IsNotFoundError(err) {
				return nil
			}
			return diag.WrapError(err)
		}
		res <- result.DevEndpoints
		if aws.ToString(result.NextToken) == "" {
			break
		}
		input.NextToken = result.NextToken
	}
	return nil
}
func resolveGlueDevEndpointArn(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	arn := aws.String(devEndpointARN(cl, aws.ToString(resource.Item.(types.DevEndpoint).EndpointName)))
	return diag.WrapError(resource.Set(c.Name, arn))
}
func resolveGlueDevEndpointTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Glue
	result, err := svc.GetTags(ctx, &glue.GetTagsInput{
		ResourceArn: aws.String(devEndpointARN(cl, aws.ToString(resource.Item.(types.DevEndpoint).EndpointName))),
	})
	if err != nil {
		if cl.IsNotFoundError(err) {
			return nil
		}
		return diag.WrapError(err)
	}
	return diag.WrapError(resource.Set(c.Name, result.Tags))
}

// ====================================================================================================================
//                                                  User Defined Helpers
// ====================================================================================================================

func devEndpointARN(cl *client.Client, name string) string {
	return cl.ARN(client.GlueService, "devEndpoint", name)
}
