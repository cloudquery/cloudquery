---
title: Getting Started with CloudQuery on GCP with Kubernetes
tag: security
date: 2022/10/20
description: >-
  Gain visibility into your GKE clusters with CloudQuery.
author: SCKelemen
---
# Getting Started with CloudQuery on GCP with Kubernetes
Samuel Kelemen
20 October 2022

In this tutorial, you'll learn how to get started with Kubernetes on GCP's Kubernetes Engine.
You'll then learn how to gather insights on your cluster using CloudQuery.

This tutorial relies on the following tools. If you haven't already installed them, take a moment
to install them before progressing: 

Prerequisites:
 - [gcloud](https://cloud.google.com/sdk/docs/install)
 - [kubectl](https://kubernetes.io/docs/tasks/tools/#kubectl)
 - [gke-gcloud-auth-plugin](https://cloud.google.com/blog/products/containers-kubernetes/kubectl-auth-changes-in-gke)
 - [postgres]()
 - [psql]()




Let’s start by creating a kubernetes cluster on GKE.
Kubernetes clusters are a pool of compute resources, on which workloads can be scheduled.
Clusters are a set of machines tied to a specific GCP region.

To create a cluster, you'll need the project name, the GCP region, and a name for your cluster.
${project-name} = cq-kelemen
${region} = europe-west1
${cluster-name} = cq-kelemen-cluster

The following command will create a cluster in the specified region, and setup the default network.
We will expose the network to the public internet, so that we, and our end users, can reach our services.
```
# create a kubernetes cluster and setup the default network
gcloud container --project "${project-name}" clusters create-auto "${cluster-name}" --region "${region}" --release-channel "regular" --network "projects/${project-name}/global/networks/default" --subnetwork "projects/${project-name}/regions/${region}/subnetworks/default" --cluster-ipv4-cidr "/17" --services-ipv4-cidr "/22"
```

Now that we have a Kubernetes cluster and an appropriate network, we can create some workloads
which will run on the cluster.

Google has micro services demo which simulates an E-Commerce store. We can deploy this on the cluster
by cloning the code and applying the kubernetes manifests. Kubernetes manifests are yaml files 
which contain all the information necessary to schedule the workloads on the clusters.

First, let's clone the demo:
```zsh
git clone https://github.com/GoogleCloudPlatform/microservices-demo.git
```
Let's take a minute to understand the kubernetes manifests.

Kubernetes manifests allow you to specify configuration to the kubernetes system.
You can configure your services from them manifests.

Next we will apply the manifests to the cluster.
```zsh
kubectl apply -f ./release/kubernetes-manifests.yaml
```

Deploying the resources could take a few minutes.
If everything is running correctly, you should be able to reach your e-commerce site.
It should look like [mine](http://35.205.158.178/).
![Screenshot: e-commerce store demo](/images/blog/getting-started-with-gcp-and-kubernetes/estore.png)

## Gathering insights with CloudQuery


GCP Plugin configuration, `gcp.yaml`:
```yaml
kind: source
spec:
  name: gcp
  version: "v1.0.1"
  destinations: ["postgresql"]
```

Kubernetes Plugin configuration, `k8s.yaml`:
```yaml
kind: source
spec:
  name: k8s
  version: "v2.0.1"
  destinations: ["postgresql"]
```

Postgres Plugin configuration, `psql.yaml`:
```yaml
kind: destination
spec:
  name: "postgresql"
  version: "v0.3.0"
  write_mode: "overwrite" # overwrite, append
  spec:
    connection_string: "postgresql://{CQ_PG_USER}:{CQ_PG_PASS}@localhost:5432/postgres?sslmode=disable"
```

Sync resources:
```zsh
cloudquery sync k8s.yaml gcp.yaml psql.yaml
```
output:
```
Loading spec(s) from k8s.yaml, gcp.yaml, psql.yaml
Downloading https://github.com/cloudquery/cloudquery/releases/download/plugins-source-k8s-v2.0.1/k8s_darwin_arm64.zip
 100% |█████████████████████████████████████████████████████████████████████████████████████████████████████| (11/11 MB, 5.5 MB/s)        
Starting sync for:  k8s -> [postgresql]
W1018 14:39:21.082077   85097 gcp.go:119] WARNING: the gcp auth plugin is deprecated in v1.22+, unavailable in v1.26+; use gcloud instead.
To learn more, consult https://cloud.google.com/blog/products/containers-kubernetes/kubectl-auth-changes-in-gke
Sync completed successfully.                          
Summary: resources: 311, errors: 0, panic: 0 failed_writes: 0
Starting sync for:  gcp -> [postgresql]
Sync completed successfully.                           
Summary: resources: 7035, errors: 18, panic: 0 failed_writes: 0
```

Our resources should now appear in our postgress database. 

View the pods in our Kubernetes cluster:
```sql
select * from k8s_core_pods;
```


Run the GCP CIS policy:
```
"psql -U postgres -h localhost -f /plugins/source/gcp/policies_v1/cis_v1.2.0/policy.sql"      

```