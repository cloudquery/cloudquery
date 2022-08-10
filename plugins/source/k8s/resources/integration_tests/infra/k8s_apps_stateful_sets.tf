// https://github.com/skydome/terraform-kubernetes-mongodb/blob/master/main.tf

resource "kubernetes_stateful_set" "mongodb" {
  metadata {
    name      = "stateful-set${var.test_prefix}${var.test_suffix}"
    labels = {
      app = "local-scheduler"
      service = "mongodb"
    }
  }
  spec {
    replicas = 1

    selector {
      match_labels = {
        app = "local-scheduler"
        service = "mongodb"
      }
    }

    service_name = "mongodb"

    template {
      metadata {
        labels = {
          app = "local-scheduler"
          service = "mongodb"
        }
      }

      spec {
        volume {
          name = "configdir"
        }

        volume {
          name = "datadir"
        }

        container {
          image   = "mongo:bionic"
          name    = "mongodb"
          command = ["mongod"]
          args    = ["--dbpath=/data/db", "--port=1235", "--bind_ip=0.0.0.0"]

          env {
            name = "EDGE_PORT"
            value = 1235
          }

          port {
            container_port = 1235
          }

          volume_mount {
            name       = "configdir"
            mount_path = "/data/configdb"
          }

          volume_mount {
            name       = "datadir"
            mount_path = "/data/db"
          }
        }
      }
    }
  }
}