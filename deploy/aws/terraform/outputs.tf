output "rds_cluster_master_password" {
  description = "The master password"
  value       = random_password.password.result
  sensitive   = true
}