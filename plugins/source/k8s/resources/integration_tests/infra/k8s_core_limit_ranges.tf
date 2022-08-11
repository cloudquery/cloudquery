resource "kubernetes_limit_range" "example" {
  metadata {
    name = "limit-range${var.test_prefix}${var.test_suffix}"
    namespace = kubernetes_namespace.limitrange.metadata.0.name
  }

  spec {
    limit {
      type = "Pod"
      max = {
        cpu    = "200m"
        memory = "1024Mi"
      }
    }
  }
}

resource "kubernetes_namespace" "limitrange" {
  metadata {
    name = "limitrangenamespace${var.test_prefix}${var.test_suffix}"
  }
}
