package models

import "github.com/aws/aws-sdk-go-v2/service/ecs/types"

type TaskDefinitionWrapper struct {
	*types.TaskDefinition
	Tags []types.Tag
}
