package quicksight

import (
	"context"
	"errors"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/quicksight"
	"github.com/aws/aws-sdk-go-v2/service/quicksight/types"
	"github.com/aws/smithy-go"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchQuicksightFolders(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Quicksight
	input := quicksight.ListFoldersInput{
		AwsAccountId: aws.String(cl.AccountID),
	}
	var ae smithy.APIError

	for {
		out, err := svc.ListFolders(ctx, &input)
		if err != nil {
			if errors.As(err, &ae) && ae.ErrorCode() == "UnsupportedUserEditionException" {
				return nil
			}

			return err
		}
		res <- out.FolderSummaryList

		if aws.ToString(out.NextToken) == "" {
			break
		}
		input.NextToken = out.NextToken
	}
	return nil
}

func getFolder(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Quicksight
	item := resource.Item.(types.FolderSummary)

	out, err := svc.DescribeFolder(ctx, &quicksight.DescribeFolderInput{
		AwsAccountId: aws.String(cl.AccountID),
		FolderId:     item.FolderId,
	})
	if err != nil {
		return err
	}

	resource.Item = out.Folder
	return nil
}
