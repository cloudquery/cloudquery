package identitystore

import (
	"github.com/aws/aws-sdk-go-v2/service/identitystore/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func GroupMemberships() *schema.Table {
	return &schema.Table{
		Name:        "aws_identitystore_group_memberships",
		Description: `https://docs.aws.amazon.com/singlesignon/latest/IdentityStoreAPIReference/API_GroupMembership.html`,
		Resolver:    fetchIdentitystoreGroupMemberships,
		Transform:   transformers.TransformWithStruct(&types.GroupMembership{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer("identitystore"),
	}
}
