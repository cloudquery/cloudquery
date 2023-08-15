package servicediscovery

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/servicediscovery"
	"github.com/aws/aws-sdk-go-v2/service/servicediscovery/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func instances() *schema.Table {
	return &schema.Table{
		Name:                "aws_servicediscovery_instances",
		Description:         `https://docs.aws.amazon.com/cloud-map/latest/api/API_Instance.html`,
		Resolver:            fetchInstances,
		PreResourceResolver: getInstance,
		Transform:           transformers.TransformWithStruct(&types.Instance{}, transformers.WithPrimaryKeys("Id")),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
			client.DefaultRegionColumn(true),
		},
	}
}

func fetchInstances(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Servicediscovery
	service := parent.Item.(*types.Service)
	config := servicediscovery.ListInstancesInput{
		ServiceId:  service.Id,
		MaxResults: aws.Int32(100),
	}
	paginator := servicediscovery.NewListInstancesPaginator(svc, &config)
	for paginator.HasMorePages() {
		output, err := paginator.NextPage(ctx, func(o *servicediscovery.Options) {
			o.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- output.Instances
	}
	return nil
}

func getInstance(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Servicediscovery
	instance := resource.Item.(types.InstanceSummary)
	service := resource.Parent.Item.(*types.Service)
	config := &servicediscovery.GetInstanceInput{
		InstanceId: instance.Id,
		ServiceId:  service.Id,
	}
	desc, err := svc.GetInstance(ctx, config, func(o *servicediscovery.Options) {
		o.Region = cl.Region
	})
	if err != nil {
		return err
	}
	resource.Item = desc.Instance
	return nil
}
