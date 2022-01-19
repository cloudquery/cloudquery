--make azure_network_public_ip_address_ip_tags table a json column of azure_network_public_ip_addresses

ALTER TABLE IF EXISTS azure_network_public_ip_addresses
    ADD COLUMN ip_tags JSON;


UPDATE azure_network_public_ip_addresses ips
SET ip_tags =
        (SELECT JSON_OBJECT_AGG(ip_tag_type, tag)
         FROM azure_network_public_ip_address_ip_tags
         WHERE public_ip_address_cq_id = ips.cq_id);


DROP TABLE IF EXISTS azure_network_public_ip_address_ip_tags;

--ip configuration columns of azure_network_public_ip_addresses

ALTER TABLE IF EXISTS azure_network_public_ip_addresses
    DROP COLUMN private_ip_address;


ALTER TABLE IF EXISTS azure_network_public_ip_addresses
    DROP COLUMN private_ip_allocation_method;


ALTER TABLE IF EXISTS azure_network_public_ip_addresses
    DROP COLUMN subnet;


ALTER TABLE IF EXISTS azure_network_public_ip_addresses
    DROP COLUMN public_ip_address;


ALTER TABLE IF EXISTS azure_network_public_ip_addresses
    ADD COLUMN ip_configuration JSON;


ALTER TABLE IF EXISTS azure_network_public_ip_addresses
    ADD COLUMN service_public_ip_address JSON;


ALTER TABLE IF EXISTS azure_network_public_ip_addresses
    ADD COLUMN nat_gateway JSON;


ALTER TABLE IF EXISTS azure_network_public_ip_addresses
    ADD COLUMN linked_public_ip_address JSON;

--change ip_address column of azure_network_public_ip_addresses from text to inet

ALTER TABLE azure_network_public_ip_addresses
    ALTER COLUMN ip_address TYPE INET USING ip_address::INET;

ALTER TABLE IF EXISTS azure_network_virtual_networks
    ADD COLUMN IF NOT EXISTS ip_allocations _text;

UPDATE azure_network_virtual_networks n
SET
    ip_allocations = (SELECT array_agg(id) AS resources
                      FROM azure_network_virtual_network_ip_allocations anvnia
                      WHERE virtual_network_cq_id = n.cq_id);

DROP TABLE IF EXISTS azure_network_virtual_network_ip_allocations;

ALTER TABLE IF EXISTS azure_network_virtual_networks
    ALTER COLUMN dhcp_options_dns_servers TYPE _inet
        USING dhcp_options_dns_servers::inet[];
