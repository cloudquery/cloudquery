package recipes

import (
	"reflect"

	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
	heroku "github.com/heroku/heroku-go/v5"
)

type Resource struct {
	// DefaultColumns columns that will be appended to the main table
	DefaultColumns []codegen.ColumnDefinition
	// Table is the table definition that will be used to generate the cloudquery table
	Table *codegen.TableDefinition
	// TableName can be used to override the default generated table name
	TableName string
	// HerokuStruct that will be used to generate the table name
	HerokuStruct interface{}
	// HerokuListResultStruct will be used to generate the cloudquery table. If empty, HerokuStruct will be used.
	// HerokuStruct will still be used for getting the name.
	HerokuListResultStruct interface{}
	// HerokuStructName is the name of the HerokuStruct because it can't be inferred by reflection
	HerokuStructName string
	// HerokuPrimaryStructName is the primary struct that will be listed first, then used as an ID in additional list calls
	// (applies only to relational resources)
	HerokuPrimaryStruct interface{}
	// HerokuPrimaryStructName is the name of the primary struct
	// (applies only to relational resources)
	HerokuPrimaryStructName string
	// SkipListParams indicates whether to skip passing nil as params argument, or not
	SkipListParams bool
	// Template is the template to use to generate the resource
	Template string
	// SkipFields fields in go struct to skip when generating the table from the go struct
	SkipFields []string
	// CreateTableOptions options to use to create the main table
	CreateTableOptions schema.TableCreationOptions
}

var listResources = []Resource{
	{
		HerokuStruct:           &heroku.AccountFeature{},
		HerokuListResultStruct: &heroku.AccountFeatureListResult{{}},
	},
	{
		HerokuStruct:           &heroku.AddOn{},
		HerokuListResultStruct: &heroku.AddOnListResult{{}},
	},
	{
		HerokuStruct:           &heroku.AddOnAttachment{},
		HerokuListResultStruct: &heroku.AddOnAttachmentListResult{{}},
	},
	{
		HerokuStruct:           &heroku.AddOnConfig{},
		HerokuListResultStruct: &heroku.AddOnConfigListResult{{}},
		HerokuPrimaryStruct:    &heroku.AddOn{},
		Template:               "relational_resource_list",
	},
	{
		HerokuStruct:           &heroku.AddOnRegionCapability{},
		HerokuListResultStruct: &heroku.AddOnRegionCapabilityListResult{{}},
	},
	{
		HerokuStruct:           &heroku.AddOnService{},
		HerokuListResultStruct: &heroku.AddOnServiceListResult{{}},
	},
	{
		HerokuStruct:           &heroku.AddOnWebhook{},
		HerokuListResultStruct: &heroku.AddOnWebhookListResult{{}},
		HerokuPrimaryStruct:    &heroku.AddOn{},
		Template:               "relational_resource_list",
	},
	{
		HerokuStruct:           &heroku.AddOnWebhookDelivery{},
		HerokuListResultStruct: &heroku.AddOnWebhookDeliveryListResult{{}},
		HerokuPrimaryStruct:    &heroku.AddOn{},
		Template:               "relational_resource_list",
	},
	{
		HerokuStruct:           &heroku.AddOnWebhookEvent{},
		HerokuListResultStruct: &heroku.AddOnWebhookEventListResult{{}},
		HerokuPrimaryStruct:    &heroku.AddOn{},
		Template:               "relational_resource_list",
	},
	{
		HerokuStruct:           &heroku.App{},
		HerokuListResultStruct: &heroku.AppListResult{{}},
	},
	{
		HerokuStruct:           &heroku.AppFeature{},
		HerokuListResultStruct: &heroku.AppFeatureListResult{{}},
		HerokuPrimaryStruct:    &heroku.App{},
		Template:               "relational_resource_list",
	},
	{
		HerokuStruct:           &heroku.AppTransfer{},
		HerokuListResultStruct: &heroku.AppTransferListResult{{}},
	},
	{
		HerokuStruct:           &heroku.AppWebhook{},
		HerokuListResultStruct: &heroku.AppWebhookListResult{{}},
		HerokuPrimaryStruct:    &heroku.App{},
		Template:               "relational_resource_list",
	},
	{
		HerokuStruct:           &heroku.AppWebhookDelivery{},
		HerokuListResultStruct: &heroku.AppWebhookDeliveryListResult{{}},
		HerokuPrimaryStruct:    &heroku.App{},
		Template:               "relational_resource_list",
	},
	{
		HerokuStruct:           &heroku.AppWebhookEvent{},
		HerokuListResultStruct: &heroku.AppWebhookEventListResult{{}},
		HerokuPrimaryStruct:    &heroku.App{},
		Template:               "relational_resource_list",
	},
	{
		HerokuStruct:           &heroku.Build{},
		HerokuListResultStruct: &heroku.BuildListResult{{}},
		HerokuPrimaryStruct:    &heroku.App{},
		Template:               "relational_resource_list",
	},
	{
		HerokuStruct:           &heroku.BuildpackInstallation{},
		HerokuListResultStruct: &heroku.BuildpackInstallationListResult{{}},
		HerokuPrimaryStruct:    &heroku.App{},
		Template:               "relational_resource_list",
	},
	{
		HerokuStruct:           &heroku.Collaborator{},
		HerokuListResultStruct: &heroku.CollaboratorListResult{{}},
		HerokuPrimaryStruct:    &heroku.App{},
		Template:               "relational_resource_list",
	},
	{
		HerokuStruct:           &heroku.Credit{},
		HerokuListResultStruct: &heroku.CreditListResult{{}},
	},
	{
		HerokuStruct:           &heroku.Domain{},
		HerokuListResultStruct: &heroku.DomainListResult{{}},
		HerokuPrimaryStruct:    &heroku.App{},
		Template:               "relational_resource_list",
	},
	{
		HerokuStruct:           &heroku.Dyno{},
		HerokuListResultStruct: &heroku.DynoListResult{{}},
		HerokuPrimaryStruct:    &heroku.App{},
		Template:               "relational_resource_list",
	},
	{
		HerokuStruct:           &heroku.DynoSize{},
		HerokuListResultStruct: &heroku.DynoSizeListResult{{}},
	},
	{
		HerokuStruct:           &heroku.EnterpriseAccount{},
		HerokuListResultStruct: &heroku.EnterpriseAccountListResult{{}},
	},
	{
		HerokuStruct:           &heroku.EnterpriseAccountMember{},
		HerokuListResultStruct: &heroku.EnterpriseAccountMemberListResult{{}},
		HerokuPrimaryStruct:    &heroku.EnterpriseAccount{},
		Template:               "relational_resource_list",
	},
	{
		HerokuStruct:           &heroku.Formation{},
		HerokuListResultStruct: &heroku.FormationListResult{{}},
		HerokuPrimaryStruct:    &heroku.App{},
		Template:               "relational_resource_list",
	},
	{
		HerokuStruct:           &heroku.InboundRuleset{},
		HerokuListResultStruct: &heroku.InboundRulesetListResult{{}},
		HerokuPrimaryStruct:    &heroku.Space{},
		Template:               "relational_resource_list",
	},
	{
		HerokuStruct:           &heroku.Invoice{},
		HerokuListResultStruct: &heroku.InvoiceListResult{{}},
	},
	{
		HerokuStruct:           &heroku.Key{},
		HerokuListResultStruct: &heroku.KeyListResult{{}},
	},
	{
		HerokuStruct:           &heroku.LogDrain{},
		HerokuListResultStruct: &heroku.LogDrainListResult{{}},
		HerokuPrimaryStruct:    &heroku.App{},
		Template:               "relational_resource_list",
	},
	{
		TableName:              "heroku_oauth_authorizations",
		HerokuStruct:           &heroku.OAuthAuthorization{},
		HerokuListResultStruct: &heroku.OAuthAuthorizationListResult{{}},
	},
	{
		TableName:              "heroku_oauth_clients",
		HerokuStruct:           &heroku.OAuthClient{},
		HerokuListResultStruct: &heroku.OAuthClientListResult{{}},
	},
	{
		HerokuStruct:           &heroku.OutboundRuleset{},
		HerokuListResultStruct: &heroku.OutboundRulesetListResult{{}},
		HerokuPrimaryStruct:    &heroku.Space{},
		Template:               "relational_resource_list",
	},
	{
		HerokuStruct:           &heroku.Peering{},
		HerokuListResultStruct: &heroku.PeeringListResult{{}},
		HerokuPrimaryStruct:    &heroku.Space{},
		Template:               "relational_resource_list",
	},
	{
		HerokuStruct:           &heroku.PermissionEntity{},
		HerokuListResultStruct: &heroku.PermissionEntityListResult{{}},
		HerokuPrimaryStruct:    &heroku.Team{},
		Template:               "relational_resource_list",
	},
	{
		HerokuStruct:           &heroku.Pipeline{},
		HerokuListResultStruct: &heroku.PipelineListResult{{}},
	},
	{
		HerokuStruct:           &heroku.PipelineBuild{},
		HerokuListResultStruct: &heroku.PipelineBuildListResult{{}},
		HerokuPrimaryStruct:    &heroku.Pipeline{},
		Template:               "relational_resource_list",
	},
	{
		HerokuStruct:           &heroku.PipelineCoupling{},
		HerokuListResultStruct: &heroku.PipelineCouplingListResult{{}},
	},
	{
		HerokuStruct:           &heroku.PipelineDeployment{},
		HerokuListResultStruct: &heroku.PipelineDeploymentListResult{{}},
		HerokuPrimaryStruct:    &heroku.Pipeline{},
		Template:               "relational_resource_list",
	},
	{
		HerokuStruct:           &heroku.PipelineRelease{},
		HerokuListResultStruct: &heroku.PipelineReleaseListResult{{}},
		HerokuPrimaryStruct:    &heroku.Pipeline{},
		Template:               "relational_resource_list",
	},
	{
		HerokuStruct:           &heroku.Region{},
		HerokuListResultStruct: &heroku.RegionListResult{{}},
	},
	{
		HerokuStruct:           &heroku.Release{},
		HerokuListResultStruct: &heroku.ReleaseListResult{{}},
		HerokuPrimaryStruct:    &heroku.App{},
		Template:               "relational_resource_list",
	},
	{
		HerokuStruct:           &heroku.ReviewApp{},
		HerokuListResultStruct: &heroku.ReviewAppListResult{{}},
		HerokuPrimaryStruct:    &heroku.Pipeline{},
		Template:               "relational_resource_list",
	},
	{
		HerokuStruct:           &heroku.Space{},
		HerokuListResultStruct: &heroku.SpaceListResult{{}},
	},
	{
		HerokuStruct:           &heroku.SpaceAppAccess{},
		HerokuListResultStruct: &heroku.SpaceAppAccessListResult{{}},
		HerokuPrimaryStruct:    &heroku.Space{},
		Template:               "relational_resource_list",
	},
	{
		HerokuStruct:           &heroku.Stack{},
		HerokuListResultStruct: &heroku.StackListResult{{}},
	},
	{
		HerokuStruct:           &heroku.Team{},
		HerokuListResultStruct: &heroku.TeamListResult{{}},
	},
	{
		HerokuStruct:           &heroku.TeamAppPermission{},
		HerokuListResultStruct: &heroku.TeamAppPermissionListResult{{}},
	},
	{
		HerokuStruct:           &heroku.TeamFeature{},
		HerokuListResultStruct: &heroku.TeamFeatureListResult{{}},
		HerokuPrimaryStruct:    &heroku.Team{},
		Template:               "relational_resource_list",
	},
	{
		HerokuStruct:           &heroku.TeamInvitation{},
		HerokuListResultStruct: &heroku.TeamInvitationListResult{{}},
		HerokuPrimaryStruct:    &heroku.Team{},
		Template:               "relational_resource_list",
	},
	{
		HerokuStruct:           &heroku.TeamInvoice{},
		HerokuListResultStruct: &heroku.TeamInvoiceListResult{{}},
		HerokuPrimaryStruct:    &heroku.Team{},
		Template:               "relational_resource_list",
	},
	{
		HerokuStruct:           &heroku.TeamMember{},
		HerokuListResultStruct: &heroku.TeamMemberListResult{{}},
		HerokuPrimaryStruct:    &heroku.Team{},
		Template:               "relational_resource_list",
	},
	{
		HerokuStruct:           &heroku.TeamSpace{},
		HerokuListResultStruct: &heroku.TeamSpaceListResult{{}},
		HerokuPrimaryStruct:    &heroku.Team{},
		Template:               "relational_resource_list",
	},
	{
		HerokuStruct:           &heroku.VPNConnection{},
		HerokuListResultStruct: &heroku.VPNConnectionListResult{{}},
		HerokuPrimaryStruct:    &heroku.Space{},
		Template:               "relational_resource_list",
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
		r.HerokuStructName = reflect.TypeOf(r.HerokuStruct).Elem().Name()
		if r.HerokuPrimaryStruct != nil {
			r.HerokuPrimaryStructName = reflect.TypeOf(r.HerokuPrimaryStruct).Elem().Name()
		}
		r.DefaultColumns = []codegen.ColumnDefinition{}
		r.SkipFields = []string{}
		if resources[i].CreateTableOptions.PrimaryKeys == nil {
			resources[i].CreateTableOptions.PrimaryKeys = []string{"id"}
		}
		resources[i] = r
	}
	return resources
}
