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

func Jobs() *schema.Table {
	return &schema.Table{
		Name:          "k8s_batch_jobs",
		Description:   "Job represents the configuration of a single job.",
		Resolver:      fetchBatchJobs,
		Multiplex:     client.ContextMultiplex,
		DeleteFilter:  client.DeleteContextFilter,
		Options:       schema.TableCreationOptions{PrimaryKeys: []string{"uid"}},
		IgnoreInTests: true,
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
				Name:        "deletion_grace_period_seconds",
				Description: "Number of seconds allowed for this object to gracefully terminate before it will be removed from the system",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("ObjectMeta.DeletionGracePeriodSeconds"),
			},
			{
				Name:        "labels",
				Description: "Map of string keys and values that can be used to organize and categorize (scope and select) objects",
				Type:        schema.TypeJSON,
				Resolver:    schema.PathResolver("ObjectMeta.Labels"),
			},
			{
				Name:        "annotations",
				Description: "Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata",
				Type:        schema.TypeJSON,
				Resolver:    schema.PathResolver("ObjectMeta.Annotations"),
			},
			{
				Name:        "owner_references",
				Description: "List of objects depended by this object",
				Type:        schema.TypeJSON,
				Resolver:    resolveBatchJobsOwnerReferences,
			},
			{
				Name:        "finalizers",
				Description: "Must be empty before the object is deleted from the registry",
				Type:        schema.TypeStringArray,
				Resolver:    schema.PathResolver("ObjectMeta.Finalizers"),
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
				Resolver:    resolveBatchJobsManagedFields,
			},
			{
				Name:        "parallelism",
				Description: "Specifies the maximum desired number of pods the job should run at any given time",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("Spec.Parallelism"),
			},
			{
				Name:        "completions",
				Description: "Specifies the desired number of successfully finished pods the job should be run with",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("Spec.Completions"),
			},
			{
				Name:        "active_deadline_seconds",
				Description: "Specifies the duration in seconds relative to the startTime that the job may be continuously active before the system tries to terminate it; value must be positive integer",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("Spec.ActiveDeadlineSeconds"),
			},
			{
				Name:        "backoff_limit",
				Description: "Specifies the number of retries before marking this job failed. Defaults to 6",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("Spec.BackoffLimit"),
			},
			{
				Name:        "selector_match_labels",
				Description: "matchLabels is a map of {key,value} pairs",
				Type:        schema.TypeJSON,
				Resolver:    schema.PathResolver("Spec.Selector.MatchLabels"),
			},
			{
				Name:        "manual_selector",
				Description: "manualSelector controls generation of pod labels and pod selectors. Leave `manualSelector` unset unless you are certain what you are doing. When false or unset, the system pick labels unique to this job and appends those labels to the pod template",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("Spec.ManualSelector"),
			},
			{
				Name:        "template",
				Description: "Describes the pod that will be created when executing a job. More info: https://kubernetes.io/docs/concepts/workloads/controllers/jobs-run-to-completion/",
				Type:        schema.TypeJSON,
				Resolver:    resolveBatchJobsTemplate,
			},
			{
				Name:        "ttl_seconds_after_finished",
				Description: "ttlSecondsAfterFinished limits the lifetime of a Job that has finished execution (either Complete or Failed)",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("Spec.TTLSecondsAfterFinished"),
			},
			{
				Name:        "completion_mode",
				Description: "CompletionMode specifies how Pod completions are tracked",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Spec.CompletionMode"),
			},
			{
				Name:        "suspend",
				Description: "Suspend specifies whether the Job controller should create Pods or not",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("Spec.Suspend"),
			},
			{
				Name:        "status_active",
				Description: "The number of actively running pods.",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("Status.Active"),
			},
			{
				Name:        "status_succeeded",
				Description: "The number of pods which reached phase Succeeded.",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("Status.Succeeded"),
			},
			{
				Name:        "status_failed",
				Description: "The number of pods which reached phase Failed.",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("Status.Failed"),
			},
			{
				Name:        "status_completed_indexes",
				Description: "CompletedIndexes holds the completed indexes when .spec.completionMode = \"Indexed\" in a text format",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Status.CompletedIndexes"),
			},
			{
				Name:        "status_uncounted_terminated_pods_succeeded",
				Description: "Succeeded holds UIDs of succeeded Pods. +listType=set",
				Type:        schema.TypeStringArray,
				Resolver:    schema.PathResolver("Status.UncountedTerminatedPods.Succeeded"),
			},
			{
				Name:        "status_uncounted_terminated_pods_failed",
				Description: "Failed holds UIDs of failed Pods. +listType=set",
				Type:        schema.TypeStringArray,
				Resolver:    schema.PathResolver("Status.UncountedTerminatedPods.Failed"),
			},
		},
		Relations: []*schema.Table{
			{
				Name:          "k8s_batch_job_selector_match_expressions",
				Description:   "A label selector requirement is a selector that contains values, a key, and an operator that relates the key and values.",
				Resolver:      fetchBatchJobSelectorMatchExpressions,
				IgnoreInTests: true,
				Columns: []schema.Column{
					{
						Name:        "job_cq_id",
						Description: "Unique CloudQuery ID of k8s_batch_jobs table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "key",
						Description: "key is the label key that the selector applies to. +patchMergeKey=key +patchStrategy=merge",
						Type:        schema.TypeString,
					},
					{
						Name:        "operator",
						Description: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
						Type:        schema.TypeString,
					},
					{
						Name:        "values",
						Description: "values is an array of string values",
						Type:        schema.TypeStringArray,
					},
				},
			},
			{
				Name:          "k8s_batch_job_status_conditions",
				Description:   "JobCondition describes current state of a job.",
				Resolver:      fetchBatchJobStatusConditions,
				IgnoreInTests: true,
				Columns: []schema.Column{
					{
						Name:        "job_cq_id",
						Description: "Unique CloudQuery ID of k8s_batch_jobs table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "type",
						Description: "Type of job condition, Complete or Failed.",
						Type:        schema.TypeString,
					},
					{
						Name:        "status",
						Description: "Status of the condition, one of True, False, Unknown.",
						Type:        schema.TypeString,
					},
					{
						Name:        "reason",
						Description: "(brief) reason for the condition's last transition.",
						Type:        schema.TypeString,
					},
					{
						Name:        "message",
						Description: "Human readable message indicating details about last transition.",
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

func fetchBatchJobs(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	client := meta.(*client.Client).Services().Jobs
	opts := metav1.ListOptions{}
	for {
		result, err := client.List(ctx, opts)
		if err != nil {
			return err
		}
		res <- result.Items
		if result.GetContinue() == "" {
			return nil
		}
		opts.Continue = result.GetContinue()
	}
}
func resolveBatchJobsOwnerReferences(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p, ok := resource.Item.(batchv1.Job)
	if !ok {
		return fmt.Errorf("not a batchv1.Job instance: %T", resource.Item)
	}
	b, err := json.Marshal(p.OwnerReferences)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, b)
}
func resolveBatchJobsManagedFields(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p, ok := resource.Item.(batchv1.Job)
	if !ok {
		return fmt.Errorf("not a batchv1.Job instance: %T", resource.Item)
	}

	b, err := json.Marshal(p.ManagedFields)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, b)
}
func resolveBatchJobsTemplate(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p, ok := resource.Item.(batchv1.Job)
	if !ok {
		return fmt.Errorf("not a batchv1.Job instance: %T", resource.Item)
	}

	b, err := json.Marshal(p.Spec.Template)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, b)
}
func fetchBatchJobSelectorMatchExpressions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	job, ok := parent.Item.(batchv1.Job)
	if !ok {
		return fmt.Errorf("not a batchv1.Job instance: %T", parent.Item)
	}

	if job.Spec.Selector == nil {
		return nil
	}
	res <- job.Spec.Selector.MatchExpressions
	return nil
}
func fetchBatchJobStatusConditions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	job, ok := parent.Item.(batchv1.Job)
	if !ok {
		return fmt.Errorf("not a batchv1.Job instance: %T", parent.Item)
	}

	res <- job.Status.Conditions
	return nil
}
