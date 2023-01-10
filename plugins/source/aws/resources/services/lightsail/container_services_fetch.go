package lightsail

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/lightsail"
	"github.com/aws/aws-sdk-go-v2/service/lightsail/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchLightsailContainerServices(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
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
func fetchLightsailContainerServiceDeployments(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
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
func fetchLightsailContainerServiceImages(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
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
