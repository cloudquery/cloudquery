package lambda

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/lambda"
	"github.com/aws/aws-sdk-go-v2/service/lambda/types"
	"github.com/cloudquery/cq-provider-aws/client"

	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func LambdaLayers() *schema.Table {
	return &schema.Table{
		Name:         "aws_lambda_layers",
		Description:  "Details about an AWS Lambda layer (https://docs.aws.amazon.com/lambda/latest/dg/configuration-layers.html). ",
		Resolver:     fetchLambdaLayers,
		Multiplex:    client.ServiceAccountRegionMultiplexer("lambda"),
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
				Name:        "latest_matching_version_compatible_runtimes",
				Description: "The layer's compatible runtimes.",
				Type:        schema.TypeStringArray,
				Resolver:    schema.PathResolver("LatestMatchingVersion.CompatibleRuntimes"),
			},
			{
				Name:        "latest_matching_version_created_date",
				Description: "The date that the version was created, in ISO 8601 format",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("LatestMatchingVersion.CreatedDate"),
			},
			{
				Name:        "latest_matching_version_description",
				Description: "The description of the version.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("LatestMatchingVersion.Description"),
			},
			{
				Name:        "latest_matching_version_layer_version_arn",
				Description: "The ARN of the layer version.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("LatestMatchingVersion.LayerVersionArn"),
			},
			{
				Name:          "latest_matching_version_license_info",
				Description:   "The layer's open-source license.",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("LatestMatchingVersion.LicenseInfo"),
				IgnoreInTests: true,
			},
			{
				Name:        "latest_matching_version",
				Description: "The version number.",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("LatestMatchingVersion.Version"),
			},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) of the function layer.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("LayerArn"),
			},
			{
				Name:        "name",
				Description: "The name of the layer.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("LayerName"),
			},
		},
		Relations: []*schema.Table{
			{
				Name:        "aws_lambda_layer_versions",
				Description: "Details about a version of an AWS Lambda layer (https://docs.aws.amazon.com/lambda/latest/dg/configuration-layers.html). ",
				Resolver:    fetchLambdaLayerVersions,
				Options:     schema.TableCreationOptions{PrimaryKeys: []string{"layer_cq_id", "version"}},
				Columns: []schema.Column{
					{
						Name:        "layer_cq_id",
						Description: "Unique CloudQuery ID of aws_lambda_layers table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "compatible_runtimes",
						Description: "The layer's compatible runtimes.",
						Type:        schema.TypeStringArray,
					},
					{
						Name:        "created_date",
						Description: "The date that the version was created, in ISO 8601 format",
						Type:        schema.TypeString,
					},
					{
						Name:        "description",
						Description: "The description of the version.",
						Type:        schema.TypeString,
					},
					{
						Name:        "layer_version_arn",
						Description: "The ARN of the layer version.",
						Type:        schema.TypeString,
					},
					{
						Name:          "license_info",
						Description:   "The layer's open-source license.",
						Type:          schema.TypeString,
						IgnoreInTests: true,
					},
					{
						Name:        "version",
						Description: "The version number.",
						Type:        schema.TypeBigInt,
					},
				},
				Relations: []*schema.Table{
					{
						Name:          "aws_lambda_layer_version_policies",
						Resolver:      fetchLambdaLayerVersionPolicies,
						Options:       schema.TableCreationOptions{PrimaryKeys: []string{"layer_version_cq_id", "revision_id"}},
						IgnoreInTests: true,
						Columns: []schema.Column{
							{
								Name:        "layer_version_cq_id",
								Description: "Unique CloudQuery ID of aws_lambda_layer_versions table (FK)",
								Type:        schema.TypeUUID,
								Resolver:    schema.ParentIdResolver,
							},
							{
								Name:        "layer_version",
								Description: "The version number.",
								Type:        schema.TypeBigInt,
								Resolver:    schema.ParentResourceFieldResolver("version"),
							},
							{
								Name:        "policy",
								Description: "The policy document.",
								Type:        schema.TypeString,
							},
							{
								Name:        "revision_id",
								Description: "A unique identifier for the current revision of the policy.",
								Type:        schema.TypeString,
							},
						},
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchLambdaLayers(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	var input lambda.ListLayersInput
	c := meta.(*client.Client)
	svc := c.Services().Lambda
	for {
		response, err := svc.ListLayers(ctx, &input, func(options *lambda.Options) {
			options.Region = c.Region
		})
		if err != nil {
			return diag.WrapError(err)
		}

		res <- response.Layers

		if aws.ToString(response.NextMarker) == "" {
			break
		}
		input.Marker = response.NextMarker
	}
	return nil
}
func fetchLambdaLayerVersions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	p, ok := parent.Item.(types.LayersListItem)
	if !ok {
		return fmt.Errorf("wrong type assertion: got %T instead of LayersListItem", p)
	}
	svc := meta.(*client.Client).Services().Lambda
	config := lambda.ListLayerVersionsInput{
		LayerName: p.LayerName,
	}

	for {
		output, err := svc.ListLayerVersions(ctx, &config)
		if err != nil {
			return diag.WrapError(err)
		}
		res <- output.LayerVersions
		if output.NextMarker == nil {
			break
		}
		config.Marker = output.NextMarker
	}
	return nil
}
func fetchLambdaLayerVersionPolicies(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	p, ok := parent.Item.(types.LayerVersionsListItem)
	if !ok {
		return fmt.Errorf("wrong type assertion: got %T instead of LayerVersionsListItem", p)
	}

	pp, ok := parent.Parent.Item.(types.LayersListItem)
	if !ok {
		return fmt.Errorf("wrong type assertion: got %T instead of LayersListItem", p)
	}
	c := meta.(*client.Client)
	svc := c.Services().Lambda

	config := lambda.GetLayerVersionPolicyInput{
		LayerName:     pp.LayerName,
		VersionNumber: p.Version,
	}

	output, err := svc.GetLayerVersionPolicy(ctx, &config)
	if err != nil {
		if client.IsAWSError(err, "ResourceNotFoundException") {
			return nil
		}
		return err
	}
	res <- output

	return nil
}
