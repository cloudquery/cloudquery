package recipes

import (
	"github.com/aws/aws-sdk-go-v2/service/frauddetector/types"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
)

func FraudDetectorResources() []*Resource {
	const arnField = "Arn"
	arnColumn := codegen.ColumnDefinition{
		Name:     "arn",
		Type:     schema.TypeString,
		Resolver: `schema.PathResolver("` + arnField + `")`,
		Options:  schema.ColumnCreationOptions{PrimaryKey: true},
	}

	resources := []*Resource{
		{
			SubService:  "batch_imports",
			Struct:      new(types.BatchImport),
			Description: "https://docs.aws.amazon.com/frauddetector/latest/api/API_BatchImport.html",
		},
		{
			SubService:  "batch_predictions",
			Struct:      new(types.BatchPrediction),
			Description: "https://docs.aws.amazon.com/frauddetector/latest/api/API_BatchPrediction.html",
		},
		{
			SubService:  "detectors",
			Struct:      new(types.Detector),
			Description: "https://docs.aws.amazon.com/frauddetector/latest/api/API_Detector.html",
		},
		{
			SubService:  "entity_types",
			Struct:      new(types.EntityType),
			Description: "https://docs.aws.amazon.com/frauddetector/latest/api/API_EntityType.html",
		},
		{
			SubService:  "event_types",
			Struct:      new(types.EventType),
			Description: "https://docs.aws.amazon.com/frauddetector/latest/api/API_EventType.html",
		},
		{
			SubService:  "external_models",
			Struct:      new(types.ExternalModel),
			Description: "https://docs.aws.amazon.com/frauddetector/latest/api/API_ExternalModel.html",
		},
		{
			SubService:  "labels",
			Struct:      new(types.Label),
			Description: "https://docs.aws.amazon.com/frauddetector/latest/api/API_Label.html",
		},
		{
			SubService:  "models",
			Struct:      new(types.Model),
			Description: "https://docs.aws.amazon.com/frauddetector/latest/api/API_Model.html",
		},
		{
			SubService:  "model_versions",
			Struct:      new(types.ModelVersionDetail),
			Description: "https://docs.aws.amazon.com/frauddetector/latest/api/API_ModelVersionDetail.html",
		},
		{
			SubService:  "outcomes",
			Struct:      new(types.Outcome),
			Description: "https://docs.aws.amazon.com/frauddetector/latest/api/API_Outcome.html",
		},
		{
			SubService:  "rules",
			Struct:      new(types.RuleDetail),
			Description: "https://docs.aws.amazon.com/frauddetector/latest/api/API_RuleDetail.html",
		},
		{
			SubService:  "variables",
			Struct:      new(types.Variable),
			Description: "https://docs.aws.amazon.com/frauddetector/latest/api/API_Variable.html",
		},
	}

	// set default values
	for _, r := range resources {
		r.Service = "frauddetector"
		r.Multiplex = `client.ServiceAccountRegionMultiplexer("frauddetector")`
		r.ExtraColumns = append(r.ExtraColumns, arnColumn)
		r.SkipFields = append(r.SkipFields, arnField)
	}
	return resources
}
