resource "kubernetes_resource_quota" "example" {
  metadata {
    name = "resource-quota${var.test_prefix}${var.test_suffix}"
  }
  spec {
    hard = {
      pods = 10
    }

    scopes = ["BestEffort"]
  }
}
