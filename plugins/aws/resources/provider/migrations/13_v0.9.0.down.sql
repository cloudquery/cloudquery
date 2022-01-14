CREATE TABLE IF NOT EXISTS public.aws_ec2_security_group_ip_permission_ipv6_ranges
(
    cq_id uuid,
    meta jsonb,
    security_group_ip_permission_cq_id uuid NOT NULL,
    cidr_ipv6 text COLLATE pg_catalog."default" NOT NULL,
    description text COLLATE pg_catalog."default",
    CONSTRAINT aws_ec2_security_group_ip_permission_ipv6_ranges_pk PRIMARY KEY (security_group_ip_permission_cq_id, cidr_ipv6),
    CONSTRAINT aws_ec2_security_group_ip_permission_ipv6_ranges_cq_id_key UNIQUE (cq_id),
    CONSTRAINT aws_ec2_security_group_ip_pe_security_group_ip_permission_fkey1 FOREIGN KEY (security_group_ip_permission_cq_id)
        REFERENCES public.aws_ec2_security_group_ip_permissions (cq_id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE CASCADE
)

CREATE TABLE IF NOT EXISTS public.aws_ec2_security_group_ip_permissions_egress_ip_ranges
(
    cq_id uuid,
    meta jsonb,
    security_group_ip_permissions_egress_cq_id uuid NOT NULL,
    cidr_ip text COLLATE pg_catalog."default" NOT NULL,
    description text COLLATE pg_catalog."default",
    CONSTRAINT aws_ec2_security_group_ip_permissions_egress_ip_ranges_pk PRIMARY KEY (security_group_ip_permissions_egress_cq_id, cidr_ip),
    CONSTRAINT aws_ec2_security_group_ip_permissions_egress_ip_range_cq_id_key UNIQUE (cq_id),
    CONSTRAINT aws_ec2_security_group_ip_pe_security_group_ip_permission_fkey4 FOREIGN KEY (security_group_ip_permissions_egress_cq_id)
        REFERENCES public.aws_ec2_security_group_ip_permissions_egresses (cq_id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE CASCADE
)

CREATE TABLE IF NOT EXISTS public.aws_ec2_security_group_ip_permissions_egress_ipv6_ranges
(
    cq_id uuid,
    meta jsonb,
    security_group_ip_permissions_egress_cq_id uuid NOT NULL,
    cidr_ipv6 text COLLATE pg_catalog."default" NOT NULL,
    description text COLLATE pg_catalog."default",
    CONSTRAINT aws_ec2_security_group_ip_permissions_egress_ipv6_ranges_pk PRIMARY KEY (security_group_ip_permissions_egress_cq_id, cidr_ipv6),
    CONSTRAINT aws_ec2_security_group_ip_permissions_egress_ipv6_ran_cq_id_key UNIQUE (cq_id),
    CONSTRAINT aws_ec2_security_group_ip_pe_security_group_ip_permission_fkey5 FOREIGN KEY (security_group_ip_permissions_egress_cq_id)
        REFERENCES public.aws_ec2_security_group_ip_permissions_egresses (cq_id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE CASCADE
)

CREATE TABLE IF NOT EXISTS public.aws_ec2_security_group_ip_permissions_egress_prefix_list_ids
(
    cq_id uuid NOT NULL,
    meta jsonb,
    security_group_ip_permissions_egress_cq_id uuid,
    description text COLLATE pg_catalog."default",
    prefix_list_id text COLLATE pg_catalog."default",
    CONSTRAINT aws_ec2_security_group_ip_permissions_egress_prefix_list_ids_pk PRIMARY KEY (cq_id),
    CONSTRAINT aws_ec2_security_group_ip_pe_security_group_ip_permission_fkey6 FOREIGN KEY (security_group_ip_permissions_egress_cq_id)
        REFERENCES public.aws_ec2_security_group_ip_permissions_egresses (cq_id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE CASCADE
)

CREATE TABLE IF NOT EXISTS public.aws_ec2_security_group_ip_permissions_egress_user_group_pairs
(
    cq_id uuid,
    meta jsonb,
    security_group_ip_permissions_egress_cq_id uuid NOT NULL,
    description text COLLATE pg_catalog."default",
    group_id text COLLATE pg_catalog."default" NOT NULL,
    group_name text COLLATE pg_catalog."default",
    peering_status text COLLATE pg_catalog."default",
    user_id text COLLATE pg_catalog."default" NOT NULL,
    vpc_id text COLLATE pg_catalog."default",
    vpc_peering_connection_id text COLLATE pg_catalog."default",
    CONSTRAINT aws_ec2_security_group_ip_permissions_egress_user_group_pair_pk PRIMARY KEY (security_group_ip_permissions_egress_cq_id, group_id, user_id),
    CONSTRAINT aws_ec2_security_group_ip_permissions_egress_user_gro_cq_id_key UNIQUE (cq_id),
    CONSTRAINT aws_ec2_security_group_ip_pe_security_group_ip_permission_fkey7 FOREIGN KEY (security_group_ip_permissions_egress_cq_id)
        REFERENCES public.aws_ec2_security_group_ip_permissions_egresses (cq_id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE CASCADE
)

CREATE TABLE IF NOT EXISTS public.aws_ec2_security_group_ip_permissions_egresses
(
    cq_id uuid NOT NULL,
    meta jsonb,
    security_group_cq_id uuid,
    from_port integer,
    ip_protocol text COLLATE pg_catalog."default",
    to_port integer,
    CONSTRAINT aws_ec2_security_group_ip_permissions_egresses_pk PRIMARY KEY (cq_id),
    CONSTRAINT aws_ec2_security_group_ip_permission_security_group_cq_id_fkey1 FOREIGN KEY (security_group_cq_id)
        REFERENCES public.aws_ec2_security_groups (cq_id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE CASCADE
)

ALTER TABLE IF EXISTS "aws_ec2_security_group_ip_permissions" DROP COLUMN "permission_type";

ALTER TABLE IF EXISTS "aws_ec2_security_group_ip_permission_ip_ranges" RENAME COLUMN "cidr" TO "cidr_ip";
ALTER TABLE IF EXISTS "aws_ec2_security_group_ip_permission_ip_ranges" DROP COLUMN "cidr_type";
ALTER TABLE IF EXISTS "public"."aws_ec2_security_group_ip_permission_ip_ranges"
    ADD CONSTRAINT "aws_ec2_security_group_ip_permission_ip_ranges_pk" PRIMARY KEY (security_group_ip_permission_cq_id, cidr_ip);
