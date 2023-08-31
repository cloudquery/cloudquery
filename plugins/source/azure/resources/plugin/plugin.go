package plugin

import (
	"github.com/cloudquery/plugin-sdk/v4/caser"
	"github.com/cloudquery/plugin-sdk/v4/docs"
	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"golang.org/x/exp/maps"
)

var (
	Version = "development"
)

var azureExceptions = map[string]string{
	"analysisservices":          "Analysis Services",
	"apimanagement":             "API Management",
	"appcomplianceautomation":   "App Compliance Automation",
	"applicationinsights":       "Application Insights",
	"appservice":                "App Service",
	"azurearcdata":              "Azure Arc Data",
	"bgp":                       "Border Gateway Protocol (BGP)",
	"cdn":                       "Content Delivery Network (CDN)",
	"cognitiveservices":         "Cognitive Services",
	"connectedvmware":           "Connected VMware",
	"containerinstance":         "Container Instance",
	"containerregistry":         "Container Registry",
	"containerservice":          "Container Service",
	"costmanagement":            "Cost Management",
	"customerinsights":          "Customer Insights",
	"databox":                   "Data Box",
	"datafactory":               "Data Factory",
	"datalakeanalytics":         "Data Lake Analytics",
	"datalakestore":             "Data Lake Store",
	"ddos":                      "Distributed Denial of Service (DDoS)",
	"desktopvirtualization":     "Desktop Virtualization",
	"devops":                    "DevOps",
	"dnsresolver":               "DNS Resolver",
	"dscp":                      "Differentiated Services Code Point (DSCP)",
	"engagementfabric":          "Engagement Fabric",
	"eventhub":                  "Event Hub",
	"hanaonazure":               "HANA on Azure",
	"hdinsight":                 "HDInsight",
	"healthcareapis":            "Healthcare APIs",
	"hybridcompute":             "Hybrid Compute",
	"hybriddatamanager":         "Hybrid Data Manager",
	"ip":                        "IP",
	"keyvault":                  "Key Vault",
	"mariadb":                   "MariaDB",
	"mysql":                     "MySQL",
	"nat":                       "Network Address Translation (NAT)",
	"networkfunction":           "Network Function",
	"operationalinsights":       "Operational Insights",
	"postgresql":                "PostgreSQL",
	"postgresqlflexibleservers": "PostgreSQL Flexible Servers",
	"postgresqlhsc":             "PostgreSQL Hyperscale (Citus)",
	"powerbidedicated":          "Power BI Dedicated",
	"saas":                      "Software as a Service (SaaS)",
	"sap":                       "SAP",
	"servicebus":                "Service Bus",
	"sqlvirtualmachine":         "SQL Virtual Machine",
	"streamanalytics":           "Stream Analytics",
	"vpn":                       "Virtual Private Network (VPN)",
	"wans":                      "Wide Area Networks (WANs)",
	"windowsiot":                "Windows IoT",
}

func titleTransformer(table *schema.Table) error {
	if table.Title != "" {
		return nil
	}

	exceptions := maps.Clone(docs.DefaultTitleExceptions)
	for k, v := range azureExceptions {
		exceptions[k] = v
	}
	csr := caser.New(caser.WithCustomExceptions(exceptions))
	table.Title = csr.ToTitle(table.Name)
	return nil
}

func Plugin() *plugin.Plugin {
	return plugin.NewPlugin(
		"azure",
		Version,
		NewClient,
	)
}
