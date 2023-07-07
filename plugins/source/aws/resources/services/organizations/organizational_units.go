package organizations

import (
	"context"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/organizations"
	"github.com/aws/aws-sdk-go-v2/service/organizations/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/services"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func OrganizationalUnits() *schema.Table {
	tableName := "aws_organizations_organizational_units"
	return &schema.Table{
		Name: tableName,
		Description: `https://docs.aws.amazon.com/organizations/latest/APIReference/API_OrganizationalUnit.html
The 'request_account_id' column is added to show from where the request was made.`,
		Resolver:            fetchOUs,
		PreResourceResolver: getOU,
		Transform: transformers.TransformWithStruct(
			&types.OrganizationalUnit{},
			transformers.WithPrimaryKeys("Arn"),
		),
		Multiplex: client.ServiceAccountRegionMultiplexer(tableName, "organizations"),
		Columns: []schema.Column{
			{
				Name:       "request_account_id",
				Type:       arrow.BinaryTypes.String,
				Resolver:   client.ResolveAWSAccount,
				PrimaryKey: true,
			},
		},
		Relations: []*schema.Table{
			organizationalUnitParents(),
		},
	}
}

func fetchOUs(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Organizations
	var input organizations.ListRootsInput
	paginator := organizations.NewListRootsPaginator(svc, &input)
	var roots []types.Root
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *organizations.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		roots = append(roots, page.Roots...)
	}

	for _, root := range roots {
		err := getOUs(ctx, meta, svc, res, *root.Id)
		if err != nil {
			return err
		}
	}
	return nil
}

func getOUs(ctx context.Context, meta schema.ClientMeta, accountsApi services.OrganizationsClient, res chan<- any, parentID string) error {
	q := []string{parentID}
	var ou string
	seenOUs := map[string]struct{}{}
	cl := meta.(*client.Client)
	for len(q) > 0 {
		ou, q = q[0], q[1:]
		if _, found := seenOUs[ou]; found {
			continue
		}
		seenOUs[ou] = struct{}{}
		// get OUs directly under this OU, and add them to the queue
		ouPaginator := organizations.NewListChildrenPaginator(accountsApi, &organizations.ListChildrenInput{
			ChildType: types.ChildTypeOrganizationalUnit,
			ParentId:  aws.String(ou),
		})
		for ouPaginator.HasMorePages() {
			output, err := ouPaginator.NextPage(ctx, func(options *organizations.Options) {
				options.Region = cl.Region
			})
			if err != nil {
				return err
			}
			res <- output.Children
			for _, child := range output.Children {
				q = append(q, *child.Id)
			}
		}
	}

	return nil
}

func getOU(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	cl := meta.(*client.Client)
	child := resource.Item.(types.Child)
	svc := cl.Services().Organizations
	ou, err := svc.DescribeOrganizationalUnit(ctx, &organizations.DescribeOrganizationalUnitInput{
		OrganizationalUnitId: child.Id,
	}, func(options *organizations.Options) {
		options.Region = cl.Region
	})
	if err != nil {
		return err
	}
	resource.Item = ou.OrganizationalUnit
	return nil
}
