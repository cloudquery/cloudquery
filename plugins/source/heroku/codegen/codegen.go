package codegen

import (
	"github.com/cloudquery/plugin-sdk/schema"
	heroku "github.com/heroku/heroku-go/v5"
	"reflect"

	"github.com/cloudquery/plugin-sdk/codegen"
)

type Resource struct {
	// DefaultColumns columns that will be appended to the main table
	DefaultColumns []codegen.ColumnDefinition
	// Table is the table definition that will be used to generate the cloudquery table
	Table *codegen.TableDefinition
	// HerokuStruct that will be used to generate the cloudquery table
	HerokuStruct interface{}
	// HerokuStructName is the name of the HerokuStruct because it can't be inferred by reflection
	HerokuStructName string
	// Template is the template to use to generate the resource
	Template string
	// SkipFields fields in go struct to skip when generating the table from the go struct
	SkipFields []string
	// CreateTableOptions options to use to create the main table
	CreateTableOptions schema.TableCreationOptions
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
	//{
	//	HerokuStruct: &heroku.AddOnConfig{},
	//},
	//{
	//	HerokuStruct: &heroku.AddOnRegionCapability{},
	//},
	//{
	//	HerokuStruct: &heroku.AddOnService{},
	//},
	//{
	//	HerokuStruct: &heroku.AddOnWebhook{},
	//},
	//{
	//	HerokuStruct: &heroku.AddOnWebhookDelivery{},
	//},
	//{
	//	HerokuStruct: &heroku.AddOnWebhookEvent{},
	//},
	//{
	//	HerokuStruct: &heroku.App{},
	//},
	//{
	//	HerokuStruct: &heroku.AppFeature{},
	//},
	//{
	//	HerokuStruct: &heroku.AppTransfer{},
	//},
	//{
	//	HerokuStruct: &heroku.AppWebhook{},
	//},
	//{
	//	HerokuStruct: &heroku.AppWebhookDelivery{},
	//},
	//{
	//	HerokuStruct: &heroku.AppWebhookEvent{},
	//},
	//{
	//	HerokuStruct: &heroku.Archive{},
	//},
	//{
	//	HerokuStruct: &heroku.AuditTrailEvent{},
	//},
	//{
	//	HerokuStruct: &heroku.Build{},
	//},
	//{
	//	HerokuStruct: &heroku.BuildpackInstallation{},
	//},
	//{
	//	HerokuStruct: &heroku.Collaborator{},
	//},
	//{
	//	HerokuStruct: &heroku.Credit{},
	//},
	//{
	//	HerokuStruct: &heroku.Domain{},
	//},
	//{
	//	HerokuStruct: &heroku.Dyno{},
	//},
	//{
	//	HerokuStruct: &heroku.DynoSize{},
	//},
	//{
	//	HerokuStruct: &heroku.EnterpriseAccount{},
	//},
	//{
	//	HerokuStruct: &heroku.EnterpriseAccountMember{},
	//},
	//{
	//	HerokuStruct: &heroku.Formation{},
	//},
	//{
	//	HerokuStruct: &heroku.InboundRuleset{},
	//},
	//{
	//	HerokuStruct: &heroku.Invoice{},
	//},
	//{
	//	HerokuStruct: &heroku.Key{},
	//},
	//{
	//	HerokuStruct: &heroku.LogDrain{},
	//},
	//{
	//	HerokuStruct: &heroku.OAuthAuthorization{},
	//},
	//{
	//	HerokuStruct: &heroku.OAuthClient{},
	//},
	//{
	//	HerokuStruct: &heroku.OutboundRuleset{},
	//},
	//{
	//	HerokuStruct: &heroku.Peering{},
	//},
	//{
	//	HerokuStruct: &heroku.PermissionEntity{},
	//},
	//{
	//	HerokuStruct: &heroku.Pipeline{},
	//},
	//{
	//	HerokuStruct: &heroku.PipelineBuild{},
	//},
	//{
	//	HerokuStruct: &heroku.PipelineCoupling{},
	//},
	//{
	//	HerokuStruct: &heroku.PipelineDeployment{},
	//},
	//{
	//	HerokuStruct: &heroku.PipelinePromotionTarget{},
	//},
	//{
	//	HerokuStruct: &heroku.PipelineRelease{},
	//},
	//{
	//	HerokuStruct: &heroku.Region{},
	//},
	//{
	//	HerokuStruct: &heroku.Release{},
	//},
	//{
	//	HerokuStruct: &heroku.ReviewApp{},
	//},
	//{
	//	HerokuStruct: &heroku.SniEndpoint{},
	//},
	//{
	//	HerokuStruct: &heroku.Space{},
	//},
	//{
	//	HerokuStruct: &heroku.SpaceAppAccess{},
	//},
	//{
	//	HerokuStruct: &heroku.SSLEndpoint{},
	//},
	{
		HerokuStruct: &heroku.Stack{},
	},
	{
		HerokuStruct: &heroku.Team{},
	},
	//{
	//	HerokuStruct: &heroku.TeamAppCollaborator{},
	//},
	//{
	//	HerokuStruct: &heroku.TeamAppPermission{},
	//},
	//{
	//	HerokuStruct: &heroku.TeamFeature{},
	//},
	//{
	//	HerokuStruct: &heroku.TeamInvitation{},
	//},
	//{
	//	HerokuStruct: &heroku.TeamInvoice{},
	//},
	//{
	//	HerokuStruct: &heroku.TeamMember{},
	//},
	//{
	//	HerokuStruct: &heroku.TeamPreferences{},
	//},
	//{
	//	HerokuStruct: &heroku.TeamSpace{},
	//},
	//{
	//	HerokuStruct: &heroku.TestCase{},
	//},
	//{
	//	HerokuStruct: &heroku.TestNode{},
	//},
	//{
	//	HerokuStruct: &heroku.TestRun{},
	//},
	//{
	//	HerokuStruct: &heroku.UserPreferences{},
	//},
	//{
	//	HerokuStruct: &heroku.VPNConnection{},
	//},
}

func All() []Resource {
	resources := listResources
	// add all shared properties
	for i, r := range resources {
		r.Template = "resource_list"
		r.HerokuStructName = reflect.TypeOf(r.HerokuStruct).Elem().Name()
		r.DefaultColumns = []codegen.ColumnDefinition{}
		r.SkipFields = []string{}
		if resources[i].CreateTableOptions.PrimaryKeys == nil {
			resources[i].CreateTableOptions.PrimaryKeys = []string{"id"}
		}
		resources[i] = r
	}
	return resources
}
