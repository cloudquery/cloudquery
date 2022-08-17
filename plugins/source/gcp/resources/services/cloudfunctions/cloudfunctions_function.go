package cloudfunctions

import (
	"context"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugins/source/gcp/client"
	"github.com/pkg/errors"
)

func CloudfunctionsFunction() *schema.Table {
	return &schema.Table{
		Name:        "gcp_cloudfunctions_functions",
		Description: "Describes a Cloud Function that contains user computation executed in response to an event It encapsulate function and triggers configurations",
		Resolver:    fetchCloudfunctionsFunctions,
		Multiplex:   client.ProjectMultiplex,

		Options: schema.TableCreationOptions{PrimaryKeys: []string{"project_id", "name"}},
		Columns: []schema.Column{
			{
				Name:        "project_id",
				Description: "GCP Project Id of the resource",
				Type:        schema.TypeString,
				Resolver:    client.ResolveProject,
			},
			{
				Name:        "available_memory_mb",
				Description: "The amount of memory in MB available for a function Defaults to 256MB",
				Type:        schema.TypeBigInt,
			},
			{
				Name:        "build_environment_variables",
				Description: "Build environment variables that shall be available during build time",
				Type:        schema.TypeJSON,
			},
			{
				Name:        "build_id",
				Description: "The Cloud Build ID of the latest successful deployment of the function",
				Type:        schema.TypeString,
			},
			{
				Name:        "build_worker_pool",
				Description: "Name of the Cloud Build Custom Worker Pool that should be used to build the function",
				Type:        schema.TypeString,
			},
			{
				Name:        "description",
				Description: "User-provided description of a function",
				Type:        schema.TypeString,
			},
			{
				Name:        "entry_point",
				Description: "The name of the function (as defined in source code) that will be executed Defaults to the resource name suffix, if not specified For backward compatibility, if function with given name is not found, then the system will try to use function named \"function\" For Nodejs this is name of a function exported by the module specified in `source_location`",
				Type:        schema.TypeString,
			},
			{
				Name:        "environment_variables",
				Description: "Environment variables that shall be available during function execution",
				Type:        schema.TypeJSON,
			},
			{
				Name:        "event_trigger_event_type",
				Description: "The type of event to observe",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("EventTrigger.EventType"),
			},
			{
				Name:        "event_trigger_resource",
				Description: "The resource(s) from which to observe events",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("EventTrigger.Resource"),
			},
			{
				Name:        "event_trigger_service",
				Description: "The hostname of the service that should be observed If no string is provided, the default service implementing the API will be used For example, `storagegoogleapiscom` is the default for all event types in the `googlestorage` namespace",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("EventTrigger.Service"),
			},
			{
				Name:        "https_trigger_security_level",
				Description: "The security level for the function  Possible values:   \"SECURITY_LEVEL_UNSPECIFIED\" - Unspecified   \"SECURE_ALWAYS\" - Requests for a URL that match this handler that do not use HTTPS are automatically redirected to the HTTPS URL with the same path Query parameters are reserved for the redirect   \"SECURE_OPTIONAL\" - Both HTTP and HTTPS requests with URLs that match the handler succeed without redirects The application can examine the request to determine which protocol was used and respond accordingly",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("HttpsTrigger.SecurityLevel"),
			},
			{
				Name:        "https_trigger_url",
				Description: "The deployed url for the function",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("HttpsTrigger.Url"),
			},
			{
				Name:        "ingress_settings",
				Description: "The ingress settings for the function, controlling what traffic can reach it  Possible values:   \"INGRESS_SETTINGS_UNSPECIFIED\" - Unspecified   \"ALLOW_ALL\" - Allow HTTP traffic from public and private sources   \"ALLOW_INTERNAL_ONLY\" - Allow HTTP traffic from only private VPC sources   \"ALLOW_INTERNAL_AND_GCLB\" - Allow HTTP traffic from private VPC sources and through GCLB",
				Type:        schema.TypeString,
			},
			{
				Name:        "labels",
				Description: "Labels associated with this Cloud Function",
				Type:        schema.TypeJSON,
			},
			{
				Name:        "max_instances",
				Description: "The limit on the maximum number of function instances that may coexist at a given time",
				Type:        schema.TypeBigInt,
			},
			{
				Name:        "name",
				Description: "A user-defined name of the function",
				Type:        schema.TypeString,
			},
			{
				Name:        "network",
				Description: "The VPC Network that this cloud function can connect to It can be either the fully-qualified URI, or the short name of the network resource",
				Type:        schema.TypeString,
			},
			{
				Name:        "runtime",
				Description: "The runtime in which to run the function Required when deploying a new function, optional when updating an existing function For a complete list of possible choices, see the `gcloud` command reference (/sdk/gcloud/reference/functions/deploy#--runtime)",
				Type:        schema.TypeString,
			},
			{
				Name:        "service_account_email",
				Description: "The email of the function's service account If empty, defaults to `{project_id}@appspotgserviceaccountcom`",
				Type:        schema.TypeString,
			},
			{
				Name:        "source_archive_url",
				Description: "The Google Cloud Storage URL, starting with gs://, pointing to the zip archive which contains the function",
				Type:        schema.TypeString,
			},
			{
				Name:        "source_repository_deployed_url",
				Description: "The URL pointing to the hosted repository where the function were defined at the time of deployment It always points to a specific commit in the format described above",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SourceRepository.DeployedUrl"),
			},
			{
				Name:        "source_repository_url",
				Description: "The URL pointing to the hosted repository where the function is defined",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SourceRepository.Url"),
			},
			{
				Name:        "source_token",
				Description: "Input only An identifier for Firebase function sources Disclaimer: This field is only supported for Firebase function deployments",
				Type:        schema.TypeString,
			},
			{
				Name:        "source_upload_url",
				Description: "The Google Cloud Storage signed URL used for source uploading, generated by googlecloudfunctionsv1",
				Type:        schema.TypeString,
			},
			{
				Name:        "status",
				Description: "Status of the function deployment  Possible values:   \"CLOUD_FUNCTION_STATUS_UNSPECIFIED\" - Not specified Invalid state   \"ACTIVE\" - Function has been successfully deployed and is serving   \"OFFLINE\" - Function deployment failed and the function isn’t serving   \"DEPLOY_IN_PROGRESS\" - Function is being created or updated   \"DELETE_IN_PROGRESS\" - Function is being deleted   \"UNKNOWN\" - Function deployment failed and the function serving state is undefined The function should be updated or deleted to move it out of this state",
				Type:        schema.TypeString,
			},
			{
				Name:        "timeout",
				Description: "The function execution timeout Execution is considered failed and can be terminated if the function is not completed at the end of the timeout period Defaults to 60 seconds",
				Type:        schema.TypeString,
			},
			{
				Name:        "update_time",
				Description: "The last update timestamp of a Cloud Function",
				Type:        schema.TypeString,
			},
			{
				Name:        "version_id",
				Description: "The version identifier of the Cloud Function Each deployment attempt results in a new version of a function being created",
				Type:        schema.TypeBigInt,
			},
			{
				Name:        "vpc_connector",
				Description: "The VPC Network Connector that this cloud function can connect to It can be either the fully-qualified URI, or the short name of the network connector resource The format of this field is `projects/*/locations/*/connectors/*` This field is mutually exclusive with `network` field and will eventually replace it See the VPC documentation (https://cloudgooglecom/compute/docs/vpc) for more information on connecting Cloud projects",
				Type:        schema.TypeString,
			},
			{
				Name:        "vpc_connector_egress_settings",
				Description: "The egress settings for the connector, controlling what traffic is diverted through it  Possible values:   \"VPC_CONNECTOR_EGRESS_SETTINGS_UNSPECIFIED\" - Unspecified   \"PRIVATE_RANGES_ONLY\" - Use the VPC Access Connector only for private IP space from RFC1918   \"ALL_TRAFFIC\" - Force the use of VPC Access Connector for all egress traffic from the function",
				Type:        schema.TypeString,
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchCloudfunctionsFunctions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	nextPageToken := ""
	for {
		output, err := c.Services.CloudFunctions.Projects.Locations.Functions.List("projects/" + c.ProjectId + "/locations/-").PageToken(nextPageToken).Do()
		if err != nil {
			return errors.WithStack(err)
		}

		res <- output.Functions
		if output.NextPageToken == "" {
			break
		}
		nextPageToken = output.NextPageToken
	}
	return nil
}
