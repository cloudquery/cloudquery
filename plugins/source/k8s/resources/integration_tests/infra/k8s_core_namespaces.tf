resource "kubernetes_namespace" "example" {
  metadata {
    annotations = {
      name = "namespace"
    }

    labels = {
      mylabel = "label-value"
    }

    name = "namespace${var.test_prefix}${var.test_suffix}"
  }
}