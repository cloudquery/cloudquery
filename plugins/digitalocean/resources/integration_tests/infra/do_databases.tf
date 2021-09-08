resource "digitalocean_database_cluster" "do_cluster" {
  name       = "do-database-cluster-${var.test_prefix}-${var.test_suffix}"
  engine     = "pg"
  version    = "11"
  size       = "db-s-1vcpu-1gb"
  region     = "nyc3"
  node_count = 1
}

resource "digitalocean_database_db" "do_database" {
  cluster_id = digitalocean_database_cluster.do_cluster.id
  name       = "do-database-${var.test_prefix}-${var.test_suffix}"
}

resource "digitalocean_database_user" "do_database_user" {
  cluster_id = digitalocean_database_cluster.do_cluster.id
  name       = "do-database-user-${var.test_prefix}-${var.test_suffix}"
}

resource "digitalocean_database_replica" "do_database_replica" {
  cluster_id = digitalocean_database_cluster.do_cluster.id
  name       = "do-database-replica-${var.test_prefix}-${var.test_suffix}"
  size       = "db-s-1vcpu-1gb"
  region     = "nyc1"

  depends_on = [digitalocean_database_cluster.do_cluster]
}

resource "digitalocean_database_firewall" "do_database_firewall" {
  cluster_id = digitalocean_database_cluster.do_cluster.id

  rule {
    type  = "ip_addr"
    value = "192.168.1.1"
  }

  rule {
    type  = "ip_addr"
    value = "192.0.2.0"
  }
}