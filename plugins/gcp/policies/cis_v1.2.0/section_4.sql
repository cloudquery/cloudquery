\set framework 'cis_v1.2.0'
\echo "Executing CIS V1.2.0 Section 4"
\set check_id "4.1"
\echo "Executiong check 4.1"
\i queries/compute/instances_with_default_service_account.sql
\set check_id "4.2"
\echo "Executiong check 4.2"
\i queries/compute/instances_with_default_service_account_with_full_access.sql
\set check_id "4.3"
\echo "Executiong check 4.3"
\i queries/compute/instances_without_block_project_wide_ssh_keys.sql
\set check_id "4.4"
\echo "Executiong check 4.4"
\i queries/compute/oslogin_disabled.sql
\set check_id "4.5"
\echo "Executiong check 4.5"
\i queries/compute/serial_port_connection_enabled.sql
\set check_id "4.6"
\echo "Executiong check 4.6"
\i queries/compute/instance_ip_forwarding_enabled.sql
\set check_id "4.7"
\echo "Executiong check 4.7"
\i queries/compute/disks_encrypted_with_csek.sql
\set check_id "4.8"
\echo "Executiong check 4.8"
\i queries/compute/instances_with_shielded_vm_disabled.sql
\set check_id "4.9"
\echo "Executiong check 4.9"
\i queries/compute/instances_with_public_ip.sql
-- MANUAL
\set check_id "4.10"
\echo "Executing check 4.10"
\echo "Ensure that App Engine applications enforce HTTPS connections (Manual)"
\i queries/manual.sql
\set check_id "4.11"
\echo "Executiong check 4.11"
\i queries/compute/instances_without_confidential_computing.sql
