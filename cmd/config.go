package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
	"strings"
)

var headerConfig = `providers:`

var awsConfig = `
  - name: aws
#    accounts: # Optional. if you want to assume role to multiple account and fetch data from them
#      - role_arn: arn:aws:iam::966797507899:role/user
#    regions: # Optional. if commented out assumes all regions
#      - us-east-1
#      - us-west-2
#    log_level: debug # Optional. if commented out will enable AWS SDK debug logging. possible values: debug, debug_with_signing, debug_with_http_body, debug_with_request_retries, debug_with_request_error, debug_with_event_stream_body
#    max_retries: # Optional. The maximum number of times that a request will be retried for failures. Defaults to -1, which defers the max retry setting to the service specific configuration. 	
    resources: # You can comment resources your are not interested in for faster fetching.
      - name: autoscaling.launch_configurations
      - name: cloudtrail.trails
      - name: directconnect.gateways
      - name: ec2.customer_gateways
      - name: ec2.flow_logs
      - name: ec2.images
      - name: ec2.instances
      - name: ec2.internet_gateways
      - name: ec2.nat_gateways
      - name: ec2.network_acls
      - name: ec2.route_tables
      - name: ec2.security_groups
      - name: ec2.subnets
      - name: ec2.vpc_peering_connections
      - name: ec2.vpcs
      - name: ecs.clusters
      - name: efs.filesystems
      - name: elasticbeanstalk.environments
      - name: elbv2.load_balancers
      - name: emr.clusters
      - name: fsx.backups
      - name: iam.groups
      - name: iam.password_policies
      - name: iam.policies
      - name: iam.roles
      - name: iam.users
      - name: kms.keys
      - name: rds.certificates
      - name: rds.clusters
      - name: rds.db_subnet_groups
      - name: redshift.clusters
      - name: redshift.cluster_subnet_groups
      - name: s3.buckets`

var gcpConfig = `
  - name: gcp
    project_id: <CHANGE_THIS_TO_YOUR_PROJECT_ID>
    resources:
      - name: compute.instances
      - name: compute.autoscalers
      - name: compute.disk_types
      - name: compute.images
      - name: compute.instances
      - name: compute.interconnects
      - name: compute.ssl_certificates
      - name: compute.vpn_gateways
      - name: iam.project_roles
      - name: iam.service_accounts
      - name: storage.buckets`

var oktaConfig = `
  - name: okta
    domain: https://<CHANGE_THIS_TO_YOUR_OKTA_DOMAIN>.okta.com
    resources:
      - name: users
      - name: applications`

var azureConfig = `
  - name: azure
#    subscriptions: # Optional. if you not specified, cloudquery tries to access all subscriptions available to tenant
#      - "subscription-id"
    resources:
      - name: resources.groups
      - name: sql.servers
      - name: sql.databases
      - name: postgresql.servers
      - name: mysql.servers
      - name: compute.disks
      - name: keyvault.vaults`

var k8sConfig = `
  - name: k8s
    resources:
      - name: pods
      - name: services`

var initialConfigs = map[string]string{
	"aws":   awsConfig,
	"gcp":   gcpConfig,
	"okta":  oktaConfig,
	"azure": azureConfig,
	"k8s": k8sConfig,
}

var validArgs = []string{"aws", "gcp", "okta", "azure", "k8s"}
var configPath = "./config.yml"

var configCmd = &cobra.Command{
	Use:       fmt.Sprintf("config [choose one or more of: %s]", strings.Join(validArgs, ",")),
	Short:     "Generate initial config.yml for fetch command",
	ValidArgs: validArgs,
	Args:      cobra.RangeArgs(1, len(validArgs)),
	RunE: func(cmd *cobra.Command, args []string) error {
		err := cobra.OnlyValidArgs(cmd, args)
		if err != nil {
			return fmt.Errorf("invalid argument %s for cloudquery gen config. choose from %v", args[0], validArgs)
		}
		var s strings.Builder
		_, err = s.WriteString(headerConfig)
		if err != nil {
			return err
		}
		for _, provider := range args {
			s.WriteString(initialConfigs[provider])
		}
		s.WriteString("\n")
		if _, err := os.Stat(configPath); err == nil {
			return fmt.Errorf("file %s already exists. Either delete it or specify other path via --path flag", configPath)
		} else if os.IsNotExist(err) {
			return ioutil.WriteFile(configPath, []byte(s.String()), 0644)
		} else {
			return err
		}
	},
}

func init() {
	genCmd.AddCommand(configCmd)
	configCmd.Flags().StringVar(&configPath, "path", configPath, "path to output generated config file")
}
