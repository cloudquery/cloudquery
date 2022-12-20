package recipes

import (
	"reflect"

	"github.com/cloudquery/plugin-sdk/codegen"
	heroku "github.com/heroku/heroku-go/v5"
)

type Resource struct {
	// Table is the table definition that will be used to generate the cloudquery table
	Table *codegen.TableDefinition
	// TableName can be used to override the default generated table name
	TableName string
	// HerokuStruct that will be used to generate the cloudquery table
	HerokuStruct any
	// HerokuStructName is the name of the HerokuStruct because it can't be inferred by reflection
	HerokuStructName string
	// HerokuPrimaryStructName is the primary struct that will be listed first, then used as an ID in additional list calls
	// (applies only to relational resources)
	HerokuPrimaryStruct any
	// HerokuPrimaryStructName is the name of the primary struct
	// (applies only to relational resources)
	HerokuPrimaryStructName string
	// SkipListParams indicates whether to skip passing nil as params argument, or not
	SkipListParams bool
	// Template is the template to use to generate the resource
	Template string
	// SkipFields fields in go struct to skip when generating the table from the go struct
	SkipFields []string
	PrimaryKey string
}

var listResources = []Resource{
	{
		HerokuStruct: &heroku.AccountFeature{},
	},
	{
		HerokuStruct: &heroku.AddOn{},
	},
	{
		HerokuStruct: &heroku.AddOnAttachment{},
	},
	{
		HerokuStruct:        &heroku.AddOnConfig{},
		HerokuPrimaryStruct: &heroku.AddOn{},
		Template:            "relational_resource_list",
	},
	{
		HerokuStruct: &heroku.AddOnRegionCapability{},
	},
	{
		HerokuStruct: &heroku.AddOnService{},
	},
	{
		HerokuStruct:        &heroku.AddOnWebhook{},
		HerokuPrimaryStruct: &heroku.AddOn{},
		Template:            "relational_resource_list",
	},
	{
		TableName:           "heroku_add_on_webhook_deliveries",
		HerokuStruct:        &heroku.AddOnWebhookDeliveryInfoResult{},
		HerokuStructName:    "AddOnWebhookDelivery",
		HerokuPrimaryStruct: &heroku.AddOn{},
		Template:            "relational_resource_list",
	},
	{
		TableName:           "heroku_add_on_webhook_events",
		HerokuStruct:        &heroku.AddOnWebhookEventInfoResult{},
		HerokuStructName:    "AddOnWebhookEvent",
		HerokuPrimaryStruct: &heroku.AddOn{},
		Template:            "relational_resource_list",
	},
	{
		HerokuStruct: &heroku.App{},
	},
	{
		HerokuStruct:        &heroku.AppFeature{},
		HerokuPrimaryStruct: &heroku.App{},
		Template:            "relational_resource_list",
	},
	{
		HerokuStruct: &heroku.AppTransfer{},
	},
	{
		TableName:           "heroku_app_webhooks",
		HerokuStruct:        &heroku.AppWebhookInfoResult{},
		HerokuStructName:    "AppWebhook",
		HerokuPrimaryStruct: &heroku.App{},
		Template:            "relational_resource_list",
	},
	{
		HerokuStruct:        &heroku.AppWebhookDelivery{},
		HerokuPrimaryStruct: &heroku.App{},
		Template:            "relational_resource_list",
	},
	{
		HerokuStruct:        &heroku.AppWebhookEvent{},
		HerokuPrimaryStruct: &heroku.App{},
		Template:            "relational_resource_list",
	},
	{
		HerokuStruct:        &heroku.Build{},
		HerokuPrimaryStruct: &heroku.App{},
		Template:            "relational_resource_list",
	},
	{
		HerokuStruct:        &heroku.BuildpackInstallation{},
		HerokuPrimaryStruct: &heroku.App{},
		Template:            "relational_resource_list",
	},
	{
		HerokuStruct:        &heroku.Collaborator{},
		HerokuPrimaryStruct: &heroku.App{},
		Template:            "relational_resource_list",
	},
	{
		HerokuStruct: &heroku.Credit{},
	},
	{
		HerokuStruct:        &heroku.Domain{},
		HerokuPrimaryStruct: &heroku.App{},
		Template:            "relational_resource_list",
	},
	{
		HerokuStruct:        &heroku.Dyno{},
		HerokuPrimaryStruct: &heroku.App{},
		Template:            "relational_resource_list",
	},
	{
		HerokuStruct: &heroku.DynoSize{},
	},
	{
		HerokuStruct: &heroku.EnterpriseAccount{},
	},
	{
		HerokuStruct:        &heroku.EnterpriseAccountMember{},
		HerokuPrimaryStruct: &heroku.EnterpriseAccount{},
		Template:            "relational_resource_list",
	},
	{
		HerokuStruct:        &heroku.Formation{},
		HerokuPrimaryStruct: &heroku.App{},
		Template:            "relational_resource_list",
	},
	{
		HerokuStruct:        &heroku.InboundRuleset{},
		HerokuPrimaryStruct: &heroku.Space{},
		Template:            "relational_resource_list",
	},
	{
		HerokuStruct: &heroku.Invoice{},
	},
	{
		HerokuStruct: &heroku.Key{},
	},
	{
		HerokuStruct:        &heroku.LogDrain{},
		HerokuPrimaryStruct: &heroku.App{},
		Template:            "relational_resource_list",
	},
	{
		TableName:    "heroku_oauth_authorizations",
		HerokuStruct: &heroku.OAuthAuthorization{},
	},
	{
		TableName:    "heroku_oauth_clients",
		HerokuStruct: &heroku.OAuthClient{},
	},
	{
		HerokuStruct:        &heroku.OutboundRuleset{},
		HerokuPrimaryStruct: &heroku.Space{},
		Template:            "relational_resource_list",
	},
	{
		HerokuStruct:        &heroku.Peering{},
		HerokuPrimaryStruct: &heroku.Space{},
		Template:            "relational_resource_list",
	},
	{
		HerokuStruct:        &heroku.PermissionEntity{},
		HerokuPrimaryStruct: &heroku.Team{},
		Template:            "relational_resource_list",
	},
	{
		HerokuStruct: &heroku.Pipeline{},
	},
	{
		TableName:           "heroku_pipeline_builds",
		HerokuStruct:        &heroku.PipelineBuildListResult{{}}[0],
		HerokuStructName:    "PipelineBuild",
		HerokuPrimaryStruct: &heroku.Pipeline{},
		Template:            "relational_resource_list",
	},
	{
		HerokuStruct: &heroku.PipelineCoupling{},
	},
	{
		TableName:           "heroku_pipeline_deployments",
		HerokuStruct:        &heroku.PipelineDeploymentListResult{{}}[0],
		HerokuStructName:    "PipelineDeployment",
		HerokuPrimaryStruct: &heroku.Pipeline{},
		Template:            "relational_resource_list",
	},
	{
		TableName:           "heroku_pipeline_releases",
		HerokuStruct:        &heroku.PipelineReleaseListResult{{}}[0],
		HerokuStructName:    "PipelineRelease",
		HerokuPrimaryStruct: &heroku.Pipeline{},
		Template:            "relational_resource_list",
	},
	{
		HerokuStruct: &heroku.Region{},
	},
	{
		HerokuStruct:        &heroku.Release{},
		HerokuPrimaryStruct: &heroku.App{},
		Template:            "relational_resource_list",
	},
	{
		HerokuStruct:        &heroku.ReviewApp{},
		HerokuPrimaryStruct: &heroku.Pipeline{},
		Template:            "relational_resource_list",
	},
	{
		HerokuStruct: &heroku.Space{},
	},
	{
		HerokuStruct:        &heroku.SpaceAppAccess{},
		HerokuPrimaryStruct: &heroku.Space{},
		Template:            "relational_resource_list",
	},
	{
		HerokuStruct: &heroku.Stack{},
	},
	{
		HerokuStruct: &heroku.Team{},
	},
	{
		HerokuStruct: &heroku.TeamAppPermission{},
	},
	{
		HerokuStruct:        &heroku.TeamFeature{},
		HerokuPrimaryStruct: &heroku.Team{},
		Template:            "relational_resource_list",
	},
	{
		HerokuStruct:        &heroku.TeamInvitation{},
		HerokuPrimaryStruct: &heroku.Team{},
		Template:            "relational_resource_list",
	},
	{
		HerokuStruct:        &heroku.TeamInvoice{},
		HerokuPrimaryStruct: &heroku.Team{},
		Template:            "relational_resource_list",
	},
	{
		HerokuStruct:        &heroku.TeamMember{},
		HerokuPrimaryStruct: &heroku.Team{},
		Template:            "relational_resource_list",
	},
	{
		TableName:           "heroku_team_spaces",
		HerokuStruct:        &heroku.TeamSpaceListResult{{}}[0],
		HerokuStructName:    "TeamSpace",
		HerokuPrimaryStruct: &heroku.Team{},
		Template:            "relational_resource_list",
	},
	{
		HerokuStruct:        &heroku.VPNConnection{},
		HerokuPrimaryStruct: &heroku.Space{},
		Template:            "relational_resource_list",
	},

	// TODO: Add support for Archive
	//{
	//	HerokuStruct:        &heroku.Archive{},
	//	HerokuPrimaryStruct: &heroku.EnterpriseAccount{},
	//	Template:            "relational_resource_list",
	//},
	// TODO: Add support for AuditTrailEvent.
	//       Its list call returns only a single item, which is probably
	//       a bug in the Heroku SDK.
	//{
	//	HerokuStruct:        &heroku.AuditTrailEvent{},
	//	HerokuPrimaryStruct: &heroku.EnterpriseAccount{},
	//	Template:            "relational_resource_list",
	//},
	// TODO: Add support for PipelinePromotion
	// Note: PipelinePromotion doesn't have a corresponding List call
	//{
	//	HerokuStruct:        &heroku.PipelinePromotionTarget{},
	//	HerokuPrimaryStruct: &heroku.PipelinePromotion{},
	//	Template:            "relational_resource_list",
	//},
	// TODO: Add support for SniEndpoint
	//{
	//	HerokuStruct:        &heroku.SniEndpoint{},
	//	HerokuPrimaryStruct: &heroku.App{},
	//	Template:            "relational_resource_list",
	//},
	// TODO: Add support for SSLEndpoint
	//{
	//	HerokuStruct:        &heroku.SSLEndpoint{},
	//	HerokuPrimaryStruct: &heroku.App{},
	//	Template:            "relational_resource_list",
	//},
	// TODO: Add support for TeamApp
	// NOTE: TeamApp only has TeamAppListByTeam
	//{
	//	HerokuStruct: &heroku.TeamApp{},
	//},
	// TODO: Add support for TeamAppCollaborator
	//{
	//	HerokuStruct: &heroku.TeamAppCollaborator{},
	//	HerokuPrimaryStruct: &heroku.TeamApp{},
	//	Template:            "relational_resource_list",
	//},
	// TODO: Add support for TeamPreferences
	//{
	//	HerokuStruct:        &heroku.TeamPreferences{},
	//	HerokuPrimaryStruct: &heroku.Team{},
	//	Template:            "relational_resource_list",
	//	SkipListParams:      true,
	//},
	// TODO: Add support for TestCase
	//{
	//	HerokuStruct:        &heroku.TestCase{},
	//	HerokuPrimaryStruct: &heroku.TestRun{},
	//	Template:            "relational_resource_list",
	//},
	// TODO: Add support for TestNode
	//{
	//	HerokuStruct:        &heroku.TestNode{},
	//	HerokuPrimaryStruct: &heroku.TestRun{},
	//	Template:            "relational_resource_list",
	//},
	// TODO: Add support for TestRun
	//{
	//	HerokuStruct:        &heroku.TestRun{},
	//	HerokuPrimaryStruct: &heroku.Pipeline{},
	//	Template:            "relational_resource_list",
	//},
	// TODO: Add support for UserPreferences
	// Note: no API function exists to list accounts, so preferences would
	// only be for current user
	//{
	//	HerokuStruct:        &heroku.UserPreferences{},
	//	HerokuPrimaryStruct: &heroku.Account{},
	//	Template:            "relational_resource_list",
	//},
}

func All() []Resource {
	resources := listResources
	// add all shared and default properties
	for i, r := range resources {
		if r.Template == "" {
			r.Template = "resource_list"
		}
		if len(r.HerokuStructName) == 0 {
			r.HerokuStructName = reflect.TypeOf(r.HerokuStruct).Elem().Name()
		}
		if r.HerokuPrimaryStruct != nil {
			r.HerokuPrimaryStructName = reflect.TypeOf(r.HerokuPrimaryStruct).Elem().Name()
		}
		r.SkipFields = []string{}
		if r.PrimaryKey == "" {
			r.PrimaryKey = "id"
		}
		resources[i] = r
	}
	return resources
}
