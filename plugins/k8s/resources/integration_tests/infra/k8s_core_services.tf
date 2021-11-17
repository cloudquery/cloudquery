resource "kubernetes_service" "core_services" {
  metadata {
    name = "service${var.test_prefix}${var.test_suffix}"
  }
  spec {
    selector         = {
      app = kubernetes_pod.core_services_pod.metadata.0.labels.app
    }
    session_affinity = "ClientIP"
    port {
      port        = 63000
    }

#    ен = "LoadBalancer"
  }

  depends_on = [kubernetes_pod.core_services_pod]
}

resource "kubernetes_pod" "core_services_pod" {
  metadata {
    name   = "coreservices-pod${var.test_prefix}${var.test_suffix}"
    labels = {
      app = "MyApp"
    }
  }

  spec {
    container {
      image = "nginx:1.20.1"
      name  = "service-pod-${var.test_prefix}${var.test_suffix}"
    }
  }
}