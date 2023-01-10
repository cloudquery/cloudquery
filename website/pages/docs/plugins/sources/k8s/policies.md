# Policies and Compliance Frameworks

Policies are a set of standard SQL queries that can be run to check the security and compliance of your cloud resources against best practice frameworks.

This page documents the available CloudQuery SQL Policies for Kubernetes. See the [readme on GitHub](https://github.com/cloudquery/cloudquery/tree/main/plugins/source/k8s/policies) for installation instructions.
## Kubernetes NSA CISA v1

### Requirements
Kubernetes NSA CISA v1 requires the following tables to be synced before the policy is executed:

```yaml
tables:
  - k8s_apps_daemon_sets
  - k8s_apps_deployments
  - k8s_apps_replica_sets
  - k8s_batch_jobs
  - k8s_core_limit_ranges
  - k8s_core_namespaces
  - k8s_core_pods
  - k8s_core_resource_quotas
  - k8s_core_service_accounts
  - k8s_networking_network_policies
```

### Queries
Kubernetes NSA CISA v1 performs the following checks:
  - Daemonset enforces cpu limits
  - Deployment enforces cpu limits
  - Job enforces cpu limits
  - Namespaces CPU default resource limit
  - Namespace enforces resource quota cpu limits
  - Replicaset enforces cpu limits
  - Daemonset enforces cpu requests
  - Deployment enforces cpu requests
  - Namespaces CPU request resource quota
  - Namespace enforces resource quota cpu request
  - Replicaset enforces cpu requests
  - Daemonset enforces memory limits
  - Deployment enforces memory limits
  - Job enforces memory limit
  - Namespaces Memory default resource limit
  - Namespace enforces resource quota memory limits
  - Replicaset enforces memory limits
  - Daemonset enforces memory requests
  - Deployment enforces memory requests
  - Job enforces memory requests
  - Namespaces Memory request resource quota
  - Namespace enforces resource quota memory request
  - Replicaset enforces memory requests
  - Network policy default deny egress
  - Network policy default deny ingress
  - Pod volume don''t have a hostPath
  - DaemonSet containers privileges disabled
  - Deployments privileges disabled
  - Job containers privileges disabled
  - Pod container privileged access disabled
  - Replicaset privileges disabled
  - DaemonSet containers privilege escalation disabled
  - Deployments privilege escalation disabled
  - Job containers privilege escalation disabled
  - Pod container privilege escalation disabled
  - ReplicaSet container privileged escalation disabled
  - Deamonset container hostNetwork disabled
  - Deployments container hostNetwork disabled
  - Jobs container hostNetwork disabled
  - Pods container hostNetwork disabled
  - ReplicaSet container hostNetwork disabled
  - Deamonset containers HostPID and HostIPC sharing disabled
  - Deployment containers HostPID and HostIPC sharing disabled
  - Job containers HostPID and HostIPC sharing disabled
  - Pod containers HostPID and HostIPC sharing disabled
  - ReplicaSet containers HostPID and HostIPC sharing disabled
  - DeamonSet containers root file system is read-only
  - Deployment containers root file system is read-only
  - Job containers root file system is read-only
  - Pod container filesystem is read-ony
  - ReplicaSet containers root file system is read-only
  - DaemonSet containers to run as non-root
  - Deployment containers to run as non-root
  - Job containers run as non-root
  - Pod container runs as non-root
  - ReplicaSet containers must run as non-root
  - Pod service account tokens disabled
  - Pod service account tokens disabled"
