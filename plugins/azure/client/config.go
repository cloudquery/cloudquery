package client

// Provider Configuration

type Config struct {
	Subscriptions []string `yaml:"subscriptions"`
	Resources     []Resource
}

type Resource struct {
	Name  string
	Other map[string]interface{} `yaml:",inline"`
}

const DefaultConfig = `
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
      - name: keyvault.vaults
      - name: network.virtual_networks`
