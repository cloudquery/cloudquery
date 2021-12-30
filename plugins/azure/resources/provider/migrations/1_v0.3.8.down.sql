ALTER TABLE IF EXISTS azure_network_virtual_network_subnets
    RENAME TO azure_networks_virtual_network_subnets;
ALTER TABLE IF EXISTS azure_network_virtual_network_peerings
    RENAME TO azure_networks_virtual_network_peerings;
ALTER TABLE IF EXISTS azure_network_virtual_network_ip_allocations
    RENAME TO azure_networks_virtual_network_ip_allocations;