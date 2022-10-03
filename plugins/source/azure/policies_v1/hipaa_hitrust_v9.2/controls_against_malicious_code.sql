\echo "Controls Against Malicious Code"
\set check_id '0201.09j1Organizational.124 - 09.j - 2'
\echo :check_id
\ir ../queries/compute/vmantimalwareextension_deploy.sql
\set check_id '0201.09j1Organizational.124 - 09.j - 3'
\echo :check_id
\ir ../queries/compute/endpoint_protection_solution_should_be_installed_on_virtual_machine_scale_sets.sql
\set check_id '0201.09j1Organizational.124 - 09.j - 4'
\echo :check_id
\ir ../queries/compute/virtualmachines_antimalwareautoupdate_auditifnotexists.sql
\set check_id '0201.09j1Organizational.124 - 09.j - 6'
\echo :check_id
\ir ../queries/compute/asc_missingsystemupdates_audit.sql
