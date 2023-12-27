package models

import "github.com/aws/aws-sdk-go-v2/service/rds/types"

type ExtendedSnapshots struct {
	types.DBSnapshot
	Attributes []ExtendedAttributes
}

type ExtendedAttributes struct {
	types.DBSnapshotAttribute
}
