resource "kubernetes_service_account" "example" {
  metadata {
    name = "service-account${var.test_prefix}${var.test_suffix}"
  }

  secret {
    name = kubernetes_secret.example.metadata[0].name
  }
}

resource "kubernetes_secret" "example" {
  metadata {
    name = "secret${var.test_prefix}${var.test_suffix}"
  }
}
