
resource "azurerm_resource_group" "container" {
  name     = "${var.prefix}-container"
  location = "East US"
}

resource "azurerm_kubernetes_cluster" "managed_clusters_cluster" {
  name                = "${var.prefix}_aks"
  location            = azurerm_resource_group.container.location
  resource_group_name = azurerm_resource_group.container.name
  dns_prefix          = "${var.prefix}-cq-aks"

  default_node_pool {
    name                = "default"
    enable_auto_scaling = true
    max_count           = 1
    node_count          = 1
    min_count           = 1
    vm_size             = "Standard_B2s"
    node_labels         = { "node-type" = "system" }
    tags = var.tags
  }

  // network_profile {
  //   network_plugin     = "azure"
  //   network_policy     = "azure"
  //   load_balancer_sku  = "standard"
  //   service_cidr       = "172.17.0.0/16"
  //   dns_service_ip     = "172.17.0.12"
  //   docker_bridge_cidr = "172.18.0.12/16"
  // }

  identity {
    type = "SystemAssigned"
  }

  role_based_access_control {
    enabled = true
  }

  tags = var.tags
}



// # Allows Kubernetes to Pull ACR images
// resource "azurerm_role_assignment" "managed_clusters_role_acr" {
//   scope                = azurerm_container_registry.managed_clusters_registry.id
//   role_definition_name = "AcrPull"
//   principal_id         = azurerm_kubernetes_cluster.managed_clusters_cluster.kubelet_identity[0].object_id
// }

// # Creates An Identity to Pod
// resource "azurerm_user_assigned_identity" "managed_clusters_pod_identity_queue_contributor" {
//   resource_group_name = azurerm_resource_group.kubernetes_cluster.name
//   location            = azurerm_resource_group.kubernetes_cluster.location
//   name                = "queuecontributoraksidentity"
// }

// # Allows Kubernetes to Manage Identity Created on AKS Nodes
// resource "azurerm_role_assignment" "managed_clusters_identity_operator" {
//   scope                = azurerm_user_assigned_identity.managed_clusters_pod_identity_queue_contributor.id
//   role_definition_name = "Managed Identity Operator"
//   principal_id         = azurerm_kubernetes_cluster.managed_clusters_cluster.kubelet_identity[0].object_id
// }

// # Allows Kubernetes to Manage VMs on AKS Nodes
// resource "azurerm_role_assignment" "managed_clusters_vm_contributor" {
//   scope                = "/subscriptions/${data.azurerm_subscription.current.subscription_id}/resourcegroups/${azurerm_kubernetes_cluster.managed_clusters_cluster.node_resource_group}"
//   role_definition_name = "Virtual Machine Contributor"
//   principal_id         = azurerm_kubernetes_cluster.managed_clusters_cluster.kubelet_identity[0].object_id
// }
