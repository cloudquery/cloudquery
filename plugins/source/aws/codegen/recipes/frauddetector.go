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
		{SubService: "batch_imports", Struct: new(types.BatchImport)},
		{SubService: "batch_predictions", Struct: new(types.BatchPrediction)},
		{SubService: "detectors", Struct: new(types.Detector)},
		{SubService: "entity_types", Struct: new(types.EntityType)},
		{SubService: "event_types", Struct: new(types.EventType)},
		{SubService: "external_models", Struct: new(types.ExternalModel)},
		{SubService: "labels", Struct: new(types.Label)},
		{SubService: "models", Struct: new(types.Model)},
		{SubService: "model_versions", Struct: new(types.ModelVersionDetail)},
		{SubService: "outcomes", Struct: new(types.Outcome)},
		{SubService: "rules", Struct: new(types.RuleDetail)},
		{SubService: "variables", Struct: new(types.Variable)},
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
