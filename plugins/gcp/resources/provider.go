package resources

import (
	"github.com/cloudquery/cq-provider-gcp/client"
	"github.com/cloudquery/cq-provider-sdk/provider"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func Provider() *provider.Provider {
	return &provider.Provider{
		Name:      "gcp",
		Configure: client.Configure,
		ResourceMap: map[string]*schema.Table{
			"kms.keys":                 KmsKeyring(),
			"compute.addresses":        ComputeAddresses(),
			"compute.autoscalers":      ComputeAutoscalers(),
			"compute.backend_services": ComputeBackendServices(),
			"compute.disk_types":       ComputeDiskTypes(),
			"compute.images":           ComputeImages(),
			"compute.instances":        ComputeInstances(),
			"compute.interconnects":    ComputeInterconnects(),
			"compute.networks":         ComputeNetworks(),
			"compute.disks":            ComputeDisks(),
			"compute.ssl_certificates": ComputeSslCertificates(),
			"compute.vpn_gateways":     ComputeVpnGateways(),
			"compute.subnetworks":      ComputeSubnetworks(),
			"compute.firewalls":        ComputeFirewalls(),
			"compute.forwarding_rules": ComputeForwardingRules(),
			"cloudfunctions.functions": CloudfunctionsFunction(),
			"iam.project_roles":        IamRoles(),
			"iam.service_accounts":     IamServiceAccounts(),
			"storage.buckets":          StorageBucket(),
			"sql.instances":            SQLInstances(),
			"domains.registrations":    DomainsRegistration(),
			"crm.projects":             CrmProjects(),
		},
		Config: func() interface{} {
			return &client.Config{}
		},
		DefaultConfigGenerator: func() (string, error) {
			return client.DefaultConfig, nil
		},
	}

}
