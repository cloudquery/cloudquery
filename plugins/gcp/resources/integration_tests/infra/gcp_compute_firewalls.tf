resource "google_compute_firewall" "google_compute_firewalls_firewall_a" {
  name    = "google-compute-firewalls-firewall-a-${var.test_suffix}"
  network = google_compute_network.network.name


  allow {
    protocol = "tcp"
    ports = [
      "80",
      "22",
      "8080",
    "1000-2000"]
  }

  source_tags = [
  "web"]
}

resource "google_compute_firewall" "google_compute_firewalls_firewall_d" {
  name    = "google-compute-firewalls-firewall-d-${var.test_suffix}"
  network = google_compute_network.network.name

  deny {
    protocol = "tcp"
    ports = [
      "123",
    ]
  }


  source_tags = [
  "web"]
}
