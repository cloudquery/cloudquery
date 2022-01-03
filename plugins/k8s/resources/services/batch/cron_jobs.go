package batch

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/cloudquery/cq-provider-k8s/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	batchv1 "k8s.io/api/batch/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func CronJobs() *schema.Table {
	return &schema.Table{
		Name:         "k8s_batch_cron_jobs",
		Description:  "CronJob represents the configuration of a single cron job.",
		Resolver:     fetchBatchCronJobs,
		Multiplex:    client.ContextMultiplex,
		DeleteFilter: client.DeleteContextFilter,
		Options:      schema.TableCreationOptions{PrimaryKeys: []string{"uid"}},
		Columns: []schema.Column{
			client.CommonContextField,
			{
				Name:        "name",
				Description: "Name must be unique within a namespace",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ObjectMeta.Name"),
			},
			{
				Name:        "generate_name",
				Description: "GenerateName is an optional prefix, used by the server, to generate a unique name ONLY IF the Name field has not been provided. If this field is used, the name returned to the client will be different than the name passed",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ObjectMeta.GenerateName"),
			},
			{
				Name:        "namespace",
				Description: "Namespace defines the space within which each name must be unique",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ObjectMeta.Namespace"),
			},
			{
				Name:        "self_link",
				Description: "SelfLink is a URL representing this object. Populated by the system. Read-only.  DEPRECATED Kubernetes will stop propagating this field in 1.20 release and the field is planned to be removed in 1.21 release.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ObjectMeta.SelfLink"),
			},
			{
				Name:        "uid",
				Description: "UID is the unique in time and space value for this object",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ObjectMeta.UID"),
			},
			{
				Name:        "resource_version",
				Description: "An opaque value that represents the internal version of this object that can be used by clients to determine when objects have changed",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ObjectMeta.ResourceVersion"),
			},
			{
				Name:        "generation",
				Description: "A sequence number representing a specific generation of the desired state. Populated by the system",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("ObjectMeta.Generation"),
			},
			{
				Name:          "deletion_grace_period_seconds",
				Description:   "Number of seconds allowed for this object to gracefully terminate before it will be removed from the system",
				Type:          schema.TypeBigInt,
				Resolver:      schema.PathResolver("ObjectMeta.DeletionGracePeriodSeconds"),
				IgnoreInTests: true,
			},
			{
				Name:          "labels",
				Description:   "Map of string keys and values that can be used to organize and categorize (scope and select) objects",
				Type:          schema.TypeJSON,
				Resolver:      schema.PathResolver("ObjectMeta.Labels"),
				IgnoreInTests: true,
			},
			{
				Name:          "annotations",
				Description:   "Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata",
				Type:          schema.TypeJSON,
				Resolver:      schema.PathResolver("ObjectMeta.Annotations"),
				IgnoreInTests: true,
			},
			{
				Name:          "owner_references",
				Description:   "List of objects depended by this object",
				Type:          schema.TypeJSON,
				Resolver:      resolveBatchCronJobsOwnerReferences,
				IgnoreInTests: true,
			},
			{
				Name:          "finalizers",
				Description:   "Must be empty before the object is deleted from the registry",
				Type:          schema.TypeStringArray,
				Resolver:      schema.PathResolver("ObjectMeta.Finalizers"),
				IgnoreInTests: true,
			},
			{
				Name:        "cluster_name",
				Description: "The name of the cluster which the object belongs to. This is used to distinguish resources with same name and namespace in different clusters. This field is not set anywhere right now and apiserver is going to ignore it if set in create or update request.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ObjectMeta.ClusterName"),
			},
			{
				Name:        "managed_fields",
				Description: "ManagedFields maps workflow-id and version to the set of fields that are managed by that workflow",
				Type:        schema.TypeJSON,
				Resolver:    resolveBatchCronJobsManagedFields,
			},
			{
				Name:        "schedule",
				Description: "The schedule in Cron format, see https://en.wikipedia.org/wiki/Cron.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Spec.Schedule"),
			},
			{
				Name:        "starting_deadline_seconds",
				Description: "Optional deadline in seconds for starting the job if it misses scheduled time for any reason",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("Spec.StartingDeadlineSeconds"),
			},
			{
				Name:        "concurrency_policy",
				Description: "Specifies how to treat concurrent executions of a Job. Valid values are: - \"Allow\" (default): allows CronJobs to run concurrently; - \"Forbid\": forbids concurrent runs, skipping next run if previous run hasn't finished yet; - \"Replace\": cancels currently running job and replaces it with a new one",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Spec.ConcurrencyPolicy"),
			},
			{
				Name:        "suspend",
				Description: "This flag tells the controller to suspend subsequent executions, it does not apply to already started executions",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("Spec.Suspend"),
			},
			{
				Name:        "job_template",
				Description: "Specifies the job that will be created when executing a CronJob.",
				Type:        schema.TypeJSON,
				Resolver:    resolveBatchCronJobsJobTemplate,
			},
			{
				Name:        "successful_jobs_history_limit",
				Description: "The number of successful finished jobs to retain",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("Spec.SuccessfulJobsHistoryLimit"),
			},
			{
				Name:        "failed_jobs_history_limit",
				Description: "The number of failed finished jobs to retain",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("Spec.FailedJobsHistoryLimit"),
			},
			{
				Name:        "status",
				Description: "Current status of a cron job. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status",
				Type:        schema.TypeJSON,
				Resolver:    resolveBatchCronJobsStatus,
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchBatchCronJobs(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	jobs := meta.(*client.Client).Services().CronJobs
	opts := metav1.ListOptions{}
	for {
		result, err := jobs.List(ctx, opts)
		if err != nil {
			return err
		}
		res <- result.Items
		next := result.GetContinue()
		if next == "" {
			return nil
		}
		opts.Continue = next
	}
}
func resolveBatchCronJobsOwnerReferences(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cronJob, ok := resource.Item.(batchv1.CronJob)
	if !ok {
		return fmt.Errorf("not a batchv1.CronJob instance: %T", resource.Item)
	}
	b, err := json.Marshal(cronJob.OwnerReferences)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, b)
}
func resolveBatchCronJobsManagedFields(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cronJob, ok := resource.Item.(batchv1.CronJob)
	if !ok {
		return fmt.Errorf("not a batchv1.CronJob instance: %T", resource.Item)
	}
	b, err := json.Marshal(cronJob.ManagedFields)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, b)
}
func resolveBatchCronJobsJobTemplate(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cronJob, ok := resource.Item.(batchv1.CronJob)
	if !ok {
		return fmt.Errorf("not a batchv1.CronJob instance: %T", resource.Item)
	}
	b, err := json.Marshal(cronJob.Spec.JobTemplate)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, b)
}
func resolveBatchCronJobsStatus(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cronJob, ok := resource.Item.(batchv1.CronJob)
	if !ok {
		return fmt.Errorf("not a batchv1.CronJob instance: %T", resource.Item)
	}
	b, err := json.Marshal(cronJob.Status)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, b)
}
