\echo "Executing K8S Pod Security NSA CISA v1"
\set check_id "container_disallow_host_path"
\echo "Disallow host path access"
\i queries/pod_security/pod_volume_host_path.sql


\echo "Verify containers have privileged access disabled"

\echo "Deamonset privileges disabled"
\set check_id "daemonset_container_privilege_disabled"
\i queries/pod_security/daemonset_container_privilege_disabled.sql

\echo "Deployment containers privileged access disabled"
\set check_id "deployment_container_privilege_disabled"
\i queries/pod_security/deployment_container_privilege_disabled.sql

\echo "Jobs container privileged access disabled"
\set check_id "job_container_privilege_disabled"
\i queries/pod_security/job_container_privilege_disabled.sql

\echo "Pod container privileged access disabled"
\set check_id "pod_container_privilege_disabled"
\i queries/pod_security/pod_container_privilege_disabled.sql

\echo "ReplicaSet container privileged access disabled"
\set check_id "replicaset_container_privilege_disabled"
\i queries/pod_security/replicaset_container_privilege_disabled.sql

\echo "Container privileged escalation disabled"

\echo "DaemonSet container privileged escalation disabled"
\set check_id "daemonset_container_privilege_escalation_disabled"
\i queries/pod_security/daemonset_container_privilege_escalation_disabled.sql

\echo "Deployment container privileged escalation disabled"
\set check_id "deployment_container_privilege_escalation_disabled"
\i queries/pod_security/deployment_container_privilege_escalation_disabled.sql

\echo "Job container privileged escalation disabled"
\set check_id "job_container_privilege_escalation_disabled"
\i queries/pod_security/job_container_privilege_escalation_disabled.sql

\echo "Pod container privileged escalation disabled"
\set check_id "pod_container_privilege_escalation_disabled"
\i queries/pod_security/pod_container_privilege_escalation_disabled.sql

\echo "ReplicaSet container privileged escalation disabled"
\set check_id "replicaset_container_privilege_escalation_disabled"
\i queries/pod_security/replicaset_container_privilege_escalation_disabled.sql


\echo "Host network access disabled"

\echo "DaemonSet container hostNetwork disabled"
\set check_id "daemonset_host_network_access_disabled"
\i queries/pod_security/daemonset_host_network_access_disabled.sql

\echo "Deployment container hostNetwork disabled"
\set check_id "deployment_host_network_access_disabled"
\i queries/pod_security/deployment_host_network_access_disabled.sql

\echo "Job container hostNetwork disabled"
\set check_id "job_host_network_access_disabled"
\i queries/pod_security/job_host_network_access_disabled.sql

\echo "Pod container hostNetwork disabled"
\set check_id "pod_container_privilege_escalation_disabled"
\i queries/pod_security/pod_host_network_access_disabled.sql

\echo "ReplicaSet container hostNetwork disabled"
\set check_id "replicaset_container_privilege_escalation_disabled"
\i queries/pod_security/replicaset_host_network_access_disabled.sql


\echo "HostPID and HostIPC sharing disabled"

\echo "DeamonSet containers HostPID and HostIPC sharing disabled"
\set check_id "daemonset_hostpid_hostipc_sharing_disabled"
\i queries/pod_security/daemonset_hostpid_hostipc_sharing_disabled.sql

\echo "Deployment containers HostPID and HostIPC sharing disabled"
\set check_id "deployment_hostpid_hostipc_sharing_disabled"
\i queries/pod_security/deployment_hostpid_hostipc_sharing_disabled.sql

\echo "Job containers HostPID and HostIPC sharing disabled"
\set check_id "job_hostpid_hostipc_sharing_disabled"
\i queries/pod_security/job_hostpid_hostipc_sharing_disabled.sql

\echo "Pod containers HostPID and HostIPC sharing disabled"
\set check_id "pod_hostpid_hostipc_sharing_disabled"
\i queries/pod_security/pod_hostpid_hostipc_sharing_disabled.sql

\echo "ReplicaSet containers HostPID and HostIPC sharing disabled"
\set check_id "replicaset_hostpid_hostipc_sharing_disabled"
\i queries/pod_security/replicaset_hostpid_hostipc_sharing_disabled.sql

\echo "Containers root file system is read-only"

\echo "DeamonSet containers root file system is read-only"
\set check_id "daemonset_immutable_container_filesystem"
\i queries/pod_security/daemonset_immutable_container_filesystem.sql

\echo "Deployment containers root file system is read-only"
\set check_id "deployment_immutable_container_filesystem"
\i queries/pod_security/deployment_immutable_container_filesystem.sql

\echo "Job containers root file system is read-only"
\set check_id "job_immutable_container_filesystem"
\i queries/pod_security/job_immutable_container_filesystem.sql

\echo "Pod containers root file system is read-only"
\set check_id "pod_immutable_container_filesystem"
\i queries/pod_security/pod_immutable_container_filesystem.sql

\echo "ReplicaSet containers root file system is read-only"
\set check_id "replicaset_immutable_container_filesystem"
\i queries/pod_security/replicaset_immutable_container_filesystem.sql


\echo "Enforce containers to run as non-root"

\echo "DeamonSet containers to run as non-root"
\set check_id "daemonset_non_root_container"
\i queries/pod_security/daemonset_non_root_container.sql

\echo "Deployment containers to run as non-root"
\set check_id "deployment_non_root_container"
\i queries/pod_security/deployment_non_root_container.sql

\echo "Job containers to run as non-root"
\set check_id "job_non_root_container"
\i queries/pod_security/job_non_root_container.sql

\echo "Pod containers to run as non-root"
\set check_id "pod_non_root_container"
\i queries/pod_security/pod_non_root_container.sql

\echo "ReplicaSet containers to run as non-root"
\set check_id "replicaset_non_root_container"
\i queries/pod_security/replicaset_non_root_container.sql


\echo "Automatic mapping of the service account tokens disabled"

\echo "Pod service account tokens disabled"
\set check_id "pod_service_account_token_disabled"
\i queries/pod_security/pod_service_account_token_disabled.sql

\echo "Service account tokens disabled"
\set check_id "service_account_token_disabled"
\i queries/pod_security/service_account_token_disabled.sql

