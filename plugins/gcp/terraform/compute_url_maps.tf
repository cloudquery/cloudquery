################################################################################
# Compute Module - Url Maps
################################################################################

resource "google_compute_url_map" "url_map" {
  name        = "${local.prefix}-urlmap"

  default_service = google_compute_backend_service.internal-backend-service.id

  host_rule {
    hosts = [
      local.domain
    ]
    path_matcher = "root"
  }

  host_rule {
    hosts = [
      "beta.${local.domain}",
    ]
    path_matcher = "secondary"
  }

  path_matcher {
    name            = "root"
    default_service = google_compute_backend_service.internal-backend-service.id

    path_rule {
      paths = [
        "/home"
      ]
      route_action {
        weighted_backend_services {
          backend_service = google_compute_backend_service.internal-backend-service.id
          weight = 400
          header_action {
            request_headers_to_remove = ["RequestHeaderToRemove"]
            request_headers_to_add {
              header_name = "RequestHeaderToAdd"
              header_value = "RequestHeaderToAddValue"
              replace = true
            }
            response_headers_to_remove = ["ResponseHeaderToRemove"]
            response_headers_to_add {
              header_name = "ResponseHeaderToAdd"
              header_value = "ResponseHeaderToAddValue"
              replace = false
            }
          }
        }
      }
    }

    path_rule {
      paths = [
        "/login"
      ]
      service = google_compute_backend_service.internal-backend-service.id
    }

    path_rule {
      paths = [
        "/static"
      ]
      service = google_compute_backend_service.internal-backend-service.id
    }
  }

  path_matcher {
    name            = "secondary"
    default_service = google_compute_backend_service.internal-backend-service.id
  }

  test {
    service = google_compute_backend_service.internal-backend-service.id
    host    = "beta.${local.domain}"
    path    = "/healthcheck"
  }

  depends_on = [
    google_compute_backend_service.internal-backend-service
  ]
}

resource "google_compute_url_map" "external_url_map" {
  name        = "${local.prefix}-external-urlmap"

  default_service = google_compute_backend_service.backend-service.id

  host_rule {
    hosts = [
      "ex.${local.domain}",
    ]
    path_matcher = "external-1"
  }

  path_matcher {
    name            = "external-1"
    default_service = google_compute_backend_service.backend-service.id

    path_rule {
      paths = [
        "/home"
      ]
      service = google_compute_backend_service.backend-service.id
    }
  }

  depends_on = [
    google_compute_backend_service.internal-backend-service
  ]
}