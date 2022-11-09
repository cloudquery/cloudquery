package recipes

import (
	"reflect"
	"time"

	"github.com/cloudquery/plugin-sdk/codegen"
	heroku "github.com/heroku/heroku-go/v5"
)

type Resource struct {
	// Table is the table definition that will be used to generate the cloudquery table
	Table *codegen.TableDefinition
	// TableName can be used to override the default generated table name
	TableName string
	// HerokuStruct that will be used to generate the cloudquery table
	HerokuStruct interface{}
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
	PrimaryKey string
}

type AddOnWebhookDelivery struct {
	heroku.AddOnWebhookDeliveryInfoResult
}

type AddOnWebhookEvent struct {
	heroku.AddOnWebhookEventInfoResult
}

type AppWebhook struct {
	heroku.AppWebhookInfoResult
}

// Element of PipelineBuildListResult
type PipelineBuild struct {
	App struct {
		ID string `json:"id" url:"id,key"` // unique identifier of app
	} `json:"app" url:"app,key"` // app that the build belongs to
	Buildpacks []struct {
		Name string `json:"name" url:"name,key"` // Buildpack Registry name of the buildpack for the app
		URL  string `json:"url" url:"url,key"`   // the URL of the buildpack for the app
	} `json:"buildpacks" url:"buildpacks,key"` // buildpacks executed for this build, in order
	CreatedAt       time.Time `json:"created_at" url:"created_at,key"`               // when build was created
	ID              string    `json:"id" url:"id,key"`                               // unique identifier of build
	OutputStreamURL string    `json:"output_stream_url" url:"output_stream_url,key"` // Build process output will be available from this URL as a stream. The
	// stream is available as either `text/plain` or `text/event-stream`.
	// Clients should be prepared to handle disconnects and can resume the
	// stream by sending a `Range` header (for `text/plain`) or a
	// `Last-Event-Id` header (for `text/event-stream`).
	Release *struct {
		ID string `json:"id" url:"id,key"` // unique identifier of release
	} `json:"release" url:"release,key"` // release resulting from the build
	Slug *struct {
		ID string `json:"id" url:"id,key"` // unique identifier of slug
	} `json:"slug" url:"slug,key"` // slug created by this build
	SourceBlob struct {
		Checksum *string `json:"checksum" url:"checksum,key"` // an optional checksum of the gzipped tarball for verifying its
		// integrity
		URL string `json:"url" url:"url,key"` // URL where gzipped tar archive of source code for build was
		// downloaded.
		Version *string `json:"version" url:"version,key"` // Version of the gzipped tarball.
	} `json:"source_blob" url:"source_blob,key"` // location of gzipped tarball of source code used to create build
	Stack     string    `json:"stack" url:"stack,key"`           // stack of build
	Status    string    `json:"status" url:"status,key"`         // status of build
	UpdatedAt time.Time `json:"updated_at" url:"updated_at,key"` // when build was updated
	User      struct {
		Email string `json:"email" url:"email,key"` // unique email address of account
		ID    string `json:"id" url:"id,key"`       // unique identifier of an account
	} `json:"user" url:"user,key"` // user that started the build
}

// Element of PipelineDeploymentListResult
type PipelineDeployment struct {
	AddonPlanNames []string `json:"addon_plan_names" url:"addon_plan_names,key"` // add-on plans installed on the app for this release
	App            struct {
		ID   string `json:"id" url:"id,key"`     // unique identifier of app
		Name string `json:"name" url:"name,key"` // unique name of app
	} `json:"app" url:"app,key"` // app involved in the release
	CreatedAt       time.Time `json:"created_at" url:"created_at,key"`               // when release was created
	Current         bool      `json:"current" url:"current,key"`                     // indicates this release as being the current one for the app
	Description     string    `json:"description" url:"description,key"`             // description of changes in this release
	ID              string    `json:"id" url:"id,key"`                               // unique identifier of release
	OutputStreamURL *string   `json:"output_stream_url" url:"output_stream_url,key"` // Release command output will be available from this URL as a stream.
	// The stream is available as either `text/plain` or
	// `text/event-stream`. Clients should be prepared to handle disconnects
	// and can resume the stream by sending a `Range` header (for
	// `text/plain`) or a `Last-Event-Id` header (for `text/event-stream`).
	Slug *struct {
		ID string `json:"id" url:"id,key"` // unique identifier of slug
	} `json:"slug" url:"slug,key"` // slug running in this release
	Status    string    `json:"status" url:"status,key"`         // current status of the release
	UpdatedAt time.Time `json:"updated_at" url:"updated_at,key"` // when release was updated
	User      struct {
		Email string `json:"email" url:"email,key"` // unique email address of account
		ID    string `json:"id" url:"id,key"`       // unique identifier of an account
	} `json:"user" url:"user,key"` // user that created the release
	Version int `json:"version" url:"version,key"` // unique version assigned to the release
}

// Element of PipelineReleaseListResult
type PipelineRelease struct {
	AddonPlanNames []string `json:"addon_plan_names" url:"addon_plan_names,key"` // add-on plans installed on the app for this release
	App            struct {
		ID   string `json:"id" url:"id,key"`     // unique identifier of app
		Name string `json:"name" url:"name,key"` // unique name of app
	} `json:"app" url:"app,key"` // app involved in the release
	CreatedAt       time.Time `json:"created_at" url:"created_at,key"`               // when release was created
	Current         bool      `json:"current" url:"current,key"`                     // indicates this release as being the current one for the app
	Description     string    `json:"description" url:"description,key"`             // description of changes in this release
	ID              string    `json:"id" url:"id,key"`                               // unique identifier of release
	OutputStreamURL *string   `json:"output_stream_url" url:"output_stream_url,key"` // Release command output will be available from this URL as a stream.
	// The stream is available as either `text/plain` or
	// `text/event-stream`. Clients should be prepared to handle disconnects
	// and can resume the stream by sending a `Range` header (for
	// `text/plain`) or a `Last-Event-Id` header (for `text/event-stream`).
	Slug *struct {
		ID string `json:"id" url:"id,key"` // unique identifier of slug
	} `json:"slug" url:"slug,key"` // slug running in this release
	Status    string    `json:"status" url:"status,key"`         // current status of the release
	UpdatedAt time.Time `json:"updated_at" url:"updated_at,key"` // when release was updated
	User      struct {
		Email string `json:"email" url:"email,key"` // unique email address of account
		ID    string `json:"id" url:"id,key"`       // unique identifier of an account
	} `json:"user" url:"user,key"` // user that created the release
	Version int `json:"version" url:"version,key"` // unique version assigned to the release
}

// Element of TeamSpaceListResult
type TeamSpace struct {
	CIDR string `json:"cidr" url:"cidr,key"` // The RFC-1918 CIDR the Private Space will use. It must be a /16 in
	// 10.0.0.0/8, 172.16.0.0/12 or 192.168.0.0/16
	CreatedAt time.Time `json:"created_at" url:"created_at,key"` // when space was created
	DataCIDR  string    `json:"data_cidr" url:"data_cidr,key"`   // The RFC-1918 CIDR that the Private Space will use for the
	// Heroku-managed peering connection that's automatically created when
	// using Heroku Data add-ons. It must be between a /16 and a /20
	ID           string `json:"id" url:"id,key"`     // unique identifier of space
	Name         string `json:"name" url:"name,key"` // unique name of space
	Organization struct {
		Name string `json:"name" url:"name,key"` // unique name of team
	} `json:"organization" url:"organization,key"` // organization that owns this space
	Region struct {
		ID   string `json:"id" url:"id,key"`     // unique identifier of region
		Name string `json:"name" url:"name,key"` // unique name of region
	} `json:"region" url:"region,key"` // identity of space region
	Shield bool   `json:"shield" url:"shield,key"` // true if this space has shield enabled
	State  string `json:"state" url:"state,key"`   // availability of this space
	Team   struct {
		ID   string `json:"id" url:"id,key"`     // unique identifier of team
		Name string `json:"name" url:"name,key"` // unique name of team
	} `json:"team" url:"team,key"` // team that owns this space
	UpdatedAt time.Time `json:"updated_at" url:"updated_at,key"` // when space was updated
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
		HerokuStruct:        &AddOnWebhookDelivery{},
		HerokuPrimaryStruct: &heroku.AddOn{},
		Template:            "relational_resource_list",
	},
	{
		HerokuStruct:        &AddOnWebhookEvent{},
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
		HerokuStruct:        &AppWebhook{},
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
		HerokuStruct:        &PipelineBuild{},
		HerokuPrimaryStruct: &heroku.Pipeline{},
		Template:            "relational_resource_list",
	},
	{
		HerokuStruct: &heroku.PipelineCoupling{},
	},
	{
		HerokuStruct:        &PipelineDeployment{},
		HerokuPrimaryStruct: &heroku.Pipeline{},
		Template:            "relational_resource_list",
	},
	{
		HerokuStruct:        &PipelineRelease{},
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
		HerokuStruct:        &TeamSpace{},
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
		r.HerokuStructName = reflect.TypeOf(r.HerokuStruct).Elem().Name()
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
