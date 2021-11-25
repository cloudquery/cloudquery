package resources

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/codebuild"
	"github.com/aws/aws-sdk-go-v2/service/codebuild/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func CodebuildProjects() *schema.Table {
	return &schema.Table{
		Name:         "aws_codebuild_projects",
		Description:  "Information about a build project.",
		Resolver:     fetchCodebuildProjects,
		Multiplex:    client.AccountRegionMultiplex,
		IgnoreError:  client.IgnoreAccessDeniedServiceDisabled,
		DeleteFilter: client.DeleteAccountRegionFilter,
		Options:      schema.TableCreationOptions{PrimaryKeys: []string{"arn"}},
		Columns: []schema.Column{
			{
				Name:        "account_id",
				Description: "The AWS Account ID of the resource.",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSAccount,
			},
			{
				Name:        "region",
				Description: "The AWS Region of the resource.",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSRegion,
			},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) of the build project.",
				Type:        schema.TypeString,
			},
			{
				Name:        "artifacts_type",
				Description: "The type of build output artifact",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Artifacts.Type"),
			},
			{
				Name:        "artifacts_artifact_identifier",
				Description: "An identifier for this artifact definition.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Artifacts.ArtifactIdentifier"),
			},
			{
				Name:        "artifacts_bucket_owner_access",
				Description: "Specifies the bucket owner's access for objects that another account uploads to their Amazon S3 bucket",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Artifacts.BucketOwnerAccess"),
			},
			{
				Name:        "artifacts_encryption_disabled",
				Description: "Set to true if you do not want your output artifacts encrypted",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("Artifacts.EncryptionDisabled"),
			},
			{
				Name:        "artifacts_location",
				Description: "Information about the build output artifact location:  * If type is set to CODEPIPELINE, CodePipeline ignores this value if specified",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Artifacts.Location"),
			},
			{
				Name:        "artifacts_name",
				Description: "Along with path and namespaceType, the pattern that CodeBuild uses to name and store the output artifact:  * If type is set to CODEPIPELINE, CodePipeline ignores this value if specified",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Artifacts.Name"),
			},
			{
				Name:        "artifacts_namespace_type",
				Description: "Along with path and name, the pattern that CodeBuild uses to determine the name and location to store the output artifact:  * If type is set to CODEPIPELINE, CodePipeline ignores this value if specified",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Artifacts.NamespaceType"),
			},
			{
				Name:        "artifacts_override_artifact_name",
				Description: "If this flag is set, a name specified in the buildspec file overrides the artifact name",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("Artifacts.OverrideArtifactName"),
			},
			{
				Name:        "artifacts_packaging",
				Description: "The type of build output artifact to create:  * If type is set to CODEPIPELINE, CodePipeline ignores this value if specified",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Artifacts.Packaging"),
			},
			{
				Name:        "artifacts_path",
				Description: "Along with namespaceType and name, the pattern that CodeBuild uses to name and store the output artifact:  * If type is set to CODEPIPELINE, CodePipeline ignores this value if specified",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Artifacts.Path"),
			},
			{
				Name:        "badge_enabled",
				Description: "Set this to true to generate a publicly accessible URL for your project's build badge.",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("Badge.BadgeEnabled"),
			},
			{
				Name:        "badge_request_url",
				Description: "The publicly-accessible URL through which you can access the build badge for your project.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Badge.BadgeRequestUrl"),
			},
			{
				Name:        "build_batch_config_batch_report_mode",
				Description: "Specifies how build status reports are sent to the source provider for the batch build",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("BuildBatchConfig.BatchReportMode"),
			},
			{
				Name:        "build_batch_config_combine_artifacts",
				Description: "Specifies if the build artifacts for the batch build should be combined into a single artifact location.",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("BuildBatchConfig.CombineArtifacts"),
			},
			{
				Name:        "build_batch_config_restrictions_compute_types_allowed",
				Description: "An array of strings that specify the compute types that are allowed for the batch build",
				Type:        schema.TypeStringArray,
				Resolver:    schema.PathResolver("BuildBatchConfig.Restrictions.ComputeTypesAllowed"),
			},
			{
				Name:        "build_batch_config_restrictions_maximum_builds_allowed",
				Description: "Specifies the maximum number of builds allowed.",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("BuildBatchConfig.Restrictions.MaximumBuildsAllowed"),
			},
			{
				Name:        "build_batch_config_service_role",
				Description: "Specifies the service role ARN for the batch build project.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("BuildBatchConfig.ServiceRole"),
			},
			{
				Name:        "build_batch_config_timeout_in_mins",
				Description: "Specifies the maximum amount of time, in minutes, that the batch build must be completed in.",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("BuildBatchConfig.TimeoutInMins"),
			},
			{
				Name:        "cache_type",
				Description: "The type of cache used by the build project",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Cache.Type"),
			},
			{
				Name:        "cache_location",
				Description: "Information about the cache location:  * NO_CACHE or LOCAL: This value is ignored.  * S3: This is the S3 bucket name/prefix.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Cache.Location"),
			},
			{
				Name:        "cache_modes",
				Description: "An array of strings that specify the local cache modes",
				Type:        schema.TypeStringArray,
				Resolver:    schema.PathResolver("Cache.Modes"),
			},
			{
				Name:        "concurrent_build_limit",
				Description: "The maximum number of concurrent builds that are allowed for this project",
				Type:        schema.TypeInt,
			},
			{
				Name:        "created",
				Description: "When the build project was created, expressed in Unix time format.",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "description",
				Description: "A description that makes the build project easy to identify.",
				Type:        schema.TypeString,
			},
			{
				Name:        "encryption_key",
				Description: "The Key Management Service customer master key (CMK) to be used for encrypting the build output artifacts",
				Type:        schema.TypeString,
			},
			{
				Name:        "environment_compute_type",
				Description: "Information about the compute resources the build project uses",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Environment.ComputeType"),
			},
			{
				Name:        "environment_image",
				Description: "The image tag or image digest that identifies the Docker image to use for this build project",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Environment.Image"),
			},
			{
				Name:        "environment_type",
				Description: "The type of build environment to use for related builds.  * The environment type ARM_CONTAINER is available only in regions US East (N",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Environment.Type"),
			},
			{
				Name:        "environment_certificate",
				Description: "The ARN of the Amazon S3 bucket, path prefix, and object key that contains the PEM-encoded certificate for the build project",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Environment.Certificate"),
			},
			{
				Name:        "environment_image_pull_credentials_type",
				Description: "The type of credentials CodeBuild uses to pull images in your build",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Environment.ImagePullCredentialsType"),
			},
			{
				Name:        "environment_privileged_mode",
				Description: "Enables running the Docker daemon inside a Docker container",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("Environment.PrivilegedMode"),
			},
			{
				Name:        "environment_registry_credential",
				Description: "The Amazon Resource Name (ARN) or name of credentials created using Secrets Manager",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Environment.RegistryCredential.Credential"),
			},
			{
				Name:        "environment_registry_credential_credential_provider",
				Description: "The service that created the credentials to access a private Docker registry. The valid value, SECRETS_MANAGER, is for Secrets Manager.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Environment.RegistryCredential.CredentialProvider"),
			},
			{
				Name:        "last_modified",
				Description: "When the build project's settings were last modified, expressed in Unix time format.",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "logs_config_cloud_watch_logs_status",
				Description: "The current status of the logs in CloudWatch Logs for a build project",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("LogsConfig.CloudWatchLogs.Status"),
			},
			{
				Name:        "logs_config_cloud_watch_logs_group_name",
				Description: "The group name of the logs in CloudWatch Logs",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("LogsConfig.CloudWatchLogs.GroupName"),
			},
			{
				Name:        "logs_config_cloud_watch_logs_stream_name",
				Description: "The prefix of the stream name of the CloudWatch Logs",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("LogsConfig.CloudWatchLogs.StreamName"),
			},
			{
				Name:        "logs_config_s3_logs_status",
				Description: "The current status of the S3 build logs",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("LogsConfig.S3Logs.Status"),
			},
			{
				Name:        "logs_config_s3_logs_bucket_owner_access",
				Description: "Specifies the bucket owner's access for objects that another account uploads to their Amazon S3 bucket",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("LogsConfig.S3Logs.BucketOwnerAccess"),
			},
			{
				Name:        "logs_config_s3_logs_encryption_disabled",
				Description: "Set to true if you do not want your S3 build log output encrypted",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("LogsConfig.S3Logs.EncryptionDisabled"),
			},
			{
				Name:        "logs_config_s3_logs_location",
				Description: "The ARN of an S3 bucket and the path prefix for S3 logs",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("LogsConfig.S3Logs.Location"),
			},
			{
				Name:        "name",
				Description: "The name of the build project.",
				Type:        schema.TypeString,
			},
			{
				Name:        "project_visibility",
				Description: "Specifies the visibility of the project's builds",
				Type:        schema.TypeString,
			},
			{
				Name:        "public_project_alias",
				Description: "Contains the project identifier used with the public build APIs.",
				Type:        schema.TypeString,
			},
			{
				Name:        "queued_timeout_in_minutes",
				Description: "The number of minutes a build is allowed to be queued before it times out.",
				Type:        schema.TypeInt,
			},
			{
				Name:        "resource_access_role",
				Description: "The ARN of the IAM role that enables CodeBuild to access the CloudWatch Logs and Amazon S3 artifacts for the project's builds.",
				Type:        schema.TypeString,
			},
			{
				Name:        "secondary_source_versions",
				Description: "An array of ProjectSourceVersion objects",
				Type:        schema.TypeJSON,
				Resolver:    resolveCodebuildProjectsSecondarySourceVersions,
			},
			{
				Name:        "service_role",
				Description: "The ARN of the IAM role that enables CodeBuild to interact with dependent Amazon Web Services services on behalf of the Amazon Web Services account.",
				Type:        schema.TypeString,
			},
			{
				Name:        "source_type",
				Description: "The type of repository that contains the source code to be built",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Source.Type"),
			},
			{
				Name:        "source_auth_type",
				Description: "This data type is deprecated and is no longer accurate or used",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Source.Auth.Type"),
			},
			{
				Name:        "source_auth_resource",
				Description: "The resource value that applies to the specified authorization type.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Source.Auth.Resource"),
			},
			{
				Name:        "source_build_status_config_context",
				Description: "Specifies the context of the build status CodeBuild sends to the source provider",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Source.BuildStatusConfig.Context"),
			},
			{
				Name:        "source_build_status_config_target_url",
				Description: "Specifies the target url of the build status CodeBuild sends to the source provider",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Source.BuildStatusConfig.TargetUrl"),
			},
			{
				Name:        "source_buildspec",
				Description: "The buildspec file declaration to use for the builds in this build project",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Source.Buildspec"),
			},
			{
				Name:        "source_git_clone_depth",
				Description: "Information about the Git clone depth for the build project.",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("Source.GitCloneDepth"),
			},
			{
				Name:        "source_git_submodules_config_fetch_submodules",
				Description: "Set to true to fetch Git submodules for your CodeBuild build project.",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("Source.GitSubmodulesConfig.FetchSubmodules"),
			},
			{
				Name:        "source_insecure_ssl",
				Description: "Enable this flag to ignore SSL warnings while connecting to the project source code.",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("Source.InsecureSsl"),
			},
			{
				Name:        "source_location",
				Description: "Information about the location of the source code to be built",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Source.Location"),
			},
			{
				Name:        "source_report_build_status",
				Description: "Set to true to report the status of a build's start and finish to your source provider",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("Source.ReportBuildStatus"),
			},
			{
				Name:        "source_identifier",
				Description: "An identifier for this project source",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Source.SourceIdentifier"),
			},
			{
				Name:        "source_version",
				Description: "A version of the build input to be built for this project",
				Type:        schema.TypeString,
			},
			{
				Name:        "tags",
				Description: "A list of tag key and value pairs associated with this build project",
				Type:        schema.TypeJSON,
				Resolver:    resolveCodebuildProjectsTags,
			},
			{
				Name:        "timeout_in_minutes",
				Description: "How long, in minutes, from 5 to 480 (8 hours), for CodeBuild to wait before timing out any related build that did not get marked as completed",
				Type:        schema.TypeInt,
			},
			{
				Name:        "vpc_config_security_group_ids",
				Description: "A list of one or more security groups IDs in your Amazon VPC.",
				Type:        schema.TypeStringArray,
				Resolver:    schema.PathResolver("VpcConfig.SecurityGroupIds"),
			},
			{
				Name:        "vpc_config_subnets",
				Description: "A list of one or more subnet IDs in your Amazon VPC.",
				Type:        schema.TypeStringArray,
				Resolver:    schema.PathResolver("VpcConfig.Subnets"),
			},
			{
				Name:        "vpc_config_vpc_id",
				Description: "The ID of the Amazon VPC.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("VpcConfig.VpcId"),
			},
			{
				Name:        "webhook_branch_filter",
				Description: "A regular expression used to determine which repository branches are built when a webhook is triggered",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Webhook.BranchFilter"),
			},
			{
				Name:        "webhook_build_type",
				Description: "Specifies the type of build this webhook will trigger.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Webhook.BuildType"),
			},
			{
				Name:        "webhook_filter_groups",
				Description: "An array of arrays of WebhookFilter objects used to determine which webhooks are triggered",
				Type:        schema.TypeJSON,
				Resolver:    resolveCodebuildProjectsWebhookFilterGroups,
			},
			{
				Name:        "webhook_last_modified_secret",
				Description: "A timestamp that indicates the last time a repository's secret token was modified.",
				Type:        schema.TypeTimestamp,
				Resolver:    schema.PathResolver("Webhook.LastModifiedSecret"),
			},
			{
				Name:        "webhook_payload_url",
				Description: "The CodeBuild endpoint where webhook events are sent.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Webhook.PayloadUrl"),
			},
			{
				Name:        "webhook_secret",
				Description: "The secret token of the associated repository",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Webhook.Secret"),
			},
			{
				Name:        "webhook_url",
				Description: "The URL to the webhook.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Webhook.Url"),
			},
		},
		Relations: []*schema.Table{
			{
				Name:        "aws_codebuild_project_environment_variables",
				Description: "Information about an environment variable for a build project or a build.",
				Resolver:    fetchCodebuildProjectEnvironmentVariables,
				Columns: []schema.Column{
					{
						Name:        "project_cq_id",
						Description: "Unique CloudQuery ID of aws_codebuild_projects table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "name",
						Description: "The name or key of the environment variable.",
						Type:        schema.TypeString,
					},
					{
						Name:        "value",
						Description: "The value of the environment variable",
						Type:        schema.TypeString,
					},
					{
						Name:        "type",
						Description: "The type of environment variable",
						Type:        schema.TypeString,
					},
				},
			},
			{
				Name:        "aws_codebuild_project_file_system_locations",
				Description: "Information about a file system created by Amazon Elastic File System (EFS)",
				Resolver:    fetchCodebuildProjectFileSystemLocations,
				Columns: []schema.Column{
					{
						Name:        "project_cq_id",
						Description: "Unique CloudQuery ID of aws_codebuild_projects table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "identifier",
						Description: "The name used to access a file system created by Amazon EFS",
						Type:        schema.TypeString,
					},
					{
						Name:        "location",
						Description: "A string that specifies the location of the file system created by Amazon EFS. Its format is efs-dns-name:/directory-path",
						Type:        schema.TypeString,
					},
					{
						Name:        "mount_options",
						Description: "The mount options for a file system created by Amazon EFS",
						Type:        schema.TypeString,
					},
					{
						Name:        "mount_point",
						Description: "The location in the container where you mount the file system.",
						Type:        schema.TypeString,
					},
					{
						Name:        "type",
						Description: "The type of the file system",
						Type:        schema.TypeString,
					},
				},
			},
			{
				Name:        "aws_codebuild_project_secondary_artifacts",
				Description: "Information about the build output artifacts for the build project.",
				Resolver:    fetchCodebuildProjectSecondaryArtifacts,
				Columns: []schema.Column{
					{
						Name:        "project_cq_id",
						Description: "Unique CloudQuery ID of aws_codebuild_projects table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "type",
						Description: "The type of build output artifact",
						Type:        schema.TypeString,
					},
					{
						Name:        "artifact_identifier",
						Description: "An identifier for this artifact definition.",
						Type:        schema.TypeString,
					},
					{
						Name:        "bucket_owner_access",
						Description: "Specifies the bucket owner's access for objects that another account uploads to their Amazon S3 bucket",
						Type:        schema.TypeString,
					},
					{
						Name:        "encryption_disabled",
						Description: "Set to true if you do not want your output artifacts encrypted",
						Type:        schema.TypeBool,
					},
					{
						Name:        "location",
						Description: "Information about the build output artifact location:  * If type is set to CODEPIPELINE, CodePipeline ignores this value if specified",
						Type:        schema.TypeString,
					},
					{
						Name:        "name",
						Description: "Along with path and namespaceType, the pattern that CodeBuild uses to name and store the output artifact:  * If type is set to CODEPIPELINE, CodePipeline ignores this value if specified",
						Type:        schema.TypeString,
					},
					{
						Name:        "namespace_type",
						Description: "Along with path and name, the pattern that CodeBuild uses to determine the name and location to store the output artifact:  * If type is set to CODEPIPELINE, CodePipeline ignores this value if specified",
						Type:        schema.TypeString,
					},
					{
						Name:        "override_artifact_name",
						Description: "If this flag is set, a name specified in the buildspec file overrides the artifact name",
						Type:        schema.TypeBool,
					},
					{
						Name:        "packaging",
						Description: "The type of build output artifact to create:  * If type is set to CODEPIPELINE, CodePipeline ignores this value if specified",
						Type:        schema.TypeString,
					},
					{
						Name:        "path",
						Description: "Along with namespaceType and name, the pattern that CodeBuild uses to name and store the output artifact:  * If type is set to CODEPIPELINE, CodePipeline ignores this value if specified",
						Type:        schema.TypeString,
					},
				},
			},
			{
				Name:        "aws_codebuild_project_secondary_sources",
				Description: "Information about the build input source code for the build project.",
				Resolver:    fetchCodebuildProjectSecondarySources,
				Columns: []schema.Column{
					{
						Name:        "project_cq_id",
						Description: "Unique CloudQuery ID of aws_codebuild_projects table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "type",
						Description: "The type of repository that contains the source code to be built",
						Type:        schema.TypeString,
					},
					{
						Name:        "auth_type",
						Description: "This data type is deprecated and is no longer accurate or used",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Auth.Type"),
					},
					{
						Name:        "auth_resource",
						Description: "The resource value that applies to the specified authorization type.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Auth.Resource"),
					},
					{
						Name:        "build_status_config_context",
						Description: "Specifies the context of the build status CodeBuild sends to the source provider",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("BuildStatusConfig.Context"),
					},
					{
						Name:        "build_status_config_target_url",
						Description: "Specifies the target url of the build status CodeBuild sends to the source provider",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("BuildStatusConfig.TargetUrl"),
					},
					{
						Name:        "buildspec",
						Description: "The buildspec file declaration to use for the builds in this build project",
						Type:        schema.TypeString,
					},
					{
						Name:        "git_clone_depth",
						Description: "Information about the Git clone depth for the build project.",
						Type:        schema.TypeInt,
					},
					{
						Name:        "git_submodules_config_fetch_submodules",
						Description: "Set to true to fetch Git submodules for your CodeBuild build project.",
						Type:        schema.TypeBool,
						Resolver:    schema.PathResolver("GitSubmodulesConfig.FetchSubmodules"),
					},
					{
						Name:        "insecure_ssl",
						Description: "Enable this flag to ignore SSL warnings while connecting to the project source code.",
						Type:        schema.TypeBool,
					},
					{
						Name:        "location",
						Description: "Information about the location of the source code to be built",
						Type:        schema.TypeString,
					},
					{
						Name:        "report_build_status",
						Description: "Set to true to report the status of a build's start and finish to your source provider",
						Type:        schema.TypeBool,
					},
					{
						Name:        "source_identifier",
						Description: "An identifier for this project source",
						Type:        schema.TypeString,
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchCodebuildProjects(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	c := meta.(*client.Client)
	svc := c.Services().Codebuild
	config := codebuild.ListProjectsInput{}
	for {
		response, err := svc.ListProjects(ctx, &config, func(options *codebuild.Options) {
			options.Region = c.Region
		})
		if err != nil {
			return err
		}
		if len(response.Projects) == 0 {
			break
		}
		projectsOutput, err := svc.BatchGetProjects(ctx, &codebuild.BatchGetProjectsInput{Names: response.Projects}, func(options *codebuild.Options) {
			options.Region = c.Region
		})
		if err != nil {
			return err
		}

		res <- projectsOutput.Projects
		if aws.ToString(response.NextToken) == "" {
			break
		}
		config.NextToken = response.NextToken
	}
	return nil
}
func resolveCodebuildProjectsSecondarySourceVersions(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p, ok := resource.Item.(types.Project)
	if !ok {
		return fmt.Errorf("not a types.Project instance: %T", resource.Item)
	}
	j := map[string]interface{}{}
	for _, v := range p.SecondarySourceVersions {
		j[*v.SourceIdentifier] = *v.SourceVersion
	}
	return resource.Set(c.Name, j)
}
func resolveCodebuildProjectsTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p, ok := resource.Item.(types.Project)
	if !ok {
		return fmt.Errorf("not a types.Project instance: %T", resource.Item)
	}
	j := map[string]interface{}{}
	for _, v := range p.Tags {
		j[*v.Key] = *v.Value
	}
	return resource.Set(c.Name, j)
}
func resolveCodebuildProjectsWebhookFilterGroups(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p, ok := resource.Item.(types.Project)
	if !ok {
		return fmt.Errorf("not a types.Project instance: %T", resource.Item)
	}
	if p.Webhook == nil {
		return nil
	}
	data, err := json.Marshal(p.Webhook.FilterGroups)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, data)
}
func fetchCodebuildProjectEnvironmentVariables(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p, ok := parent.Item.(types.Project)
	if !ok {
		return fmt.Errorf("not a types.Project instance: %T", parent.Item)
	}
	if p.Environment == nil {
		return nil
	}
	res <- p.Environment.EnvironmentVariables
	return nil
}
func fetchCodebuildProjectFileSystemLocations(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p, ok := parent.Item.(types.Project)
	if !ok {
		return fmt.Errorf("not a types.Project instance: %T", parent.Item)
	}
	res <- p.FileSystemLocations
	return nil
}
func fetchCodebuildProjectSecondaryArtifacts(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p, ok := parent.Item.(types.Project)
	if !ok {
		return fmt.Errorf("not a types.Project instance: %T", parent.Item)
	}
	res <- p.SecondaryArtifacts
	return nil
}
func fetchCodebuildProjectSecondarySources(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p, ok := parent.Item.(types.Project)
	if !ok {
		return fmt.Errorf("not a types.Project instance: %T", parent.Item)
	}
	res <- p.SecondarySources
	return nil
}
