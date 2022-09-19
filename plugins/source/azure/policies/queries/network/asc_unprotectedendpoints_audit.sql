WITH network_security_groups AS (
	SELECT DISTINCT
		sr->>'name' AS security_group_name,
		azure_network_interfaces.id AS network_interface_id
	FROM
		azure_network_security_groups,
		JSON_ARRAY_ELEMENTS(azure_network_security_groups.security_rules) AS sr,
		azure_network_interfaces
	WHERE
		azure_network_interfaces.resource_guid = azure_network_security_groups.resource_guid
		AND sr->>'access' = 'Allow'
		AND sr->>'direction' = 'Inbound'
		AND sr->>'protocol' IN ( 'TCP', '*' )
		AND (
            array(select json_array_elements_text(sr->'sourceAddressPrefixes')) && ARRAY [ '*', '0.0.0.0', '0.0.0.0/0', 'Internet', '<nw>/0', '/0' ] -- TODO CHECK
			OR sr->>'sourceAddressPrefix' IN ( '*', '0.0.0.0', '0.0.0.0/0', 'Internet', '<nw>/0', '/0' )
		)
	)
insert into azure_policy_results
SELECT
  :'execution_time',
  :'framework',
  :'check_id',
  'All network ports should be restricted on network security groups associated to your virtual machine',
    subscription_id,
	machines.id,
	case
    when network_security_groups.security_group_name IS DISTINCT FROM NULL
    then 'fail' else 'pass'
  end
FROM
	azure_compute_virtual_machines machines,
	json_array_elements ( machines.network_profile->'networkInterfaces' ) AS interface
	LEFT JOIN network_security_groups ON interface->>'id' = network_security_groups.network_interface_id -- TODO check match
