package users

import (
	"context"

	"github.com/PagerDuty/go-pagerduty"
	"github.com/cloudquery/cloudquery/plugins/source/pagerduty/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
)

func fetchUserContactMethods(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cqClient := meta.(*client.Client)
	concreteParent := parent.Item.(pagerduty.User)

	response, err := cqClient.PagerdutyClient.ListUserContactMethodsWithContext(ctx, concreteParent.ID)
	if err != nil {
		return err
	}

	if len(response.ContactMethods) == 0 {
		return nil
	}

	res <- response.ContactMethods

	return nil
}
