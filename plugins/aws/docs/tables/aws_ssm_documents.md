
# Table: aws_ssm_documents
Describes a Amazon Web Services Systems Manager document (SSM document).
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The AWS Account ID of the resource.|
|region|text|The AWS Region of the resource.|
|arn|text|The Amazon Resource Name (ARN) of the managed instance.|
|approved_version|text|The version of the document currently approved for use in the organization.|
|attachments_information|jsonb|Details about the document attachments, including names, locations, sizes, and so on.|
|author|text|The user in your organization who created the document.|
|created_date|timestamp without time zone|The date when the document was created.|
|default_version|text|The default version.|
|description|text|A description of the document.|
|display_name|text|The friendly name of the SSM document|
|document_format|text|The document format, either JSON or YAML.|
|document_type|text|The type of document.|
|document_version|text|The document version.|
|hash|text|The Sha256 or Sha1 hash created by the system when the document was created. Sha1 hashes have been deprecated.|
|hash_type|text|The hash type of the document|
|latest_version|text|The latest version of the document.|
|name|text|The name of the SSM document.|
|owner|text|The Amazon Web Services user account that created the document.|
|parameters|jsonb|A description of the parameters for a document.|
|pending_review_version|text|The version of the document that is currently under review.|
|platform_types|text[]|The list of OS platforms compatible with this SSM document.|
|requires|jsonb|A list of SSM documents required by a document.|
|review_status|text|The current status of the review.|
|schema_version|text|The schema version.|
|sha1|text|The SHA1 hash of the document, which you can use for verification.|
|status|text|The status of the SSM document.|
|status_information|text|A message returned by Amazon Web Services Systems Manager that explains the Status value|
|target_type|text|The target type which defines the kinds of resources the document can run on. For example, /AWS::EC2::Instance|
|version_name|text|The version of the artifact associated with the document.|
|review_information|jsonb|Details about the review of a document.|
|tags|jsonb|The tags, or metadata, that have been applied to the document.|
|account_ids|text[]|The account IDs that have permission to use this document|
|account_sharing_info_list|jsonb|A list of Amazon Web Services accounts where the current document is shared and the version shared with each account.|
