# Example from https://github.com/terraform-google-modules/terraform-google-vm/blob/master/examples/compute_instance/simple/main.tf

# module "instance_template" {
#   source          = "../../../modules/instance_template"
#   region          = var.region
#   project_id      = var.project_id
#   subnetwork      = var.subnetwork
# }

# module "compute_instance" {
#   source              = "../../../modules/compute_instance"
#   region              = var.region
#   zone                = var.zone
#   subnetwork          = var.subnetwork
#   num_instances       = var.num_instances
#   hostname            = "instance-simple"
#   instance_template   = module.instance_template.self_link
#   deletion_protection = false

#   access_config = [{
#     nat_ip       = var.nat_ip
#     network_tier = var.network_tier
#   }, ]
# }

resource "google_compute_instance" "default" {
  project      = var.project_id
  zone         = "${var.region}-b"
  name         = "${var.prefix}-compute"
  machine_type = "e2-micro"
  boot_disk {
    initialize_params {
      image = "debian-cloud/debian-9"
    }
  }
  network_interface {
    network = "default"
  }
  tags = ["health-check", "ssh"]
}