
# Table: gcp_serviceusage_service_apis
Api is a light-weight descriptor for an API Interface Interfaces are also described as "protocol buffer services" in some contexts, such as by the "service" keyword in a proto file, but they are different from API Services, which represent a concrete implementation of an interface as opposed to simply a description of methods and bindings
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|service_cq_id|uuid|Unique CloudQuery ID of gcp_serviceusage_services table (FK)|
|methods|jsonb|The methods of this interface, in unspecified order|
|mixins|jsonb|Included interfaces|
|name|text|The fully qualified name of this interface, including package name followed by the interface's simple name|
|options|jsonb|Any metadata attached to the interface|
|source_context_file_name|text|The path-qualified name of the proto file that contained the associated protobuf element|
|syntax|text|"SYNTAX_PROTO2" - Syntax `proto2`   "SYNTAX_PROTO3" - Syntax `proto3`|
|version|text|A version string for this interface|
