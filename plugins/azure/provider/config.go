package provider

var configYaml = `
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
