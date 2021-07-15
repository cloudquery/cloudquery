
# Table: aws_ec2_instance_elastic_gpu_associations
Describes the association between an instance and an Elastic Graphics accelerator.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|instance_cq_id|uuid|Unique CloudQuery ID of aws_ec2_instances table (FK)|
|elastic_gpu_association_id|text|The ID of the association.|
|elastic_gpu_association_state|text|The state of the association between the instance and the Elastic Graphics accelerator.|
|elastic_gpu_association_time|text|The time the Elastic Graphics accelerator was associated with the instance.|
|elastic_gpu_id|text|The ID of the Elastic Graphics accelerator.|
