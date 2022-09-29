# Source Plugin Release Stages

[Official](./sources) source plugins go through multiple release stages: Alpha, Beta and GA.

## Comparison Matrix

See the following comparison matrix to understand differences and expectations.

|                                | Alpha | Beta                | GA                  |
| ------------------------------ | ----- | ------------------- | ------------------- |
| Spec (Config) Breaking Changes | Yes   | No                  | No                  |
| Schema Changes                 | Yes   | Semantic Versioning | Semantic Versioning |
| Long Term Support (Bugfixes)   | No    | Yes                 | Yes                 |

- Spec (Config) Breaking Changes - While a plugin is in Alpha, we might change how the plugin spec is defined for a plugin in a backward-incompatible way. Starting Beta we guarantee no **breaking** changes (we will still be adding features).
- Schema Changes - While a plugin is in Alpha, breaking changes to schema are possible without the plugin following semantic versioning. Starting Beta, we do our best to not break the schema in non-backward compatible way, but if we do then plugin will follow [semantic versioning](#semantic-versioning).
- Long Term Support - When a plugin is an Alpha, it is still experimental, and we might decide not to support it and/or deprecate it (including bugfixes). When a plugin is entering Beta phase, we guarantee long-term support with community Discord and bugfixes (additional features and resources roadmap can be found in GitHub issues).

## Semantic Versioning

Official plugins follow semantic versioning with the following logic:

For version `x.y.z`:

- `x`: `Spec` configuration changed in backward incompatible way.
- `y`: Incompatible schema where one or more columns/tables changed in a way that will require you to re-create/delete columns/tables. Documentation on which columns/tables are affected will be documented in the plugin release notes.
- `z`: No breaking changes. Features and/or bugfixes are documented in the plugin release notes.
