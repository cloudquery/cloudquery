package lightsail

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/lightsail"
	"github.com/aws/aws-sdk-go-v2/service/lightsail/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func ContainerServices() *schema.Table {
	return &schema.Table{
		Name:        "aws_lightsail_container_services",
		Description: "Describes an Amazon Lightsail container service",
		Resolver:    fetchLightsailContainerServices,
		Multiplex:   client.ServiceAccountRegionMultiplexer("lightsail"),
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
				Name:            "arn",
				Description:     "The Amazon Resource Name (ARN) of the container service",
				Type:            schema.TypeString,
				CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
			},
			{
				Name:        "container_service_name",
				Description: "The name of the container service",
				Type:        schema.TypeString,
			},
			{
				Name:        "created_at",
				Description: "The timestamp when the container service was created",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:     "current_deployment",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("CurrentDeployment"),
			},
			{
				Name:        "is_disabled",
				Description: "A Boolean value indicating whether the container service is disabled",
				Type:        schema.TypeBool,
			},
			{
				Name:     "location",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Location"),
			},
			{
				Name:          "next_deployment",
				Type:          schema.TypeJSON,
				Resolver:      schema.PathResolver("NextDeployment"),
				IgnoreInTests: true,
			},
			{
				Name:        "power",
				Description: "The power specification of the container service",
				Type:        schema.TypeString,
			},
			{
				Name:        "power_id",
				Description: "The ID of the power of the container service",
				Type:        schema.TypeString,
			},
			{
				Name:        "principal_arn",
				Description: "The principal ARN of the container service",
				Type:        schema.TypeString,
			},
			{
				Name:        "private_domain_name",
				Description: "The private domain name of the container service",
				Type:        schema.TypeString,
			},
			{
				Name:     "private_registry_access",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("PrivateRegistryAccess"),
			},
			{
				Name:          "public_domain_names",
				Description:   "The public domain name of the container service, such as examplecom and wwwexamplecom",
				Type:          schema.TypeJSON,
				IgnoreInTests: true,
			},
			{
				Name:        "resource_type",
				Description: "The Lightsail resource type of the container service (ie, ContainerService)",
				Type:        schema.TypeString,
			},
			{
				Name:        "scale",
				Description: "The scale specification of the container service",
				Type:        schema.TypeInt,
			},
			{
				Name:        "state",
				Description: "The current state of the container service",
				Type:        schema.TypeString,
			},
			{
				Name:     "state_detail",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("StateDetail"),
			},
			{
				Name:        "tags",
				Description: "The tag keys and optional values for the resource",
				Type:        schema.TypeJSON,
				Resolver:    client.ResolveTags,
			},
			{
				Name:        "url",
				Description: "The publicly accessible URL of the container service",
				Type:        schema.TypeString,
			},
		},
		Relations: []*schema.Table{
			{
				Name:        "aws_lightsail_container_service_deployments",
				Description: "Describes a container deployment configuration of an Amazon Lightsail container service",
				Resolver:    fetchLightsailContainerServiceDeployments,
				Columns: []schema.Column{
					{
						Name:        "container_service_cq_id",
						Description: "Unique CloudQuery ID of aws_lightsail_container_services table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "containers",
						Description: "An object that describes the configuration for the containers of the deployment",
						Type:        schema.TypeJSON,
					},
					{
						Name:        "created_at",
						Description: "The timestamp when the deployment was created",
						Type:        schema.TypeTimestamp,
					},
					{
						Name:     "public_endpoint",
						Type:     schema.TypeJSON,
						Resolver: schema.PathResolver("PublicEndpoint"),
					},
					{
						Name:        "state",
						Description: "The state of the deployment",
						Type:        schema.TypeString,
					},
					{
						Name:        "version",
						Description: "The version number of the deployment",
						Type:        schema.TypeInt,
					},
				},
			},
			{
				Name:        "aws_lightsail_container_service_images",
				Description: "Describes a container image that is registered to an Amazon Lightsail container service",
				Resolver:    fetchLightsailContainerServiceImages,
				Columns: []schema.Column{
					{
						Name:        "container_service_cq_id",
						Description: "Unique CloudQuery ID of aws_lightsail_container_services table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "created_at",
						Description: "The timestamp when the container image was created",
						Type:        schema.TypeTimestamp,
					},
					{
						Name:        "digest",
						Description: "The digest of the container image",
						Type:        schema.TypeString,
					},
					{
						Name:        "image",
						Description: "The name of the container image",
						Type:        schema.TypeString,
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchLightsailContainerServices(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	var input lightsail.GetContainerServicesInput
	c := meta.(*client.Client)
	svc := c.Services().Lightsail
	response, err := svc.GetContainerServices(ctx, &input)
	if err != nil {
		return err
	}
	res <- response.ContainerServices
	return nil
}
func fetchLightsailContainerServiceDeployments(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	r := parent.Item.(types.ContainerService)
	input := lightsail.GetContainerServiceDeploymentsInput{
		ServiceName: r.ContainerServiceName,
	}
	c := meta.(*client.Client)
	svc := c.Services().Lightsail
	deployments, err := svc.GetContainerServiceDeployments(ctx, &input)
	if err != nil {
		return err
	}
	res <- deployments.Deployments
	return nil
}
func fetchLightsailContainerServiceImages(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	r := parent.Item.(types.ContainerService)
	input := lightsail.GetContainerImagesInput{
		ServiceName: r.ContainerServiceName,
	}
	c := meta.(*client.Client)
	svc := c.Services().Lightsail
	deployments, err := svc.GetContainerImages(ctx, &input)
	if err != nil {
		return err
	}
	res <- deployments.ContainerImages
	return nil
}
