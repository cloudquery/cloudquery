
data "aws_availability_zones" "available" {}
resource "aws_glue_connection" "aws_glue_connection" {
  connection_properties = {
    JDBC_CONNECTION_URL = "jdbc:mysql://example.com/exampledatabase"
    PASSWORD            = "examplepassword"
    USERNAME            = "exampleusername"
  }

  name =  "${var.prefix}-glue-connection"

  physical_connection_requirements {
    availability_zone      =   data.aws_availability_zones.available.names[0]
    security_group_id_list = [ module.vpc.default_security_group_id]
    subnet_id              = module.vpc.private_subnets[0]
  }
}


