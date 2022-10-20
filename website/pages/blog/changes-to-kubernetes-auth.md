---
title: Changes to Kubernetes Authentication
tag: kubernetes
date: 2022/10/20
description: >-
  Cluster authentication changes coming in v1.26
author: SCKelemen
---

import { BlogHeader } from "../../components/BlogHeader"

<BlogHeader/>

Kubernetes users may see the following message when running the K8s plugin on GKE Clusters:
```
WARNING: the gcp auth plugin is deprecated in v1.22+, unavailable in v1.26+; use gcloud instead.
```

As part of an initiative to remove platform specific code from Kubernetes, authentication will begin to be delegated to authentication plugins, starting in version 1.26.

## What does this mean for CloudQuery users?

CloudQuery does not use any specific resources. 
Users can upgrade by 


For more information, read [Google's press release]().