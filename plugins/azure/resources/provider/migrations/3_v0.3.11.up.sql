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
