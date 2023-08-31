package plugin

import (
	"strings"

	"github.com/cloudquery/plugin-sdk/v4/caser"
	"github.com/cloudquery/plugin-sdk/v4/docs"
	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"golang.org/x/exp/maps"
)

var (
	Version = "development"
)

var gcpExceptions = map[string]string{
	"aiplatform":           "AI Platform",
	"apigateway":           "API Gateway",
	"apikeys":              "API Keys",
	"appengine":            "App Engine",
	"artifactregistry":     "Artifact Registry",
	"baremetalsolution":    "Bare Metal Solution",
	"bigquery":             "BigQuery",
	"billingbudgets":       "Billing Budgets",
	"binaryauthorization":  "Binary Authorization",
	"cloudasset":           "Cloud Asset",
	"cloudbuild":           "Cloud Build",
	"cloudchannel":         "Cloud Channel",
	"clouddms":             "Cloud DMS",
	"cloudfunctions":       "Cloud Functions",
	"cloudidentity":        "Cloud Identity",
	"cloudiot":             "Cloud IoT",
	"cloudkms":             "Cloud KMS",
	"cloudresourcemanager": "Cloud Resource Manager",
	"cloudscheduler":       "Cloud Scheduler",
	"cloudtasks":           "Cloud Tasks",
	"cloudtrace":           "Cloud Trace",
	"composer":             "Composer",
	"compute":              "Compute",
	"container":            "Container",
	"containeranalysis":    "Container Analysis",
	"dataproc":             "Dataproc",
	"dataprocmetastore":    "Dataproc Metastore",
	"datashare":            "Datashare",
	"deploymentmanager":    "Deployment Manager",
	"dialogflow":           "Dialogflow",
	"dlp":                  "DLP",
	"domains":              "Domains",
	"ekm":                  "Cloud External Key Manager (EKM)",
	"featurestore":         "Feature Store",
	"featurestores":        "featurestores",
	"gameservices":         "Game Services",
	"http":                 "HTTP",
	"https":                "HTTPs",
	"iamcredentials":       "IAM Credentials",
	"iap":                  "IAP",
	"identityplatform":     "Identity Platform",
	"indexendpoint":        "Index Endpoint",
	"iot":                  "IoT",
	"kms":                  "Cloud Key Management Service (KMS)",
	"memcache":             "Memcache",
	"ml":                   "ML",
	"networkmanagement":    "Network Management",
	"networkservices":      "Network Services",
	"nfs":                  "NFS",
	"osconfig":             "OS Config",
	"oslogin":              "OS Login",
	"privateca":            "Private CA",
	"pubsub":               "PubSub",
	"secretmanager":        "Secret Manager",
	"specialistpool":       "Specialist Pool",
	"sql":                  "SQL",
	"ssl":                  "SSL",
	"tensorboard":          "TensorBoard",
	"videotranscoder":      "Video Transcoder",
	"vm":                   "VM",
	"vmmigration":          "VM Migration",
	"vms":                  "Virtual Machines (VMs)",
	"vpcaccess":            "VPC Access",
	"websecurityscanner":   "Web Security Scanner",
}

func titleTransformer(table *schema.Table) error {
	if table.Title != "" {
		return nil
	}
	exceptions := maps.Clone(docs.DefaultTitleExceptions)
	for k, v := range gcpExceptions {
		exceptions[k] = v
	}
	csr := caser.New(caser.WithCustomExceptions(exceptions))
	t := csr.ToTitle(table.Name)
	table.Title = strings.Trim(strings.ReplaceAll(t, "  ", " "), " ")
	return nil
}

func Plugin() *plugin.Plugin {
	return plugin.NewPlugin(
		"gcp",
		Version,
		NewClient,
	)
}
