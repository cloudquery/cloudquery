\echo "Executing K8S Network Hardening NSA CISA v1"

\echo "Enforce CPU resource limits"

\echo "Deamonsets enforce cpu limit"
\set check_id "daemonset_cpu_limit"
\i queries/network_hardening/daemonset_cpu_limit.sql

\echo "Deployments enforce cpu limit"
\set check_id "deployment_cpu_limit"
\i queries/network_hardening/deployment_cpu_limit.sql

\echo "Jobs enforce cpu limit"
\set check_id "job_cpu_limit"
\i queries/network_hardening/job_cpu_limit.sql


\echo "Namespaces CPU limit range default"
\set check_id "namespace_limit_range_default_cpu_limit"
\i queries/network_hardening/namespace_limit_range_default_cpu_limit.sql


\echo "Namespaces CPU limit resource quota"
\set check_id "namespace_resource_quota_cpu_limit"
\i queries/network_hardening/namespace_resource_quota_cpu_limit.sql


\echo "ReplciaSets enforce cpu limit"
\set check_id "replicaset_cpu_limit"
\i queries/network_hardening/replicaset_cpu_limit.sql


\echo "Enforce CPU request"

\echo "Deamonsets enforce cpu request"
\set check_id "daemonset_cpu_request"
\i queries/network_hardening/daemonset_cpu_request.sql

\echo "Deployments enforce cpu request"
\set check_id "deployment_cpu_request"
\i queries/network_hardening/deployment_cpu_request.sql

\echo "Jobs enforce cpu request"
\set check_id "job_cpu_limit"
\i queries/network_hardening/job_cpu_limit.sql


\echo "Namespaces CPU request range default"
\set check_id "namespace_limit_range_default_cpu_request"
\i queries/network_hardening/namespace_limit_range_default_cpu_request.sql


\echo "Namespaces CPU request resource quota"
\set check_id "namespace_resource_quota_cpu_request"
\i queries/network_hardening/namespace_resource_quota_cpu_request.sql


\echo "ReplciaSets enforce cpu request"
\set check_id "replicaset_cpu_request"
\i queries/network_hardening/replicaset_cpu_request.sql

\echo "Ensure contorl plane hardening"
\echo "Endpoint API served on secure port"
\set check_id "endpoint_api_serve_on_secure_port"
\i queries/network_hardening/endpoint_api_serve_on_secure_port.sql


\echo "Ensure memory limits set"

\echo "Deamonsets enforce memory limit"
\set check_id "daemonset_memory_limit"
\i queries/network_hardening/daemonset_memory_limit.sql

\echo "Deployments enforce memory limit"
\set check_id "deployment_memory_limit"
\i queries/network_hardening/deployment_memory_limit.sql

\echo "Jobs enforce memory limit"
\set check_id "job_memory_limit"
\i queries/network_hardening/job_memory_limit.sql


\echo "Namespaces CPU memory range default"
\set check_id "namespace_limit_range_default_memory_limit"
\i queries/network_hardening/namespace_limit_range_default_memory_limit.sql


\echo "Namespaces CPU memory resource quota"
\set check_id "namespace_resource_quota_memory_limit"
\i queries/network_hardening/namespace_resource_quota_memory_limit.sql


\echo "ReplciaSets enforce memory limit"
\set check_id "replicaset_memory_limit"
\i queries/network_hardening/replicaset_memory_limit.sql


\echo "Enforce Memory request"

\echo "Deamonsets enforce memory request"
\set check_id "daemonset_memory_request"
\i queries/network_hardening/daemonset_memory_request.sql

\echo "Deployments enforce memory request"
\set check_id "deployment_memory_request"
\i queries/network_hardening/deployment_memory_request.sql

\echo "Jobs enforce memory request"
\set check_id "job_memory_request"
\i queries/network_hardening/job_memory_request.sql


\echo "Namespaces Memory request range default"
\set check_id "namespace_limit_range_default_memory_request"
\i queries/network_hardening/namespace_limit_range_default_memory_request.sql


\echo "Namespaces Memory request resource quota"
\set check_id "namespace_resource_quota_memory_request"
\i queries/network_hardening/namespace_resource_quota_memory_request.sql


\echo "ReplciaSets enforce cpu request"
\set check_id "replicaset_memory_request"
\i queries/network_hardening/replicaset_memory_request.sql


\echo "Enforce default deny network policy"

\echo "Network policy default deny egress"
\set check_id "network_policy_default_deny_egress"
\i queries/network_hardening/network_policy_default_deny_egress.sql


\echo "Network policy default deny ingress"
\set check_id "network_policy_default_deny_ingress"
\i queries/network_hardening/network_policy_default_deny_ingress.sql

\echo "Network policy default don't allow egress"
\set check_id "network_policy_default_dont_allow_egress"
\i queries/network_hardening/network_policy_default_dont_allow_egress.sql

\echo "Network policy default don't allow ingress"
\set check_id "network_policy_default_dont_allow_ingress"
\i queries/network_hardening/network_policy_default_dont_allow_ingress.sql