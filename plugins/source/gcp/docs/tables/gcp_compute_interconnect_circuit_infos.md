
# Table: gcp_compute_interconnect_circuit_infos
Describes a single physical circuit between the Customer and Google CircuitInfo objects are created by Google, so all fields are output only
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|interconnect_cq_id|uuid|Unique ID of gcp_compute_interconnects table (FK)|
|interconnect_id|text||
|customer_demarc_id|text|Customer-side demarc ID for this circuit|
|google_circuit_id|text|Google-assigned unique ID for this circuit Assigned at circuit turn-up|
|google_demarc_id|text|Google-side demarc ID for this circuit Assigned at circuit turn-up and provided by Google to the customer in the LOA|
