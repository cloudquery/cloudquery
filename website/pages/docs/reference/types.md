---
title: CloudQuery Types
description: CloudQuery uses Apache Arrow as its internal type system, and defines a number of custom types as Arrow extensions.
---

# CloudQuery Types

CloudQuery uses [Apache Arrow](https://arrow.apache.org/docs/index.html) to represent data internally. Source integrations define columns in terms of Arrow types, and destinations support converting from Arrow to their own native types.

Apart from the native Arrow types, CloudQuery also defines a number of custom types, implemented as Arrow extensions:

- `JSON`: A valid JSON object, stored as binary.
- `Inet`: An IP address or network, stored as binary.
- `MAC`: A MAC address, stored as binary.
- `UUID`: A 16-byte UUID using Arrow's fixed size binary type as its storage layer.