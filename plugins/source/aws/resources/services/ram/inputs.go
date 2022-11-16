package ram

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ram"
)

func listPrincipalsInput() ram.ListPrincipalsInput {
	return ram.ListPrincipalsInput{
		MaxResults: aws.Int32(500),
	}
}

func listResourcesInput() ram.ListResourcesInput {
	return ram.ListResourcesInput{
		MaxResults: aws.Int32(500),
	}
}

func getResourceShareInvitationsInput() ram.GetResourceShareInvitationsInput {
	return ram.GetResourceShareInvitationsInput{
		MaxResults: aws.Int32(500),
	}
}
