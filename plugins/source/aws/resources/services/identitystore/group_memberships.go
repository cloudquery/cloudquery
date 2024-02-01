package identitystore

import (
	"context"

	"github.com/apache/arrow/go/v15/arrow"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/arn"
	"github.com/aws/aws-sdk-go-v2/service/identitystore"
	"github.com/aws/aws-sdk-go-v2/service/identitystore/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func groupMemberships() *schema.Table {
	tableName := "aws_identitystore_group_memberships"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/singlesignon/latest/IdentityStoreAPIReference/API_GroupMembership.html`,
		Resolver:    fetchGroupMemberships,
		Transform:   transformers.TransformWithStruct(&types.GroupMembership{}),
		Columns: []schema.Column{
			client.RequestAccountIDColumn(true),
			client.RequestRegionColumn(true),
			{
				Name:                "group_arn",
				Type:                arrow.BinaryTypes.String,
				Resolver:            schema.ParentColumnResolver("arn"),
				PrimaryKeyComponent: true,
			},
			{
				Name:                "arn",
				Type:                arrow.BinaryTypes.String,
				Resolver:            resolveMembershipARN,
				PrimaryKeyComponent: true,
			},
			{
				Name:     "member_id",
				Type:     arrow.BinaryTypes.String,
				Resolver: resolveMemberID,
			},
		},
	}
}

func fetchGroupMemberships(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceIdentitystore).Identitystore
	group := parent.Item.(types.Group)
	config := identitystore.ListGroupMembershipsInput{
		GroupId:         group.GroupId,
		IdentityStoreId: group.IdentityStoreId,
	}
	paginator := identitystore.NewListGroupMembershipsPaginator(svc, &config)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *identitystore.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- page.GroupMemberships
	}
	return nil
}

func resolveMemberID(_ context.Context, _ schema.ClientMeta, r *schema.Resource, c schema.Column) error {
	groupMembership := r.Item.(types.GroupMembership)
	switch val := groupMembership.MemberId.(type) {
	case *types.MemberIdMemberUserId:
		return r.Set(c.Name, val.Value)
	case *types.UnknownUnionMember:
		return r.Set(c.Name, val.Tag)
	default:
		return nil
	}
}

func resolveMembershipARN(_ context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	return resource.Set(c.Name, membershipARN(cl, aws.ToString(resource.Item.(types.GroupMembership).MembershipId)))
}

func membershipARN(cl *client.Client, membershipID string) string {
	// https: //docs.aws.amazon.com/service-authorization/latest/reference/list_awsidentitystore.html#awsidentitystore-resources-for-iam-policies
	return arn.ARN{
		Partition: cl.Partition,
		Service:   string(client.IdentitystoreService),
		Resource:  "membership/" + membershipID,
	}.String()
}
