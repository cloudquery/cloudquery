resource "digitalocean_custom_image" "do_image" {
  name    = "do_image${random_id.test_id.hex}"
  url     = "https://mirror.pkgbuild.com/images/latest/Arch-Linux-x86_64-cloudimg-20210815.31636.qcow2"
  regions = ["nyc3"]
}