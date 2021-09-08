resource "digitalocean_custom_image" "do_image" {
  name    = "do_image${var.test_prefix}-${var.test_suffix}"
  url     = "https://mirror.pkgbuild.com/images/latest/Arch-Linux-x86_64-cloudimg-20210815.31636.qcow2"
  regions = ["nyc3"]
}