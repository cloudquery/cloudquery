resource "kubernetes_resource_quota" "example" {
  metadata {
    name = "resource-quota${var.test_prefix}${var.test_suffix}"
    namespace = kubernetes_namespace.resourcequota.metadata.0.name
  }

  spec {
    hard = {
      pods = 10
    }

    scopes = ["BestEffort"]
  }
}

resource "kubernetes_namespace" "resourcequota" {
  metadata {
    name = "resourcequotanamespace${var.test_prefix}${var.test_suffix}"
  }
}
