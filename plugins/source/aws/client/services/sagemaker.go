// Code generated by codegen; DO NOT EDIT.
package services

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker"
	
)

//go:generate mockgen -package=mocks -destination=../mocks/sagemaker.go -source=sagemaker.go SagemakerClient
type SagemakerClient interface {
	DescribeAction(context.Context, *sagemaker.DescribeActionInput, ...func(*sagemaker.Options)) (*sagemaker.DescribeActionOutput, error)
	DescribeAlgorithm(context.Context, *sagemaker.DescribeAlgorithmInput, ...func(*sagemaker.Options)) (*sagemaker.DescribeAlgorithmOutput, error)
	DescribeApp(context.Context, *sagemaker.DescribeAppInput, ...func(*sagemaker.Options)) (*sagemaker.DescribeAppOutput, error)
	DescribeAppImageConfig(context.Context, *sagemaker.DescribeAppImageConfigInput, ...func(*sagemaker.Options)) (*sagemaker.DescribeAppImageConfigOutput, error)
	DescribeArtifact(context.Context, *sagemaker.DescribeArtifactInput, ...func(*sagemaker.Options)) (*sagemaker.DescribeArtifactOutput, error)
	DescribeAutoMLJob(context.Context, *sagemaker.DescribeAutoMLJobInput, ...func(*sagemaker.Options)) (*sagemaker.DescribeAutoMLJobOutput, error)
	DescribeCodeRepository(context.Context, *sagemaker.DescribeCodeRepositoryInput, ...func(*sagemaker.Options)) (*sagemaker.DescribeCodeRepositoryOutput, error)
	DescribeCompilationJob(context.Context, *sagemaker.DescribeCompilationJobInput, ...func(*sagemaker.Options)) (*sagemaker.DescribeCompilationJobOutput, error)
	DescribeContext(context.Context, *sagemaker.DescribeContextInput, ...func(*sagemaker.Options)) (*sagemaker.DescribeContextOutput, error)
	DescribeDataQualityJobDefinition(context.Context, *sagemaker.DescribeDataQualityJobDefinitionInput, ...func(*sagemaker.Options)) (*sagemaker.DescribeDataQualityJobDefinitionOutput, error)
	DescribeDevice(context.Context, *sagemaker.DescribeDeviceInput, ...func(*sagemaker.Options)) (*sagemaker.DescribeDeviceOutput, error)
	DescribeDeviceFleet(context.Context, *sagemaker.DescribeDeviceFleetInput, ...func(*sagemaker.Options)) (*sagemaker.DescribeDeviceFleetOutput, error)
	DescribeDomain(context.Context, *sagemaker.DescribeDomainInput, ...func(*sagemaker.Options)) (*sagemaker.DescribeDomainOutput, error)
	DescribeEdgeDeploymentPlan(context.Context, *sagemaker.DescribeEdgeDeploymentPlanInput, ...func(*sagemaker.Options)) (*sagemaker.DescribeEdgeDeploymentPlanOutput, error)
	DescribeEdgePackagingJob(context.Context, *sagemaker.DescribeEdgePackagingJobInput, ...func(*sagemaker.Options)) (*sagemaker.DescribeEdgePackagingJobOutput, error)
	DescribeEndpoint(context.Context, *sagemaker.DescribeEndpointInput, ...func(*sagemaker.Options)) (*sagemaker.DescribeEndpointOutput, error)
	DescribeEndpointConfig(context.Context, *sagemaker.DescribeEndpointConfigInput, ...func(*sagemaker.Options)) (*sagemaker.DescribeEndpointConfigOutput, error)
	DescribeExperiment(context.Context, *sagemaker.DescribeExperimentInput, ...func(*sagemaker.Options)) (*sagemaker.DescribeExperimentOutput, error)
	DescribeFeatureGroup(context.Context, *sagemaker.DescribeFeatureGroupInput, ...func(*sagemaker.Options)) (*sagemaker.DescribeFeatureGroupOutput, error)
	DescribeFeatureMetadata(context.Context, *sagemaker.DescribeFeatureMetadataInput, ...func(*sagemaker.Options)) (*sagemaker.DescribeFeatureMetadataOutput, error)
	DescribeFlowDefinition(context.Context, *sagemaker.DescribeFlowDefinitionInput, ...func(*sagemaker.Options)) (*sagemaker.DescribeFlowDefinitionOutput, error)
	DescribeHub(context.Context, *sagemaker.DescribeHubInput, ...func(*sagemaker.Options)) (*sagemaker.DescribeHubOutput, error)
	DescribeHubContent(context.Context, *sagemaker.DescribeHubContentInput, ...func(*sagemaker.Options)) (*sagemaker.DescribeHubContentOutput, error)
	DescribeHumanTaskUi(context.Context, *sagemaker.DescribeHumanTaskUiInput, ...func(*sagemaker.Options)) (*sagemaker.DescribeHumanTaskUiOutput, error)
	DescribeHyperParameterTuningJob(context.Context, *sagemaker.DescribeHyperParameterTuningJobInput, ...func(*sagemaker.Options)) (*sagemaker.DescribeHyperParameterTuningJobOutput, error)
	DescribeImage(context.Context, *sagemaker.DescribeImageInput, ...func(*sagemaker.Options)) (*sagemaker.DescribeImageOutput, error)
	DescribeImageVersion(context.Context, *sagemaker.DescribeImageVersionInput, ...func(*sagemaker.Options)) (*sagemaker.DescribeImageVersionOutput, error)
	DescribeInferenceExperiment(context.Context, *sagemaker.DescribeInferenceExperimentInput, ...func(*sagemaker.Options)) (*sagemaker.DescribeInferenceExperimentOutput, error)
	DescribeInferenceRecommendationsJob(context.Context, *sagemaker.DescribeInferenceRecommendationsJobInput, ...func(*sagemaker.Options)) (*sagemaker.DescribeInferenceRecommendationsJobOutput, error)
	DescribeLabelingJob(context.Context, *sagemaker.DescribeLabelingJobInput, ...func(*sagemaker.Options)) (*sagemaker.DescribeLabelingJobOutput, error)
	DescribeLineageGroup(context.Context, *sagemaker.DescribeLineageGroupInput, ...func(*sagemaker.Options)) (*sagemaker.DescribeLineageGroupOutput, error)
	DescribeModel(context.Context, *sagemaker.DescribeModelInput, ...func(*sagemaker.Options)) (*sagemaker.DescribeModelOutput, error)
	DescribeModelBiasJobDefinition(context.Context, *sagemaker.DescribeModelBiasJobDefinitionInput, ...func(*sagemaker.Options)) (*sagemaker.DescribeModelBiasJobDefinitionOutput, error)
	DescribeModelCard(context.Context, *sagemaker.DescribeModelCardInput, ...func(*sagemaker.Options)) (*sagemaker.DescribeModelCardOutput, error)
	DescribeModelCardExportJob(context.Context, *sagemaker.DescribeModelCardExportJobInput, ...func(*sagemaker.Options)) (*sagemaker.DescribeModelCardExportJobOutput, error)
	DescribeModelExplainabilityJobDefinition(context.Context, *sagemaker.DescribeModelExplainabilityJobDefinitionInput, ...func(*sagemaker.Options)) (*sagemaker.DescribeModelExplainabilityJobDefinitionOutput, error)
	DescribeModelPackage(context.Context, *sagemaker.DescribeModelPackageInput, ...func(*sagemaker.Options)) (*sagemaker.DescribeModelPackageOutput, error)
	DescribeModelPackageGroup(context.Context, *sagemaker.DescribeModelPackageGroupInput, ...func(*sagemaker.Options)) (*sagemaker.DescribeModelPackageGroupOutput, error)
	DescribeModelQualityJobDefinition(context.Context, *sagemaker.DescribeModelQualityJobDefinitionInput, ...func(*sagemaker.Options)) (*sagemaker.DescribeModelQualityJobDefinitionOutput, error)
	DescribeMonitoringSchedule(context.Context, *sagemaker.DescribeMonitoringScheduleInput, ...func(*sagemaker.Options)) (*sagemaker.DescribeMonitoringScheduleOutput, error)
	DescribeNotebookInstance(context.Context, *sagemaker.DescribeNotebookInstanceInput, ...func(*sagemaker.Options)) (*sagemaker.DescribeNotebookInstanceOutput, error)
	DescribeNotebookInstanceLifecycleConfig(context.Context, *sagemaker.DescribeNotebookInstanceLifecycleConfigInput, ...func(*sagemaker.Options)) (*sagemaker.DescribeNotebookInstanceLifecycleConfigOutput, error)
	DescribePipeline(context.Context, *sagemaker.DescribePipelineInput, ...func(*sagemaker.Options)) (*sagemaker.DescribePipelineOutput, error)
	DescribePipelineDefinitionForExecution(context.Context, *sagemaker.DescribePipelineDefinitionForExecutionInput, ...func(*sagemaker.Options)) (*sagemaker.DescribePipelineDefinitionForExecutionOutput, error)
	DescribePipelineExecution(context.Context, *sagemaker.DescribePipelineExecutionInput, ...func(*sagemaker.Options)) (*sagemaker.DescribePipelineExecutionOutput, error)
	DescribeProcessingJob(context.Context, *sagemaker.DescribeProcessingJobInput, ...func(*sagemaker.Options)) (*sagemaker.DescribeProcessingJobOutput, error)
	DescribeProject(context.Context, *sagemaker.DescribeProjectInput, ...func(*sagemaker.Options)) (*sagemaker.DescribeProjectOutput, error)
	DescribeSpace(context.Context, *sagemaker.DescribeSpaceInput, ...func(*sagemaker.Options)) (*sagemaker.DescribeSpaceOutput, error)
	DescribeStudioLifecycleConfig(context.Context, *sagemaker.DescribeStudioLifecycleConfigInput, ...func(*sagemaker.Options)) (*sagemaker.DescribeStudioLifecycleConfigOutput, error)
	DescribeSubscribedWorkteam(context.Context, *sagemaker.DescribeSubscribedWorkteamInput, ...func(*sagemaker.Options)) (*sagemaker.DescribeSubscribedWorkteamOutput, error)
	DescribeTrainingJob(context.Context, *sagemaker.DescribeTrainingJobInput, ...func(*sagemaker.Options)) (*sagemaker.DescribeTrainingJobOutput, error)
	DescribeTransformJob(context.Context, *sagemaker.DescribeTransformJobInput, ...func(*sagemaker.Options)) (*sagemaker.DescribeTransformJobOutput, error)
	DescribeTrial(context.Context, *sagemaker.DescribeTrialInput, ...func(*sagemaker.Options)) (*sagemaker.DescribeTrialOutput, error)
	DescribeTrialComponent(context.Context, *sagemaker.DescribeTrialComponentInput, ...func(*sagemaker.Options)) (*sagemaker.DescribeTrialComponentOutput, error)
	DescribeUserProfile(context.Context, *sagemaker.DescribeUserProfileInput, ...func(*sagemaker.Options)) (*sagemaker.DescribeUserProfileOutput, error)
	DescribeWorkforce(context.Context, *sagemaker.DescribeWorkforceInput, ...func(*sagemaker.Options)) (*sagemaker.DescribeWorkforceOutput, error)
	DescribeWorkteam(context.Context, *sagemaker.DescribeWorkteamInput, ...func(*sagemaker.Options)) (*sagemaker.DescribeWorkteamOutput, error)
	GetDeviceFleetReport(context.Context, *sagemaker.GetDeviceFleetReportInput, ...func(*sagemaker.Options)) (*sagemaker.GetDeviceFleetReportOutput, error)
	GetLineageGroupPolicy(context.Context, *sagemaker.GetLineageGroupPolicyInput, ...func(*sagemaker.Options)) (*sagemaker.GetLineageGroupPolicyOutput, error)
	GetModelPackageGroupPolicy(context.Context, *sagemaker.GetModelPackageGroupPolicyInput, ...func(*sagemaker.Options)) (*sagemaker.GetModelPackageGroupPolicyOutput, error)
	GetSagemakerServicecatalogPortfolioStatus(context.Context, *sagemaker.GetSagemakerServicecatalogPortfolioStatusInput, ...func(*sagemaker.Options)) (*sagemaker.GetSagemakerServicecatalogPortfolioStatusOutput, error)
	GetSearchSuggestions(context.Context, *sagemaker.GetSearchSuggestionsInput, ...func(*sagemaker.Options)) (*sagemaker.GetSearchSuggestionsOutput, error)
	ListActions(context.Context, *sagemaker.ListActionsInput, ...func(*sagemaker.Options)) (*sagemaker.ListActionsOutput, error)
	ListAlgorithms(context.Context, *sagemaker.ListAlgorithmsInput, ...func(*sagemaker.Options)) (*sagemaker.ListAlgorithmsOutput, error)
	ListAliases(context.Context, *sagemaker.ListAliasesInput, ...func(*sagemaker.Options)) (*sagemaker.ListAliasesOutput, error)
	ListAppImageConfigs(context.Context, *sagemaker.ListAppImageConfigsInput, ...func(*sagemaker.Options)) (*sagemaker.ListAppImageConfigsOutput, error)
	ListApps(context.Context, *sagemaker.ListAppsInput, ...func(*sagemaker.Options)) (*sagemaker.ListAppsOutput, error)
	ListArtifacts(context.Context, *sagemaker.ListArtifactsInput, ...func(*sagemaker.Options)) (*sagemaker.ListArtifactsOutput, error)
	ListAssociations(context.Context, *sagemaker.ListAssociationsInput, ...func(*sagemaker.Options)) (*sagemaker.ListAssociationsOutput, error)
	ListAutoMLJobs(context.Context, *sagemaker.ListAutoMLJobsInput, ...func(*sagemaker.Options)) (*sagemaker.ListAutoMLJobsOutput, error)
	ListCandidatesForAutoMLJob(context.Context, *sagemaker.ListCandidatesForAutoMLJobInput, ...func(*sagemaker.Options)) (*sagemaker.ListCandidatesForAutoMLJobOutput, error)
	ListCodeRepositories(context.Context, *sagemaker.ListCodeRepositoriesInput, ...func(*sagemaker.Options)) (*sagemaker.ListCodeRepositoriesOutput, error)
	ListCompilationJobs(context.Context, *sagemaker.ListCompilationJobsInput, ...func(*sagemaker.Options)) (*sagemaker.ListCompilationJobsOutput, error)
	ListContexts(context.Context, *sagemaker.ListContextsInput, ...func(*sagemaker.Options)) (*sagemaker.ListContextsOutput, error)
	ListDataQualityJobDefinitions(context.Context, *sagemaker.ListDataQualityJobDefinitionsInput, ...func(*sagemaker.Options)) (*sagemaker.ListDataQualityJobDefinitionsOutput, error)
	ListDeviceFleets(context.Context, *sagemaker.ListDeviceFleetsInput, ...func(*sagemaker.Options)) (*sagemaker.ListDeviceFleetsOutput, error)
	ListDevices(context.Context, *sagemaker.ListDevicesInput, ...func(*sagemaker.Options)) (*sagemaker.ListDevicesOutput, error)
	ListDomains(context.Context, *sagemaker.ListDomainsInput, ...func(*sagemaker.Options)) (*sagemaker.ListDomainsOutput, error)
	ListEdgeDeploymentPlans(context.Context, *sagemaker.ListEdgeDeploymentPlansInput, ...func(*sagemaker.Options)) (*sagemaker.ListEdgeDeploymentPlansOutput, error)
	ListEdgePackagingJobs(context.Context, *sagemaker.ListEdgePackagingJobsInput, ...func(*sagemaker.Options)) (*sagemaker.ListEdgePackagingJobsOutput, error)
	ListEndpointConfigs(context.Context, *sagemaker.ListEndpointConfigsInput, ...func(*sagemaker.Options)) (*sagemaker.ListEndpointConfigsOutput, error)
	ListEndpoints(context.Context, *sagemaker.ListEndpointsInput, ...func(*sagemaker.Options)) (*sagemaker.ListEndpointsOutput, error)
	ListExperiments(context.Context, *sagemaker.ListExperimentsInput, ...func(*sagemaker.Options)) (*sagemaker.ListExperimentsOutput, error)
	ListFeatureGroups(context.Context, *sagemaker.ListFeatureGroupsInput, ...func(*sagemaker.Options)) (*sagemaker.ListFeatureGroupsOutput, error)
	ListFlowDefinitions(context.Context, *sagemaker.ListFlowDefinitionsInput, ...func(*sagemaker.Options)) (*sagemaker.ListFlowDefinitionsOutput, error)
	ListHubContentVersions(context.Context, *sagemaker.ListHubContentVersionsInput, ...func(*sagemaker.Options)) (*sagemaker.ListHubContentVersionsOutput, error)
	ListHubContents(context.Context, *sagemaker.ListHubContentsInput, ...func(*sagemaker.Options)) (*sagemaker.ListHubContentsOutput, error)
	ListHubs(context.Context, *sagemaker.ListHubsInput, ...func(*sagemaker.Options)) (*sagemaker.ListHubsOutput, error)
	ListHumanTaskUis(context.Context, *sagemaker.ListHumanTaskUisInput, ...func(*sagemaker.Options)) (*sagemaker.ListHumanTaskUisOutput, error)
	ListHyperParameterTuningJobs(context.Context, *sagemaker.ListHyperParameterTuningJobsInput, ...func(*sagemaker.Options)) (*sagemaker.ListHyperParameterTuningJobsOutput, error)
	ListImageVersions(context.Context, *sagemaker.ListImageVersionsInput, ...func(*sagemaker.Options)) (*sagemaker.ListImageVersionsOutput, error)
	ListImages(context.Context, *sagemaker.ListImagesInput, ...func(*sagemaker.Options)) (*sagemaker.ListImagesOutput, error)
	ListInferenceExperiments(context.Context, *sagemaker.ListInferenceExperimentsInput, ...func(*sagemaker.Options)) (*sagemaker.ListInferenceExperimentsOutput, error)
	ListInferenceRecommendationsJobSteps(context.Context, *sagemaker.ListInferenceRecommendationsJobStepsInput, ...func(*sagemaker.Options)) (*sagemaker.ListInferenceRecommendationsJobStepsOutput, error)
	ListInferenceRecommendationsJobs(context.Context, *sagemaker.ListInferenceRecommendationsJobsInput, ...func(*sagemaker.Options)) (*sagemaker.ListInferenceRecommendationsJobsOutput, error)
	ListLabelingJobs(context.Context, *sagemaker.ListLabelingJobsInput, ...func(*sagemaker.Options)) (*sagemaker.ListLabelingJobsOutput, error)
	ListLabelingJobsForWorkteam(context.Context, *sagemaker.ListLabelingJobsForWorkteamInput, ...func(*sagemaker.Options)) (*sagemaker.ListLabelingJobsForWorkteamOutput, error)
	ListLineageGroups(context.Context, *sagemaker.ListLineageGroupsInput, ...func(*sagemaker.Options)) (*sagemaker.ListLineageGroupsOutput, error)
	ListModelBiasJobDefinitions(context.Context, *sagemaker.ListModelBiasJobDefinitionsInput, ...func(*sagemaker.Options)) (*sagemaker.ListModelBiasJobDefinitionsOutput, error)
	ListModelCardExportJobs(context.Context, *sagemaker.ListModelCardExportJobsInput, ...func(*sagemaker.Options)) (*sagemaker.ListModelCardExportJobsOutput, error)
	ListModelCardVersions(context.Context, *sagemaker.ListModelCardVersionsInput, ...func(*sagemaker.Options)) (*sagemaker.ListModelCardVersionsOutput, error)
	ListModelCards(context.Context, *sagemaker.ListModelCardsInput, ...func(*sagemaker.Options)) (*sagemaker.ListModelCardsOutput, error)
	ListModelExplainabilityJobDefinitions(context.Context, *sagemaker.ListModelExplainabilityJobDefinitionsInput, ...func(*sagemaker.Options)) (*sagemaker.ListModelExplainabilityJobDefinitionsOutput, error)
	ListModelMetadata(context.Context, *sagemaker.ListModelMetadataInput, ...func(*sagemaker.Options)) (*sagemaker.ListModelMetadataOutput, error)
	ListModelPackageGroups(context.Context, *sagemaker.ListModelPackageGroupsInput, ...func(*sagemaker.Options)) (*sagemaker.ListModelPackageGroupsOutput, error)
	ListModelPackages(context.Context, *sagemaker.ListModelPackagesInput, ...func(*sagemaker.Options)) (*sagemaker.ListModelPackagesOutput, error)
	ListModelQualityJobDefinitions(context.Context, *sagemaker.ListModelQualityJobDefinitionsInput, ...func(*sagemaker.Options)) (*sagemaker.ListModelQualityJobDefinitionsOutput, error)
	ListModels(context.Context, *sagemaker.ListModelsInput, ...func(*sagemaker.Options)) (*sagemaker.ListModelsOutput, error)
	ListMonitoringAlertHistory(context.Context, *sagemaker.ListMonitoringAlertHistoryInput, ...func(*sagemaker.Options)) (*sagemaker.ListMonitoringAlertHistoryOutput, error)
	ListMonitoringAlerts(context.Context, *sagemaker.ListMonitoringAlertsInput, ...func(*sagemaker.Options)) (*sagemaker.ListMonitoringAlertsOutput, error)
	ListMonitoringExecutions(context.Context, *sagemaker.ListMonitoringExecutionsInput, ...func(*sagemaker.Options)) (*sagemaker.ListMonitoringExecutionsOutput, error)
	ListMonitoringSchedules(context.Context, *sagemaker.ListMonitoringSchedulesInput, ...func(*sagemaker.Options)) (*sagemaker.ListMonitoringSchedulesOutput, error)
	ListNotebookInstanceLifecycleConfigs(context.Context, *sagemaker.ListNotebookInstanceLifecycleConfigsInput, ...func(*sagemaker.Options)) (*sagemaker.ListNotebookInstanceLifecycleConfigsOutput, error)
	ListNotebookInstances(context.Context, *sagemaker.ListNotebookInstancesInput, ...func(*sagemaker.Options)) (*sagemaker.ListNotebookInstancesOutput, error)
	ListPipelineExecutionSteps(context.Context, *sagemaker.ListPipelineExecutionStepsInput, ...func(*sagemaker.Options)) (*sagemaker.ListPipelineExecutionStepsOutput, error)
	ListPipelineExecutions(context.Context, *sagemaker.ListPipelineExecutionsInput, ...func(*sagemaker.Options)) (*sagemaker.ListPipelineExecutionsOutput, error)
	ListPipelineParametersForExecution(context.Context, *sagemaker.ListPipelineParametersForExecutionInput, ...func(*sagemaker.Options)) (*sagemaker.ListPipelineParametersForExecutionOutput, error)
	ListPipelines(context.Context, *sagemaker.ListPipelinesInput, ...func(*sagemaker.Options)) (*sagemaker.ListPipelinesOutput, error)
	ListProcessingJobs(context.Context, *sagemaker.ListProcessingJobsInput, ...func(*sagemaker.Options)) (*sagemaker.ListProcessingJobsOutput, error)
	ListProjects(context.Context, *sagemaker.ListProjectsInput, ...func(*sagemaker.Options)) (*sagemaker.ListProjectsOutput, error)
	ListSpaces(context.Context, *sagemaker.ListSpacesInput, ...func(*sagemaker.Options)) (*sagemaker.ListSpacesOutput, error)
	ListStageDevices(context.Context, *sagemaker.ListStageDevicesInput, ...func(*sagemaker.Options)) (*sagemaker.ListStageDevicesOutput, error)
	ListStudioLifecycleConfigs(context.Context, *sagemaker.ListStudioLifecycleConfigsInput, ...func(*sagemaker.Options)) (*sagemaker.ListStudioLifecycleConfigsOutput, error)
	ListSubscribedWorkteams(context.Context, *sagemaker.ListSubscribedWorkteamsInput, ...func(*sagemaker.Options)) (*sagemaker.ListSubscribedWorkteamsOutput, error)
	ListTags(context.Context, *sagemaker.ListTagsInput, ...func(*sagemaker.Options)) (*sagemaker.ListTagsOutput, error)
	ListTrainingJobs(context.Context, *sagemaker.ListTrainingJobsInput, ...func(*sagemaker.Options)) (*sagemaker.ListTrainingJobsOutput, error)
	ListTrainingJobsForHyperParameterTuningJob(context.Context, *sagemaker.ListTrainingJobsForHyperParameterTuningJobInput, ...func(*sagemaker.Options)) (*sagemaker.ListTrainingJobsForHyperParameterTuningJobOutput, error)
	ListTransformJobs(context.Context, *sagemaker.ListTransformJobsInput, ...func(*sagemaker.Options)) (*sagemaker.ListTransformJobsOutput, error)
	ListTrialComponents(context.Context, *sagemaker.ListTrialComponentsInput, ...func(*sagemaker.Options)) (*sagemaker.ListTrialComponentsOutput, error)
	ListTrials(context.Context, *sagemaker.ListTrialsInput, ...func(*sagemaker.Options)) (*sagemaker.ListTrialsOutput, error)
	ListUserProfiles(context.Context, *sagemaker.ListUserProfilesInput, ...func(*sagemaker.Options)) (*sagemaker.ListUserProfilesOutput, error)
	ListWorkforces(context.Context, *sagemaker.ListWorkforcesInput, ...func(*sagemaker.Options)) (*sagemaker.ListWorkforcesOutput, error)
	ListWorkteams(context.Context, *sagemaker.ListWorkteamsInput, ...func(*sagemaker.Options)) (*sagemaker.ListWorkteamsOutput, error)
	Search(context.Context, *sagemaker.SearchInput, ...func(*sagemaker.Options)) (*sagemaker.SearchOutput, error)
}
