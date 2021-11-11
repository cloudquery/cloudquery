
# Table: aws_route53_domains
The domain names registered with Amazon Route 53.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The AWS Account ID of the resource.|
|admin_contact_address_line1|text|First line of the contact's address.|
|admin_contact_address_line2|text|Second line of contact's address, if any.|
|admin_contact_city|text|The city of the contact's address.|
|admin_contact_type|text|Indicates whether the contact is a person, company, association, or public organization.|
|admin_contact_country_code|text|Code for the country of the contact's address.|
|admin_contact_email|text|Email address of the contact.|
|admin_contact_fax|text|Fax number of the contact.|
|admin_contact_first_name|text|First name of contact.|
|admin_contact_last_name|text|Last name of contact.|
|admin_contact_organization_name|text|Name of the organization for contact types other than PERSON.|
|admin_contact_phone_number|text|The phone number of the contact.|
|admin_contact_state|text|The state or province of the contact's city.|
|admin_contact_zip_code|text|The zip or postal code of the contact's address.|
|admin_contact_extra_params|jsonb|A mapping of name to value parameter pairs required by certain top-level domains.|
|domain_name|text|The name of a domain.|
|registrant_contact_address_line1|text|First line of the contact's address.|
|registrant_contact_address_line2|text|Second line of contact's address, if any.|
|registrant_contact_city|text|The city of the contact's address.|
|registrant_contact_type|text|Indicates whether the contact is a person, company, association, or public organization.|
|registrant_contact_country_code|text|Code for the country of the contact's address.|
|registrant_contact_email|text|Email address of the contact.|
|registrant_contact_fax|text|Fax number of the contact.|
|registrant_contact_first_name|text|First name of contact.|
|registrant_contact_last_name|text|Last name of contact.|
|registrant_contact_organization_name|text|Name of the organization for contact types other than PERSON.|
|registrant_contact_phone_number|text|The phone number of the contact.|
|registrant_contact_state|text|The state or province of the contact's city.|
|registrant_contact_zip_code|text|The zip or postal code of the contact's address.|
|registrant_contact_extra_params|jsonb|A mapping of name to value parameter pairs required by certain top-level domains.|
|tech_contact_address_line1|text|First line of the contact's address.|
|tech_contact_address_line2|text|Second line of contact's address, if any.|
|tech_contact_city|text|The city of the contact's address.|
|tech_contact_type|text|Indicates whether the contact is a person, company, association, or public organization.|
|tech_contact_country_code|text|Code for the country of the contact's address.|
|tech_contact_email|text|Email address of the contact.|
|tech_contact_fax|text|Fax number of the contact.|
|tech_contact_first_name|text|First name of contact.|
|tech_contact_last_name|text|Last name of contact.|
|tech_contact_organization_name|text|Name of the organization for contact types other than PERSON.|
|tech_contact_phone_number|text|The phone number of the contact.|
|tech_contact_state|text|The state or province of the contact's city.|
|tech_contact_zip_code|text|The zip or postal code of the contact's address.|
|tech_contact_extra_params|jsonb|A mapping of name to value parameter pairs required by certain top-level domains.|
|abuse_contact_email|text|Email address to contact to report incorrect contact information for a domain, to report that the domain is being used to send spam, to report that someone is cybersquatting on a domain name, or report some other type of abuse.|
|abuse_contact_phone|text|Phone number for reporting abuse.|
|admin_privacy|boolean|Specifies whether contact information is concealed from WHOIS queries|
|auto_renew|boolean|Specifies whether the domain registration is set to renew automatically.|
|creation_date|timestamp without time zone|The date when the domain was created as found in the response to a WHOIS query.|
|dns_sec|text|Reserved for future use.|
|expiration_date|timestamp without time zone|The date when the registration for the domain is set to expire|
|registrant_privacy|boolean|Specifies whether contact information is concealed from WHOIS queries|
|registrar_name|text|Name of the registrar of the domain as identified in the registry|
|registrar_url|text|Web address of the registrar.|
|registry_domain_id|text|Reserved for future use.|
|reseller|text|Reseller of the domain|
|status_list|text[]|An array of domain name status codes, also known as Extensible Provisioning Protocol (EPP) status codes|
|tech_privacy|boolean|Specifies whether contact information is concealed from WHOIS queries|
|updated_date|timestamp without time zone|The last updated date of the domain as found in the response to a WHOIS query.|
|who_is_server|text|The fully qualified name of the WHOIS server that can answer the WHOIS query for the domain.|
|tags|jsonb|A list of tags|
