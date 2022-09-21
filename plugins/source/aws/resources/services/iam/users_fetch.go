package iam

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

type AccessKeyWrapper struct {
	types.AccessKeyMetadata
	LastRotated time.Time
}

func fetchIamUsers(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	return client.ListAndDetailResolver(ctx, meta, res, listUsers, fetchUserDetail)
}

func listUsers(ctx context.Context, meta schema.ClientMeta, detailChan chan<- interface{}) error {
	config := iam.ListUsersInput{}
	c := meta.(*client.Client)
	svc := c.Services().IAM
	p := iam.NewListUsersPaginator(svc, &config)
	for p.HasMorePages() {
		response, err := p.NextPage(ctx)
		if err != nil {
			return err
		}
		for _, user := range response.Users {
			detailChan <- user
		}
	}
	return nil
}

func fetchUserDetail(ctx context.Context, meta schema.ClientMeta, resultsChan chan<- interface{}, errorChan chan<- error, listInfo interface{}) {
	c := meta.(*client.Client)

	listUser := listInfo.(types.User)
	svc := meta.(*client.Client).Services().IAM
	userDetail, err := svc.GetUser(ctx, &iam.GetUserInput{
		UserName: aws.String(*listUser.UserName),
	})
	if err != nil {
		if c.IsNotFoundError(err) {
			return
		}
		errorChan <- err
		return
	}
	resultsChan <- userDetail.User
}

func fetchIamUserGroups(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	var config iam.ListGroupsForUserInput
	p := parent.Item.(*types.User)
	svc := meta.(*client.Client).Services().IAM
	config.UserName = p.UserName
	for {
		output, err := svc.ListGroupsForUser(ctx, &config)
		if err != nil {
			return err
		}
		res <- output.Groups
		if output.Marker == nil {
			break
		}
		config.Marker = output.Marker
	}
	return nil
}

func fetchIamUserAccessKeys(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	var config iam.ListAccessKeysInput
	p := parent.Item.(*types.User)
	svc := meta.(*client.Client).Services().IAM
	config.UserName = p.UserName
	for {
		output, err := svc.ListAccessKeys(ctx, &config)
		if err != nil {
			return err
		}

		keys := make([]AccessKeyWrapper, len(output.AccessKeyMetadata))
		for i, key := range output.AccessKeyMetadata {
			switch i {
			case 0:
				rotated := parent.Get("access_key_1_last_rotated")
				if rotated != nil {
					keys[i] = AccessKeyWrapper{key, rotated.(time.Time)}
				} else {
					keys[i] = AccessKeyWrapper{key, *key.CreateDate}
				}
			case 1:
				rotated := parent.Get("access_key_2_last_rotated")
				if rotated != nil {
					keys[i] = AccessKeyWrapper{key, rotated.(time.Time)}
				} else {
					keys[i] = AccessKeyWrapper{key, *key.CreateDate}
				}
			default:
				keys[i] = AccessKeyWrapper{key, time.Time{}}
			}
		}
		res <- keys
		if output.Marker == nil {
			break
		}
		config.Marker = output.Marker
	}
	return nil
}

func postIamUserAccessKeyResolver(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	r := resource.Item.(AccessKeyWrapper)
	if r.AccessKeyId == nil {
		return nil
	}
	svc := meta.(*client.Client).Services().IAM
	output, err := svc.GetAccessKeyLastUsed(ctx, &iam.GetAccessKeyLastUsedInput{AccessKeyId: r.AccessKeyId})
	if err != nil {
		return err
	}
	if output.AccessKeyLastUsed != nil {
		if err := resource.Set("last_used", output.AccessKeyLastUsed.LastUsedDate); err != nil {
			return err
		}
		if err := resource.Set("last_used_service_name", output.AccessKeyLastUsed.ServiceName); err != nil {
			return err
		}
	}
	return nil
}

func fetchIamUserAttachedPolicies(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	var config iam.ListAttachedUserPoliciesInput
	p := parent.Item.(*types.User)
	svc := meta.(*client.Client).Services().IAM
	config.UserName = p.UserName
	for {
		output, err := svc.ListAttachedUserPolicies(ctx, &config)
		if err != nil {
			return err
		}
		res <- output.AttachedPolicies
		if output.Marker == nil {
			break
		}
		config.Marker = output.Marker
	}
	return nil
}
