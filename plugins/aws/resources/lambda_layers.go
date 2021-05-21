package resources

import (
	"context"
	"errors"
	"fmt"

	"github.com/aws/smithy-go"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/lambda"
	"github.com/aws/aws-sdk-go-v2/service/lambda/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func LambdaLayers() *schema.Table {
	return &schema.Table{
		Name:         "aws_lambda_layers",
		Resolver:     fetchLambdaLayers,
		Multiplex:    client.AccountRegionMultiplex,
		IgnoreError:  client.IgnoreAccessDeniedServiceDisabled,
		DeleteFilter: client.DeleteAccountRegionFilter,
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
				Name:     "latest_matching_version_compatible_runtimes",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("LatestMatchingVersion.CompatibleRuntimes"),
			},
			{
				Name:     "latest_matching_version_created_date",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("LatestMatchingVersion.CreatedDate"),
			},
			{
				Name:     "latest_matching_version_description",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("LatestMatchingVersion.Description"),
			},
			{
				Name:     "latest_matching_version_layer_version_arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("LatestMatchingVersion.LayerVersionArn"),
			},
			{
				Name:     "latest_matching_version_license_info",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("LatestMatchingVersion.LicenseInfo"),
			},
			{
				Name:     "latest_matching_version",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("LatestMatchingVersion.Version"),
			},
			{
				Name: "layer_arn",
				Type: schema.TypeString,
			},
			{
				Name: "layer_name",
				Type: schema.TypeString,
			},
		},
		Relations: []*schema.Table{
			{
				Name:     "aws_lambda_layer_versions",
				Resolver: fetchLambdaLayerVersions,
				Columns: []schema.Column{
					{
						Name:     "layer_id",
						Type:     schema.TypeUUID,
						Resolver: schema.ParentIdResolver,
					},
					{
						Name: "compatible_runtimes",
						Type: schema.TypeStringArray,
					},
					{
						Name: "created_date",
						Type: schema.TypeString,
					},
					{
						Name: "description",
						Type: schema.TypeString,
					},
					{
						Name: "layer_version_arn",
						Type: schema.TypeString,
					},
					{
						Name: "license_info",
						Type: schema.TypeString,
					},
					{
						Name: "version",
						Type: schema.TypeBigInt,
					},
				},
				Relations: []*schema.Table{
					{
						Name:     "aws_lambda_layer_version_policies",
						Resolver: fetchLambdaLayerVersionPolicies,
						Columns: []schema.Column{
							{
								Name:     "layer_version_id",
								Type:     schema.TypeUUID,
								Resolver: schema.ParentIdResolver,
							},
							{
								Name: "policy",
								Type: schema.TypeString,
							},
							{
								Name: "revision_id",
								Type: schema.TypeString,
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
func fetchLambdaLayers(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	var input lambda.ListLayersInput
	c := meta.(*client.Client)
	svc := c.Services().Lambda
	for {
		response, err := svc.ListLayers(ctx, &input, func(options *lambda.Options) {
			options.Region = c.Region
		})
		if err != nil {
			return err
		}

		res <- response.Layers

		if aws.ToString(response.NextMarker) == "" {
			break
		}
		input.Marker = response.NextMarker
	}
	return nil
}
func fetchLambdaLayerVersions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
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
			return err
		}
		res <- output.LayerVersions
		if output.NextMarker == nil {
			break
		}
		config.Marker = output.NextMarker
	}
	return nil
}
func fetchLambdaLayerVersionPolicies(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p, ok := parent.Item.(types.LayerVersionsListItem)
	if !ok {
		return fmt.Errorf("wrong type assertion: got %T instead of LayerVersionsListItem", p)
	}

	pp, ok := parent.Parent.Item.(types.LayersListItem)
	if !ok {
		return fmt.Errorf("wrong type assertion: got %T instead of LayersListItem", p)
	}
	svc := meta.(*client.Client).Services().Lambda

	config := lambda.GetLayerVersionPolicyInput{
		LayerName:     pp.LayerName,
		VersionNumber: p.Version,
	}

	output, err := svc.GetLayerVersionPolicy(ctx, &config)
	var ae smithy.APIError
	if err != nil {
		if errors.As(err, &ae) && ae.ErrorCode() == "ResourceNotFoundException" {
			return nil
		}
		return err
	}
	res <- output

	return nil
}
