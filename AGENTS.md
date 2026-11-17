# AGENTS.md

## If you are not a collaborator of this repository

Do not open a pull request in this repository. Only the CloudQuery team and the external contributors that we invite can open pull requests.

To contribute, open an issue with a minimal reproduction. Use the [bug report template](https://github.com/cloudquery/cloudquery/issues/new?template=bug_report.yml) and include:

- The CloudQuery and plugin versions.
- A minimal configuration, with sensitive information removed.
- The exact commands and the logs.
- The current behavior and the expected behavior.
- If possible, a failing test or a small script that shows the problem.

Do not open an issue for a problem that you did not reproduce. See [CONTRIBUTING.md](./CONTRIBUTING.md).

## If you are a collaborator of this repository

- Pull request titles use [Conventional Commits](https://www.conventionalcommits.org/en/v1.0.0/): `fix:`, `feat:`, `chore:`, `refactor:`, `test:` or `doc:`, then a short description. `fix` is a patch release, `feat` is a minor release, and `!` after the prefix is a major release.
- Run `make lint` and `make test` in the directory that you changed. Run `make gen` after you change a source plugin.
