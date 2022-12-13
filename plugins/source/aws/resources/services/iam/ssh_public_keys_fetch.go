package iam

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchIamSshPublicKeys(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	user := parent.Item.(*types.User)
	config := iam.ListSSHPublicKeysInput{UserName: user.UserName}
	svc := meta.(*client.Client).Services().Iam
	for {
		response, err := svc.ListSSHPublicKeys(ctx, &config)
		if err != nil {
			return err
		}
		res <- response.SSHPublicKeys
		if aws.ToString(response.Marker) == "" {
			break
		}
		config.Marker = response.Marker
	}
	return nil
}
