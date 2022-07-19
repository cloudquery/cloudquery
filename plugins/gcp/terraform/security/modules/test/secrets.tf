data "google_project" "project" {}

resource "google_secret_manager_secret" "secret-basic" {
  secret_id = "${var.prefix}-secret"

  replication {
    user_managed {
      replicas {
        location = "us-east1"
        customer_managed_encryption {
          kms_key_name = google_kms_crypto_key.key-us-east1.id
        }
      }
    }
  }

  rotation {
    next_rotation_time = "2050-01-01T00:00:01Z"
    rotation_period = "3153600000s"
  }

  topics {
    name = google_pubsub_topic.pubsub-basic.id
  }

  depends_on = [google_pubsub_topic_iam_member.binding, google_kms_crypto_key_iam_member.crypto_key]
}

resource "google_pubsub_topic_iam_member" "binding" {
  project = var.project_id
  topic = google_pubsub_topic.pubsub-basic.id
  role = "roles/pubsub.publisher"
  member = "serviceAccount:service-${data.google_project.project.number}@gcp-sa-secretmanager.iam.gserviceaccount.com"
}

resource "google_pubsub_topic" "pubsub-basic" {
  name = "${var.prefix}-secrets-topic"
}

resource "google_kms_key_ring" "keyring-us-east1" {
  name     = "${var.prefix}-secrets-keyring"
  location = "us-east1"
}

resource "google_kms_crypto_key" "key-us-east1" {
  name            = "${var.prefix}-secrets-cryptokey"
  key_ring        = google_kms_key_ring.keyring-us-east1.id
}

resource "google_kms_crypto_key_iam_member" "crypto_key" {
  crypto_key_id = google_kms_crypto_key.key-us-east1.id
  role          = "roles/cloudkms.cryptoKeyEncrypterDecrypter"
  member        = "serviceAccount:service-${data.google_project.project.number}@gcp-sa-secretmanager.iam.gserviceaccount.com"
}