ALTER TABLE IF EXISTS azure_compute_virtual_machines
    DROP COLUMN "windows_configuration_patch_settings_assessment_mode";
ALTER TABLE IF EXISTS azure_compute_virtual_machines
    DROP COLUMN "linux_configuration_patch_settings_assessment_mode";
ALTER TABLE IF EXISTS azure_compute_virtual_machines
    DROP COLUMN "network_profile_network_api_version";
ALTER TABLE IF EXISTS azure_compute_virtual_machines
    DROP COLUMN "network_profile_network_interface_configurations";
ALTER TABLE IF EXISTS azure_compute_virtual_machines
    DROP COLUMN "scheduled_events_profile";
ALTER TABLE IF EXISTS azure_compute_virtual_machines
    DROP COLUMN "user_data";


ALTER TABLE IF EXISTS azure_compute_virtual_machine_resources
    ADD COLUMN "virtual_machine_extension_properties" jsonb;
ALTER TABLE IF EXISTS azure_compute_virtual_machine_resources
    DROP COLUMN "force_update_tag";
ALTER TABLE IF EXISTS azure_compute_virtual_machine_resources
    DROP COLUMN "publisher";
ALTER TABLE IF EXISTS azure_compute_virtual_machine_resources
    DROP COLUMN "type_handler_version";
ALTER TABLE IF EXISTS azure_compute_virtual_machine_resources
    DROP COLUMN "auto_upgrade_minor_version";
ALTER TABLE IF EXISTS azure_compute_virtual_machine_resources
    DROP COLUMN "enable_automatic_upgrade";
ALTER TABLE IF EXISTS azure_compute_virtual_machine_resources
    DROP COLUMN "settings";
ALTER TABLE IF EXISTS azure_compute_virtual_machine_resources
    DROP COLUMN "protected_settings";
ALTER TABLE IF EXISTS azure_compute_virtual_machine_resources
    DROP COLUMN "provisioning_state";
ALTER TABLE IF EXISTS azure_compute_virtual_machine_resources
    DROP COLUMN "extension_type";
ALTER TABLE IF EXISTS azure_compute_virtual_machine_resources
    DROP COLUMN "instance_view";


CREATE TABLE IF NOT EXISTS "azure_compute_virtual_machine_network_interfaces"
(
    "cq_id"                                          uuid                        NOT NULL,
    "cq_meta"                                        jsonb,
    "cq_fetch_date"                                  timestamp without time zone NOT NULL,
    "virtual_machine_cq_id"                          uuid,
    "virtual_machine_id"                             text,
    "network_interface_reference_properties_primary" boolean,
    "id"                                             text,
    CONSTRAINT azure_compute_virtual_machine_network_interfaces_pk PRIMARY KEY (cq_fetch_date, virtual_machine_cq_id, id),
    UNIQUE (cq_fetch_date, cq_id)
);
CREATE INDEX ON azure_compute_virtual_machine_network_interfaces (cq_fetch_date, virtual_machine_cq_id);
SELECT setup_tsdb_child('azure_compute_virtual_machine_network_interfaces', 'virtual_machine_cq_id',
                        'azure_compute_virtual_machines', 'cq_id');

-- Resource: security.jit_network_access_policies
DROP TABLE IF EXISTS azure_security_jit_network_access_policy_virtual_machines;
DROP TABLE IF EXISTS azure_security_jit_network_access_policy_requests;
DROP TABLE IF EXISTS azure_security_jit_network_access_policies;

DROP TABLE IF EXISTS azure_resources_links;
