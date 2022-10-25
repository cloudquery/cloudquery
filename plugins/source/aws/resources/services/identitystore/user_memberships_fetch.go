package identitystore

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/identitystore"
	"github.com/aws/aws-sdk-go-v2/service/identitystore/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchIdentitystoreUserMemberships(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client).Services().Identitystore
	user := parent.Item.(types.User)
	memberId := &types.MemberIdMemberUserId{Value: *user.UserId}
	config := identitystore.ListGroupMembershipsForMemberInput{
		IdentityStoreId: user.IdentityStoreId,
		MemberId:        memberId,
	}
	for {
		response, err := svc.ListGroupMembershipsForMember(ctx, &config)
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
