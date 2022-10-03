\set check_id 'EFS.1'
\echo "Executing check EFS.1"
\ir ../queries/efs/unencrypted_efs_filesystems.sql

\set check_id 'EFS.2'
\echo "Executing check EFS.2"
\ir ../queries/efs/efs_filesystems_with_disabled_backups.sql
