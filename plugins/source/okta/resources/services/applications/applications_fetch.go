package applications

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/okta/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/okta/okta-sdk-golang/v3/okta"
)

func fetchApplications(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)

	req := cl.ApplicationApi.ListApplications(ctx).Limit(200)
	items, resp, err := cl.ApplicationApi.ListApplicationsExecute(req)
	if err != nil {
		return err
	}
	if len(items) == 0 {
		return nil
	}

	convertAndPush := func(items []okta.ListApplications200ResponseInner) {
		list := make([]*okta.Application, 0, len(items))
		for i := range items {
			if aa := appToApp(&items[i]); aa != nil {
				list = append(list, aa)
			}
		}
		res <- list
	}

	convertAndPush(items)

	for resp != nil && resp.HasNextPage() {
		var nextItems []okta.ListApplications200ResponseInner
		resp, err = resp.Next(&nextItems)
		if err != nil {
			return err
		}
		convertAndPush(nextItems)
	}
	return nil
}

func appToApp(obj *okta.ListApplications200ResponseInner) *okta.Application {
	// order copied from ListApplications200ResponseInner.GetActualInstance()
	if obj == nil {
		return nil
	}
	if obj.AutoLoginApplication != nil {
		return &obj.AutoLoginApplication.Application
	}

	if obj.BasicAuthApplication != nil {
		return &obj.BasicAuthApplication.Application
	}

	if obj.BookmarkApplication != nil {
		return &obj.BookmarkApplication.Application
	}

	if obj.BrowserPluginApplication != nil {
		return &obj.BrowserPluginApplication.Application
	}

	if obj.OpenIdConnectApplication != nil {
		return &obj.OpenIdConnectApplication.Application
	}

	if obj.SamlApplication != nil {
		return &obj.SamlApplication.Application
	}

	if obj.SecurePasswordStoreApplication != nil {
		return &obj.SecurePasswordStoreApplication.Application
	}

	if obj.WsFederationApplication != nil {
		return &obj.WsFederationApplication.Application
	}

	// all schemas are nil
	return nil
}
