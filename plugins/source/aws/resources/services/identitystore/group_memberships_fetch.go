package identitystore

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/identitystore"
	"github.com/aws/aws-sdk-go-v2/service/identitystore/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchIdentitystoreGroupMemberships(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	svc := meta.(*client.Client).Services().Identitystore
	group := parent.Item.(types.Group)
	config := identitystore.ListGroupMembershipsInput{
		GroupId:         group.GroupId,
		IdentityStoreId: group.IdentityStoreId,
	}
	for {
		response, err := svc.ListGroupMemberships(ctx, &config)
		if err != nil {
			return err
		}
		res <- response.GroupMemberships

		if aws.ToString(response.NextToken) == "" {
			break
		}
		config.NextToken = response.NextToken
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
