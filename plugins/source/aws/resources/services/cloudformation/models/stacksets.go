package models

import "github.com/aws/aws-sdk-go-v2/service/cloudformation/types"

type ExpandedSummary struct {
	types.StackSetSummary
	types.CallAs
}

type ExpandedStackSet struct {
	types.StackSet
	types.CallAs
}
type ExpandedStackSetOperationSummary struct {
	types.StackSetOperationSummary
	types.CallAs
}
type ExpandedStackSetOperation struct {
	types.StackSetOperation
	types.CallAs
}
