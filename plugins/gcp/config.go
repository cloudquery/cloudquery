package main

var configYaml = `
  - name: gcp
#    project_filter: "" # Optional. Filter as described https://cloud.google.com/sdk/gcloud/reference/projects/list --filter
#    project_ids: # Optional. If not specified either using all projects accessible.
#     - <CHANGE_THIS_TO_YOUR_PROJECT_ID>
    resources:
      - name: compute.addresses
      - name: compute.autoscalers
      - name: compute.backend_services
      - name: compute.disk_types
      - name: compute.images
      - name: compute.instances
      - name: compute.interconnects
      - name: compute.networks
      - name: compute.ssl_certificates
      - name: compute.vpn_gateways
      - name: compute.subnetworks
      - name: compute.firewalls
      - name: compute.forwarding_rules
      - name: cloudfunctions.functions
      - name: iam.project_roles
      - name: iam.service_accounts
      - name: storage.buckets
      - name: sql.instances
      - name: crm.projects`
