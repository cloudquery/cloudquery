package group

import (
	"context"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/cloudquery/plugins/source/snyk/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/pavel-snyk/snyk-sdk-go/snyk"
)

func groupMembers() *schema.Table {
	return &schema.Table{
		Name:        "snyk_group_members",
		Description: `https://snyk.docs.apiary.io/#reference/groups/group-settings/list-all-members-in-a-group`,
		Resolver:    fetchGroupMembers,
		Multiplex:   client.ByOrganization,
		Transform: transformers.TransformWithStruct(&snyk.GroupMember{},
			transformers.WithPrimaryKeys("ID"),
		),
		Columns: []schema.Column{
			{
				Name:       "group_id",
				Type:       arrow.BinaryTypes.String,
				PrimaryKey: true,
				Resolver:   schema.ParentColumnResolver("id"),
			},
		},
	}
}

func fetchGroupMembers(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	groupID := parent.Item.(*snyk.Group).ID
	result, _, err := c.Groups.ListMembers(ctx, groupID)
	if err != nil {
		return err
	}
	for _, member := range result {
		res <- member
	}
	return nil
}
