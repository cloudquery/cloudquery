package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func vpcEndpointServicePermissions() *schema.Table {
	return &schema.Table{
		Name:        "aws_ec2_vpc_endpoint_service_permissions",
		Description: `https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_AllowedPrincipal.html`,
		Resolver:    fetchEc2VpcEndpointServicePermissions,
		Multiplex:   client.ServiceAccountRegionMultiplexer("ec2"),
		Transform:   transformers.TransformWithStruct(&types.AllowedPrincipal{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: client.ResolveTags,
			},
		},
	}
}
func fetchEc2VpcEndpointServicePermissions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	endpointService := parent.Item.(types.ServiceDetail)
	if aws.ToString(endpointService.Owner) == "amazon" {
		return nil
	}
	svc := meta.(*client.Client).Services().Ec2
	paginator := ec2.NewDescribeVpcEndpointServicePermissionsPaginator(svc, &ec2.DescribeVpcEndpointServicePermissionsInput{
		ServiceId: endpointService.ServiceId,
	})
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- page.AllowedPrincipals
	}
	return nil
}
