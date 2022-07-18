\echo "Executing CIS V1.3.0 Section 7"
\set check_id "7.1"
\echo "Executing check 7.1"
\i queries/compute/vms_utilizing_managed_disks.sql
\set check_id "7.2"
\echo "Executing check 7.2"
\i queries/compute/os_and_data_disks_encrypted_with_cmk.sql
\set check_id "7.3"
\echo "Executing check 7.3"
\i queries/compute/unattached_disks_are_encrypted_with_cmk.sql
\set check_id "7.4"
\echo "Executing check 7.4"
\echo "Check must be done manually"
\set check_id "7.5"
\echo "Executing check 7.5"
\echo "Check must be done manually"
\set check_id "7.6"
\echo "Executing check 7.6"
\echo "Check must be done manually"
\set check_id "7.7"
\echo "Executing check 7.7"
\i queries/compute/vhds_not_encrypted.sql