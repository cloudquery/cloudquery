# Source Plugins Release Stages

[Official](./sources) source plugins go through multiple release stages: Alpha, Beta and GA.

## Comparision Matrix

See the following comparision matrix to understand difference and expectations.

|                                | Alpha | Beta                | GA                  |
| ------------------------------ | ----- | ------------------- | ------------------- |
| Spec (Config) Breaking Changes | Yes   | No                  | No                  |
| Schema Changes                 | Yes   | Semantic Versioning | Semantic Versioning |
| Long Term Support (Bugfixes)   | No    | Yes                 | Yes                 |

- Spec (Config) Breaking Changes - While a plugin is in Alpha we might change how the plugin spec is defined for a plugin in a backward incompatible way. Starting beta we guarantee no **breaking** changes (we still be adding features).
- Schema Changes - While a plugin is in Alpha breaking changes to schema are possible without plugin following semantic versioning. Starting Beta we do our best effort to not break schema in non backward compatible way but if we do then plugin will follow [semantic versioning](#semantic-versioning).
- Long Term Support - When plugin is an alpha it is still in experiment mode and we might decide not to support and/or deprecate it (including bugfixes). When plugin is entering Beta phase we guarantee long term support with community discord and bugfixes (additional features and resources roadmap can be found in GitHub issues).

## Semantic Versioning

Official plugins follow semantic versiniong with the following logic:

For version x.y.z changes in:

- x: `Spec` changed in backward incompatible way.
- y: Incompatible schema where one or more columns/tables changed in a way that will require to re-create/delete columns/tables. Documentation on which columns/tables are affected will be documented in the plugin release notes.
- z: No breaking changes. Either feature and/or bugfixes as documented in the plugin release notes.
