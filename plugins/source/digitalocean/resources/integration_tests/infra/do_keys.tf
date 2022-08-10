resource "tls_private_key" "ssh_sky" {
  algorithm = "RSA"
  rsa_bits  = "4096"
}

resource "digitalocean_ssh_key" "do_ssh_key" {
  name       = "do_ssh_key-${random_id.test_id.hex}"
  public_key = tls_private_key.ssh_sky.public_key_openssh
}