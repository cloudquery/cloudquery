package models

import "github.com/aws/aws-sdk-go-v2/service/resourcegroups/types"

// ResourceGroupWrapper fields are extracted from types.ResourceGroup and types.ResourceQuery
type ResourceGroupWrapper struct {
	*types.Group
	*types.ResourceQuery
}
