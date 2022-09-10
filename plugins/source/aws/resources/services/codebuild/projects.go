package codebuild

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/codebuild"
	"github.com/aws/aws-sdk-go-v2/service/codebuild/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func CodebuildProjects() *schema.Table {
	return &schema.Table{
		Name:          "aws_codebuild_projects",
		Description:   "Information about a build project.",
		Resolver:      fetchCodebuildProjects,
		Multiplex:     client.ServiceAccountRegionMultiplexer("codebuild"),
		IgnoreInTests: true,
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
				Name:            "arn",
				Description:     "The Amazon Resource Name (ARN) of the build project.",
				Type:            schema.TypeString,
				CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
			},
			{
				Name:     "artifacts",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Artifacts"),
			},
			{
				Name:     "badge",
				Type:     schema.TypeJSON,
			},
			{
				Name:     "build_batch_config",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("BuildBatchConfig"),
			},
			{
				Name:     "cache",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Cache"),
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
				Name:        "environment",
				Description: "Information about a CodeBuild environment",
				Type:        schema.TypeJSON,
				Resolver:    schema.PathResolver("Environment"),
			},
			{
				Name:        "last_modified",
				Description: "When the build project's settings were last modified, expressed in Unix time format.",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:     "logs_config",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("LogsConfig"),
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
				Name:     "source",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Source"),
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
				Resolver:    client.ResolveTags,
			},
			{
				Name:        "timeout_in_minutes",
				Description: "How long, in minutes, from 5 to 480 (8 hours), for CodeBuild to wait before timing out any related build that did not get marked as completed",
				Type:        schema.TypeInt,
			},
			{
				Name:     "vpc_config",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("VpcConfig"),
			},
			{
				Name:     "webhook",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Webhook"),
			},
			{
				Name:        "file_system_locations",
				Description: "Information about a file system created by Amazon Elastic File System (EFS)",
				Type:        schema.TypeJSON,
				Resolver:    schema.PathResolver("FileSystemLocations"),
			},
			{
				Name:        "secondary_artifacts",
				Description: "Information about the build output artifacts for the build project.",
				Type:        schema.TypeJSON,
				Resolver:    schema.PathResolver("SecondaryArtifacts"),
			},
			{
				Name:        "secondary_sources",
				Description: "Information about the build input source code for the build project.",
				Type:        schema.TypeJSON,
				Resolver:    schema.PathResolver("SecondarySources"),
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchCodebuildProjects(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	svc := c.Services().Codebuild
	config := codebuild.ListProjectsInput{}
	for {
		response, err := svc.ListProjects(ctx, &config)
		if err != nil {
			return err
		}
		if len(response.Projects) == 0 {
			break
		}
		projectsOutput, err := svc.BatchGetProjects(ctx, &codebuild.BatchGetProjectsInput{Names: response.Projects})
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
	p := resource.Item.(types.Project)
	j := map[string]interface{}{}
	for _, v := range p.SecondarySourceVersions {
		j[*v.SourceIdentifier] = *v.SourceVersion
	}
	return resource.Set(c.Name, j)
}

