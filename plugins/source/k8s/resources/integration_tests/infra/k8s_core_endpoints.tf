resource "kubernetes_endpoints" "example" {
  metadata {
    name = "endpoint${var.test_prefix}${var.test_suffix}"
  }

  subset {
    address {
      ip = "10.0.0.4"
    }

    port {
      name     = "http"
      port     = 80
      protocol = "TCP"
    }
  }
}
