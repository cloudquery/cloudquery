package recipes

import (
	"reflect"
	"runtime"
	"strings"

	aiplatform "cloud.google.com/go/aiplatform/apiv1"
	pb "cloud.google.com/go/aiplatform/apiv1/aiplatformpb"
	longrunningpb "cloud.google.com/go/longrunning/autogen/longrunningpb"
	locationpb "google.golang.org/genproto/googleapis/cloud/location"
	"google.golang.org/protobuf/types/known/structpb"
)

func getLocationsServiceName(f interface{}) string {
	fName := strings.Split(runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name(), "apiv1.")
	return strings.ToLower(strings.TrimSuffix(strings.TrimPrefix(fName[1], "New"), "Client"))
}

func isFieldMockable(field reflect.StructField) bool {
	nonMockables := []any{&structpb.Value{}, &structpb.Struct{}, &pb.Model{}, &pb.PipelineJob_RuntimeConfig{}}
	for _, nonMockable := range nonMockables {
		if field.Type == reflect.TypeOf(nonMockable) {
			return false
		}
	}

	return true
}

func init() {
	resources := []*Resource{
		{
			SubService:     "batch_prediction_jobs",
			Struct:         &pb.BatchPredictionJob{},
			NewFunction:    aiplatform.NewJobClient,
			Description:    "https://cloud.google.com/vertex-ai/docs/reference/rest/v1/projects.locations.batchPredictionJobs#BatchPredictionJob",
			RegisterServer: pb.RegisterJobServiceServer,
		},
		{
			SubService:     "custom_jobs",
			Struct:         &pb.CustomJob{},
			NewFunction:    aiplatform.NewJobClient,
			Description:    "https://cloud.google.com/vertex-ai/docs/reference/rest/v1/projects.locations.customJobs#CustomJob",
			RegisterServer: pb.RegisterJobServiceServer,
		},
		{
			SubService:     "data_labeling_jobs",
			Struct:         &pb.DataLabelingJob{},
			NewFunction:    aiplatform.NewJobClient,
			Description:    "https://cloud.google.com/vertex-ai/docs/reference/rest/v1/projects.locations.dataLabelingJobs#DataLabelingJob",
			RegisterServer: pb.RegisterJobServiceServer,
		},
		{
			SubService:     "datasets",
			Struct:         &pb.Dataset{},
			NewFunction:    aiplatform.NewDatasetClient,
			Description:    "https://cloud.google.com/vertex-ai/docs/reference/rest/v1/projects.locations.datasets#Dataset",
			RegisterServer: pb.RegisterDatasetServiceServer,
		},
		{
			SubService:     "endpoints",
			Struct:         &pb.Endpoint{},
			NewFunction:    aiplatform.NewEndpointClient,
			Description:    "https://cloud.google.com/vertex-ai/docs/reference/rest/v1/projects.locations.endpoints#Endpoint",
			RegisterServer: pb.RegisterEndpointServiceServer,
		},
		{
			SubService:     "featurestores",
			Struct:         &pb.Featurestore{},
			NewFunction:    aiplatform.NewFeaturestoreClient,
			Description:    "https://cloud.google.com/vertex-ai/docs/reference/rest/v1/projects.locations.featurestores#Featurestore",
			RegisterServer: pb.RegisterFeaturestoreServiceServer,
		},
		{
			SubService:     "hyperparameter_tuning_jobs",
			Struct:         &pb.HyperparameterTuningJob{},
			NewFunction:    aiplatform.NewJobClient,
			Description:    "https://cloud.google.com/vertex-ai/docs/reference/rest/v1/projects.locations.hyperparameterTuningJobs#HyperparameterTuningJob",
			RegisterServer: pb.RegisterJobServiceServer,
		},
		{
			SubService:       "index_endpoints",
			Struct:           &pb.IndexEndpoint{},
			NewFunction:      aiplatform.NewIndexEndpointClient,
			Description:      "https://cloud.google.com/vertex-ai/docs/reference/rest/v1/projects.locations.indexEndpoints#IndexEndpoint",
			RegisterServer:   pb.RegisterIndexEndpointServiceServer,
			ParentFilterFunc: "filterIndexesLocations",
		},
		{
			SubService:       "indexes",
			Struct:           &pb.Index{},
			NewFunction:      aiplatform.NewIndexClient,
			Description:      "https://cloud.google.com/vertex-ai/docs/reference/rest/v1/projects.locations.indexes#Index",
			RegisterServer:   pb.RegisterIndexServiceServer,
			ParentFilterFunc: "filterIndexesLocations",
		},
		{
			SubService:     "metadata_stores",
			Struct:         &pb.MetadataStore{},
			NewFunction:    aiplatform.NewMetadataClient,
			ListFunction:   (&aiplatform.MetadataClient{}).ListMetadataStores,
			Description:    "https://cloud.google.com/vertex-ai/docs/reference/rest/v1/projects.locations.metadataStores#MetadataStore",
			RegisterServer: pb.RegisterMetadataServiceServer,
		},
		{
			SubService:     "model_deployment_monitoring_jobs",
			Struct:         &pb.ModelDeploymentMonitoringJob{},
			NewFunction:    aiplatform.NewJobClient,
			Description:    "https://cloud.google.com/vertex-ai/docs/reference/rest/v1/projects.locations.modelDeploymentMonitoringJobs#ModelDeploymentMonitoringJob",
			RegisterServer: pb.RegisterJobServiceServer,
		},
		{
			SubService:     "models",
			Struct:         &pb.Model{},
			NewFunction:    aiplatform.NewModelClient,
			ListFunction:   (&aiplatform.ModelClient{}).ListModels,
			Description:    "https://cloud.google.com/vertex-ai/docs/reference/rest/v1/projects.locations.models#Model",
			RegisterServer: pb.RegisterModelServiceServer,
		},

		{
			SubService:     "pipeline_jobs",
			Struct:         &pb.PipelineJob{},
			NewFunction:    aiplatform.NewPipelineClient,
			Description:    "https://cloud.google.com/vertex-ai/docs/reference/rest/v1/projects.locations.pipelineJobs#PipelineJob",
			RegisterServer: pb.RegisterPipelineServiceServer,
		},
		{
			SubService:     "specialist_pools",
			Struct:         &pb.SpecialistPool{},
			NewFunction:    aiplatform.NewSpecialistPoolClient,
			Description:    "https://cloud.google.com/vertex-ai/docs/reference/rest/v1/projects.locations.specialistPools#SpecialistPool",
			RegisterServer: pb.RegisterSpecialistPoolServiceServer,
		},
		{
			SubService:       "studies",
			Struct:           &pb.Study{},
			NewFunction:      aiplatform.NewVizierClient,
			Description:      "https://cloud.google.com/vertex-ai/docs/reference/rest/v1/projects.locations.studies#Study",
			RegisterServer:   pb.RegisterVizierServiceServer,
			ParentFilterFunc: "filterStudiesLocation",
		},
		{
			SubService:     "tensorboards",
			Struct:         &pb.Tensorboard{},
			NewFunction:    aiplatform.NewTensorboardClient,
			Description:    "https://cloud.google.com/vertex-ai/docs/reference/rest/v1/projects.locations.tensorboards#Tensorboard",
			RegisterServer: pb.RegisterTensorboardServiceServer,
		},
		{
			SubService:           "training_pipelines",
			Struct:               &pb.TrainingPipeline{},
			NewFunction:          aiplatform.NewPipelineClient,
			Description:          "https://cloud.google.com/vertex-ai/docs/reference/rest/v1/projects.locations.trainingPipelines#TrainingPipeline",
			RegisterServer:       pb.RegisterPipelineServiceServer,
			IgnoreInTestsColumns: []string{"model_to_upload"},
		},
	}

	locationsMap := map[string]*Resource{}
	locationResources := []*Resource{}
	for _, resource := range resources {
		resource.ProtobufImport = "cloud.google.com/go/aiplatform/apiv1/aiplatformpb"
		resource.ChildTable = true
		resource.RequestStructFields = `Parent: parent.Item.(*location.Location).Name,`
		resource.MockImports = []string{"google.golang.org/api/option", "google.golang.org/genproto/googleapis/cloud/location"}
		resource.ClientOptions = []string{`option.WithEndpoint(parent.Item.(*location.Location).LocationId + "-aiplatform.googleapis.com:443")`}
		if resource.ParentFilterFunc == "" {
			resource.ParentFilterFunc = "filterLocation"
		}
		resource.SkipMock = true

		locationSubService := getLocationsServiceName(resource.NewFunction) + "_locations"
		locationResource := locationsMap[locationSubService]
		if locationResource == nil {
			locationResource = &Resource{}
			locationResource.SubService = locationSubService
			locationResource.Struct = &locationpb.Location{}
			locationResource.PrimaryKeys = []string{ProjectIdColumn.Name, "name"}
			locationResource.NewFunction = resource.NewFunction
			locationResource.RequestStructFields = `Name: "projects/" + c.ProjectId,`
			locationResource.Description = "https://cloud.google.com/api-gateway/docs/reference/rest/v1/projects.locations#Location"
			locationResource.RegisterServer = locationpb.RegisterLocationsServer
			locationResource.ProtobufImport = "google.golang.org/genproto/googleapis/cloud/location"
			locationResource.ClientOptions = []string{`option.WithEndpoint("us-central1-aiplatform.googleapis.com:443")`}
			locationResource.MockImports = []string{"google.golang.org/api/option"}
			locationResource.RelationsTestData = RelationsTestData{
				RegisterServer: resource.RegisterServer,
			}

			locationsMap[locationSubService] = locationResource
			locationResources = append(locationResources, locationResource)
		}
		locationResource.Relations = append(locationResource.Relations, Caser.ToPascal(resource.SubService)+"()")
	}

	operationsResource := &Resource{
		SubService:          "operations",
		Struct:              &longrunningpb.Operation{},
		NewFunction:         aiplatform.NewPipelineClient,
		RequestStructFields: `Name: "projects/" + c.ProjectId + "/locations/-",`,
		Description:         "https://cloud.google.com/vertex-ai/docs/reference/rest/v1/projects.locations.operations#Operation",
		RegisterServer:      longrunningpb.RegisterOperationsServer,
		ProtobufImport:      "cloud.google.com/go/longrunning/autogen/longrunningpb",
		ClientOptions:       []string{`option.WithEndpoint("us-central1-aiplatform.googleapis.com:443")`},
		MockImports:         []string{"google.golang.org/api/option"},
	}

	allResources := append(resources, locationResources...)
	allResources = append(allResources, operationsResource)

	for _, resource := range allResources {
		resource.PrimaryKeys = []string{ProjectIdColumn.Name, "name"}
		resource.Service = "aiplatform"
		resource.Template = "newapi_list"
		resource.MockTemplate = "newapi_list_grpc_mock"
		resource.MockImports = append(resource.MockImports, "cloud.google.com/go/aiplatform/apiv1")

		structElem := reflect.TypeOf(resource.Struct).Elem()
		for i := 0; i < structElem.NumField(); i++ {
			field := structElem.Field(i)
			if isFieldMockable(field) {
				continue
			}
			resource.IgnoreInTestsColumns = append(resource.IgnoreInTestsColumns, Caser.ToSnake(field.Name))
		}
	}

	Resources = append(Resources, allResources...)
}
