
# Table: aws_ec2_instance_elastic_inference_accelerator_associations
Describes the association between an instance and an elastic inference accelerator.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|instance_cq_id|uuid|Unique CloudQuery ID of aws_ec2_instances table (FK)|
|elastic_inference_accelerator_arn|text|The Amazon Resource Name (ARN) of the elastic inference accelerator.|
|elastic_inference_accelerator_association_id|text|The ID of the association.|
|elastic_inference_accelerator_association_state|text|The state of the elastic inference accelerator.|
|elastic_inference_accelerator_association_time|timestamp without time zone|The time at which the elastic inference accelerator is associated with an instance.|
