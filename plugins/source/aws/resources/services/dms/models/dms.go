package models

import "github.com/aws/aws-sdk-go-v2/service/databasemigrationservice/types"

type ReplicationInstanceWrapper struct {
	types.ReplicationInstance
	Tags map[string]interface{}
}
