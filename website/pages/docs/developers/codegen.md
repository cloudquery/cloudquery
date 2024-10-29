---
title: Codegen
description: Learn how to generate code for CloudQuery integrations.
---

# Generating resources

Adding resources to an integrations can sometimes be a tedious task, some resources can have more than hundreds of fields and relations, and adding them all can
take a long time. To remedy this issue, we provide utilities as part of our [plugin-sdk](https://github.com/cloudquery/plugin-sdk) to automatically infer columns 
from Go structs. In particular, see the [`transformers.TransformWithStruct()` method](https://github.com/cloudquery/plugin-sdk/blob/main/transformers/struct.go).  
