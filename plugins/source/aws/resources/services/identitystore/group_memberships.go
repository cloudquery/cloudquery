package identitystore

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/identitystore"
	"github.com/aws/aws-sdk-go-v2/service/identitystore/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/transformers"
)

func groupMemberships() *schema.Table {
	tableName := "aws_identitystore_group_memberships"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/singlesignon/latest/IdentityStoreAPIReference/API_GroupMembership.html`,
		Resolver:    fetchIdentitystoreGroupMemberships,
		Transform:   transformers.TransformWithStruct(&types.GroupMembership{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "identitystore"),
		Columns: []schema.Column{
			{
				Name:     "member_id",
				Type:     schema.TypeString,
				Resolver: resolveMemberID,
			},
		},
	}
}

func fetchIdentitystoreGroupMemberships(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	svc := meta.(*client.Client).Services().Identitystore
	group := parent.Item.(types.Group)
	config := identitystore.ListGroupMembershipsInput{
		GroupId:         group.GroupId,
		IdentityStoreId: group.IdentityStoreId,
	}
	paginator := identitystore.NewListGroupMembershipsPaginator(svc, &config)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- page.GroupMemberships
	}
	return nil
}

func resolveMemberID(_ context.Context, meta schema.ClientMeta, r *schema.Resource, c schema.Column) error {
	groupMembership := r.Item.(types.GroupMembership)
	if v, ok := groupMembership.MemberId.(*types.MemberIdMemberUserId); ok {
		return r.Set(c.Name, v.Value)
	}
	return nil
}
