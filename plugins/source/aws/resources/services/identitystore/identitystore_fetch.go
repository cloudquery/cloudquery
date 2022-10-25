package identitystore

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/identitystore"
	"github.com/aws/aws-sdk-go-v2/service/ssoadmin"
	types "github.com/aws/aws-sdk-go-v2/service/ssoadmin/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

//var iam_users chan interface{}
//var iam_groups chan interface{}

func getIamInstance(ctx context.Context, meta schema.ClientMeta) (types.InstanceMetadata, error) {
	svc := meta.(*client.Client).Services().SSOAdmin
	config := ssoadmin.ListInstancesInput{}
	response, err := svc.ListInstances(ctx, &config)
	if err == nil {
		for _, i := range response.Instances {
			return i, err
		}
	}
	return types.InstanceMetadata{}, err
}

func fetchIdentitystoreGroups(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	instance, err := getIamInstance(ctx, meta)
	if err != nil {
		return err
	}
	svc := meta.(*client.Client).Services().Identitystore
	config := identitystore.ListGroupsInput{}
	config.IdentityStoreId = instance.IdentityStoreId
	for {
		response, err := svc.ListGroups(ctx, &config)
		if err != nil {
			return err
		}
		res <- response.Groups
		//for _, group := range response.Groups {
		//	res <- group
		//	iam_groups <- group
		//}

		if aws.ToString(response.NextToken) == "" {
			break
		}
		config.NextToken = response.NextToken
	}
	return nil
}
