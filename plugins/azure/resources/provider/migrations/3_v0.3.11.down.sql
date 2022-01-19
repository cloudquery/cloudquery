--make json column of azure_network_public_ip_addresses a table azure_network_public_ip_address_ip_tags

CREATE TABLE IF NOT EXISTS public.azure_network_public_ip_address_ip_tags
(
    cq_id                   UUID  NOT NULL,
    meta                    JSONB NULL,
    public_ip_address_cq_id UUID  NULL,
    ip_tag_type             TEXT  NULL,
    tag                     TEXT  NULL,
    CONSTRAINT azure_network_public_ip_address_ip_tags_pk PRIMARY KEY (cq_id)
);


ALTER TABLE public.azure_network_public_ip_address_ip_tags
    ADD CONSTRAINT azure_network_public_ip_address_ip_public_ip_address_cq_id_fkey
        FOREIGN KEY (public_ip_address_cq_id) REFERENCES public.azure_network_public_ip_addresses (cq_id) ON
            DELETE CASCADE;


INSERT INTO azure_network_public_ip_address_ip_tags(cq_id, public_ip_address_cq_id, ip_tag_type, tag)
SELECT gen_random_uuid(),
       cq_id,
       json_data.key,
       json_data.value
FROM azure_network_public_ip_addresses,
     json_each_text(ip_tags) AS json_data;


ALTER TABLE IF EXISTS azure_network_public_ip_addresses
    DROP COLUMN ip_tags;

--ip configuration columns of azure_network_public_ip_addresses

ALTER TABLE IF EXISTS azure_network_public_ip_addresses
    ADD COLUMN private_ip_address TEXT;

ALTER TABLE IF EXISTS azure_network_public_ip_addresses
    ADD COLUMN private_ip_allocation_method TEXT;


ALTER TABLE IF EXISTS azure_network_public_ip_addresses
    ADD COLUMN subnet JSON;


ALTER TABLE IF EXISTS azure_network_public_ip_addresses
    ADD COLUMN public_ip_address JSON;


ALTER TABLE IF EXISTS azure_network_public_ip_addresses
    DROP COLUMN ip_configuration;


ALTER TABLE IF EXISTS azure_network_public_ip_addresses
    DROP COLUMN service_public_ip_address;


ALTER TABLE IF EXISTS azure_network_public_ip_addresses
    DROP COLUMN nat_gateway;


ALTER TABLE IF EXISTS azure_network_public_ip_addresses
    DROP COLUMN linked_public_ip_address;

--change ip_address column of azure_network_public_ip_addresses from cidr to text

ALTER TABLE azure_network_public_ip_addresses
    ALTER COLUMN ip_address TYPE TEXT;
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