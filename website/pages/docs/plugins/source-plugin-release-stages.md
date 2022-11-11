# Source Plugin Release Stages

[Official source plugins](sources) go through two release stages: Preview, and GA.

Both Preview and GA plugins follow [semantic versioning](#semantic-versioning).
The main differences between the two stages are:

1. Preview plugins are still experimental and may have frequent breaking changes.
2. Preview plugins might get deprecated due to lack of usage.
3. Long Term Support with community Discord and bug fixes is only guaranteed for GA plugins.

## Semantic Versioning

For version `Major.Minor.Patch`:

- `Major` is incremented when there are breaking changes (e.g. breaking configuration spec structure, column type changes).
- `Minor` is incremented when we add features in a backwards compatible way.
- `Patch` is incremented when we fix bugs in a backwards compatible way.
