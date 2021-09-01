resource "aws_fsx_backup" "test_fsx_backup" {
  file_system_id = aws_fsx_lustre_file_system.test_fsx.id
  tags = {
    Name = "fsx-${var.test_prefix}${var.test_suffix}"
  }
}

resource "aws_fsx_lustre_file_system" "test_fsx" {
  storage_capacity            = 6000
  storage_type                = "HDD"
  drive_cache_type            = "NONE"
  deployment_type             = "PERSISTENT_1"
  per_unit_storage_throughput = 12
  subnet_ids                  = [aws_subnet.aws_vpc_subnet2.id]

}

