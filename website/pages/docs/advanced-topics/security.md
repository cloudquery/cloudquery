---
title: Security
---

# Security

This section will list key security points regarding CloudQuery. Make sure you follow best practices if you decide to "host it yourself."

## Plugin Authentication Credentials

- Plugin Authentication Credentials should always be read-only.
- The machine where CloudQuery is running should be secured with the correct permissions, as it contains the credentials to your cloud infrastructure.

## CloudQuery Database

Even though the CloudQuery database contains only configuration and meta-data, you should protect it and keep it secure with correct access and permissions.
