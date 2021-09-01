resource "aws_mq_configuration" "mq_test" {
  description    = "MQ Test Configuration"
  name           = "mq_test_configuration${var.test_suffix}"
  engine_type    = "ActiveMQ"
  engine_version = "5.15.0"

  data = <<DATA
<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<broker xmlns="http://activemq.apache.org/schema/core">
  <plugins>
    <forcePersistencyModeBrokerPlugin persistenceFlag="true"/>
    <statisticsBrokerPlugin/>
    <timeStampingBrokerPlugin ttlCeiling="86400000" zeroExpirationOverride="86400000"/>
  </plugins>
</broker>
DATA

}

resource "aws_mq_broker" "mq_test_broker" {
  broker_name = "mq_test_broker_${var.test_suffix}"

  configuration {
    id       = aws_mq_configuration.mq_test.id
    revision = aws_mq_configuration.mq_test.latest_revision
  }

  engine_type        = "ActiveMQ"
  engine_version     = "5.15.9"
  host_instance_type = "mq.t2.micro"
  security_groups    = [aws_security_group.mq_test_sg.id]

  user {
    username = "ExampleUser"
    password = "someLongPasswordWeDontCare"
  }

  publicly_accessible = false
}

resource "aws_security_group" "mq_test_sg" {
  name        = "mq_test_sg_${var.test_prefix}_${var.test_suffix}"
  description = "Managed By Terraform, AWS e2e Testing Resource"
}