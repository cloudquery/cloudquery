################################################################################
# Cloud Run Module
################################################################################

data "google_project" "project" {
}

resource "google_secret_manager_secret" "secret" {
  secret_id = "${lower(var.prefix)}-cloudrun-secret"
  replication {
    automatic = true
  }
}

resource "google_secret_manager_secret_version" "secret-version-data" {
  secret = google_secret_manager_secret.secret.name
  secret_data = "secret-data"
}

resource "google_secret_manager_secret_iam_member" "secret-access" {
  secret_id = google_secret_manager_secret.secret.id
  role      = "roles/secretmanager.secretAccessor"
  member    = "serviceAccount:${data.google_project.project.number}-compute@developer.gserviceaccount.com"
  depends_on = [google_secret_manager_secret.secret]
}

resource "google_cloud_run_service" "default" {
  name     = "${lower(var.prefix)}-cloudrun-srv"
  location = "us-central1"

  template {
    spec {
      containers {
        image = "us-docker.pkg.dev/cloudrun/container/hello"
        command = ["/server"]
        resources {
          limits = {
            cpu = "1000m"
            memory = "512Mi"
          }
          requests = {}
        }
        volume_mounts {
          name = "a-volume"
          mount_path = "/secrets"
        }
        env {
          name = "ENV_VAR"
          value = "test"
        }
        env {
          name = "SECRET_ENV_VAR"
          value_from {
            secret_key_ref {
              name = google_secret_manager_secret.secret.secret_id
              key = "1"
            }
          }
        }
      }
      volumes {
        name = "a-volume"
        secret {
          secret_name = google_secret_manager_secret.secret.secret_id
          default_mode = 292 # 0444
          items {
            key = "1"
            path = "my-secret"
            mode = 256 # 0400
          }
        }
      }
    }
    metadata {
      annotations = {
        "autoscaling.knative.dev/maxScale"      = "1"
        "run.googleapis.com/client-name"        = "terraform"
      }
      labels = {
        "key" = "value"
      }
    }
  }

  metadata {
    annotations = {
      generated-by = "magic-modules"
      "run.googleapis.com/ingress" = "internal"
    }
    labels = {
      key = "value"
    }
  }

  traffic {
    percent         = 100
    latest_revision = true
  }
  autogenerate_revision_name = true

  depends_on = [google_secret_manager_secret_version.secret-version-data]
}