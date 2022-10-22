package recipes

import (
	"github.com/aws/aws-sdk-go-v2/service/frauddetector/types"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
)

func FraudDetectorResources() []*Resource {
	const arnField = "Arn"
	skipFields := []string{arnField}
	extraColumns := append(defaultRegionalColumns,
		codegen.ColumnDefinition{
			Name:     "arn",
			Type:     schema.TypeString,
			Resolver: `schema.PathResolver("` + arnField + `")`,
			Options:  schema.ColumnCreationOptions{PrimaryKey: true},
		},
	)
	tagsCol := codegen.ColumnDefinition{
		Name:     "tags",
		Type:     schema.TypeJSON,
		Resolver: `resolveResourceTags`,
	}

	return []*Resource{
		{
			Service:      "frauddetector",
			SubService:   "batch_imports",
			Struct:       new(types.BatchImport),
			Multiplex:    `client.ServiceAccountRegionMultiplexer("frauddetector")`,
			Description:  "https://docs.aws.amazon.com/frauddetector/latest/api/API_BatchImport.html",
			ExtraColumns: extraColumns,
			SkipFields:   skipFields,
		},
		{
			Service:      "frauddetector",
			SubService:   "batch_predictions",
			Struct:       new(types.BatchPrediction),
			Multiplex:    `client.ServiceAccountRegionMultiplexer("frauddetector")`,
			Description:  "https://docs.aws.amazon.com/frauddetector/latest/api/API_BatchPrediction.html",
			ExtraColumns: extraColumns,
			SkipFields:   skipFields,
		},
		{
			Service:      "frauddetector",
			SubService:   "detectors",
			Struct:       new(types.Detector),
			Multiplex:    `client.ServiceAccountRegionMultiplexer("frauddetector")`,
			Description:  "https://docs.aws.amazon.com/frauddetector/latest/api/API_Detector.html",
			ExtraColumns: append(extraColumns, tagsCol),
			SkipFields:   skipFields,
			Relations:    []string{"Rules()"},
		},
		{
			Service:      "frauddetector",
			SubService:   "entity_types",
			Struct:       new(types.EntityType),
			Multiplex:    `client.ServiceAccountRegionMultiplexer("frauddetector")`,
			Description:  "https://docs.aws.amazon.com/frauddetector/latest/api/API_EntityType.html",
			ExtraColumns: append(extraColumns, tagsCol),
			SkipFields:   skipFields,
		},
		{
			Service:      "frauddetector",
			SubService:   "event_types",
			Struct:       new(types.EventType),
			Multiplex:    `client.ServiceAccountRegionMultiplexer("frauddetector")`,
			Description:  "https://docs.aws.amazon.com/frauddetector/latest/api/API_EventType.html",
			ExtraColumns: append(extraColumns, tagsCol),
			SkipFields:   skipFields,
		},
		{
			Service:      "frauddetector",
			SubService:   "external_models",
			Struct:       new(types.ExternalModel),
			Multiplex:    `client.ServiceAccountRegionMultiplexer("frauddetector")`,
			Description:  "https://docs.aws.amazon.com/frauddetector/latest/api/API_ExternalModel.html",
			ExtraColumns: extraColumns,
			SkipFields:   skipFields,
		},
		{
			Service:      "frauddetector",
			SubService:   "labels",
			Struct:       new(types.Label),
			Multiplex:    `client.ServiceAccountRegionMultiplexer("frauddetector")`,
			Description:  "https://docs.aws.amazon.com/frauddetector/latest/api/API_Label.html",
			ExtraColumns: append(extraColumns, tagsCol),
			SkipFields:   skipFields,
		},
		{
			Service:      "frauddetector",
			SubService:   "models",
			Struct:       new(types.Model),
			Multiplex:    `client.ServiceAccountRegionMultiplexer("frauddetector")`,
			Description:  "https://docs.aws.amazon.com/frauddetector/latest/api/API_Model.html",
			ExtraColumns: extraColumns,
			SkipFields:   skipFields,
			Relations:    []string{"ModelVersions()"},
		},
		{
			Service:      "frauddetector",
			SubService:   "model_versions", // relation for models
			Struct:       new(types.ModelVersionDetail),
			Multiplex:    "", // we skip multiplexing here as it's a relation
			Description:  "https://docs.aws.amazon.com/frauddetector/latest/api/API_ModelVersionDetail.html",
			ExtraColumns: extraColumns,
			SkipFields:   skipFields,
		},
		{
			Service:      "frauddetector",
			SubService:   "outcomes",
			Struct:       new(types.Outcome),
			Multiplex:    `client.ServiceAccountRegionMultiplexer("frauddetector")`,
			Description:  "https://docs.aws.amazon.com/frauddetector/latest/api/API_Outcome.html",
			ExtraColumns: append(extraColumns, tagsCol),
			SkipFields:   skipFields,
		},
		{
			Service:      "frauddetector",
			SubService:   "rules", // relation for detectors
			Struct:       new(types.RuleDetail),
			Multiplex:    "", // we skip multiplexing here as it's a relation
			Description:  "https://docs.aws.amazon.com/frauddetector/latest/api/API_RuleDetail.html",
			ExtraColumns: extraColumns,
			SkipFields:   skipFields,
		},
		{
			Service:      "frauddetector",
			SubService:   "variables",
			Struct:       new(types.Variable),
			Multiplex:    `client.ServiceAccountRegionMultiplexer("frauddetector")`,
			Description:  "https://docs.aws.amazon.com/frauddetector/latest/api/API_Variable.html",
			ExtraColumns: append(extraColumns, tagsCol),
			SkipFields:   skipFields,
		},
	}
}
