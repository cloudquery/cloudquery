CREATE TABLE IF NOT EXISTS public.azure_network_virtual_network_ip_allocations
(
    cq_id uuid NULL,
    meta jsonb NULL,
    virtual_network_cq_id uuid NOT NULL,
    id TEXT NOT NULL,
    CONSTRAINT azure_network_virtual_network_ip_allocations_cq_id_key UNIQUE (cq_id),
    CONSTRAINT azure_network_virtual_network_ip_allocations_pk PRIMARY KEY (virtual_network_cq_id, id)
);
-- public.azure_network_virtual_network_ip_allocations foreign keys;
ALTER TABLE public.azure_network_virtual_network_ip_allocations
    ADD CONSTRAINT azure_network_virtual_network_ip_all_virtual_network_cq_id_fkey FOREIGN KEY (virtual_network_cq_id) REFERENCES public.azure_network_virtual_networks (cq_id) ON
        DELETE CASCADE;

INSERT INTO azure_network_virtual_network_ip_allocations(cq_id, virtual_network_cq_id, id)
SELECT gen_random_uuid(),
       cq_id,
       UNNEST(ip_allocations)
FROM azure_network_virtual_networks;

ALTER TABLE IF EXISTS azure_network_virtual_networks DROP COLUMN ip_allocations;

ALTER TABLE IF EXISTS azure_network_virtual_networks
    ALTER
        COLUMN dhcp_options_dns_servers TYPE _text;