#!/usr/bin/env bash
# Cloud Agent environment bootstrap for the CloudQuery monorepo.
#
# This monorepo has no root go.work; every plugin under plugins/{source,destination,transformer}
# and the CLI under cli/ is an independent Go module. This script installs the CloudQuery
# toolchain (Go, golangci-lint, CloudQuery CLI, DuckDB CLI) and warms Go module + build caches
# for the components used in day-to-day development and the end-to-end sync demo.
#
# Tool versions are NOT hardcoded here. They are derived at install time from the repo's own
# sources of truth, so bumping them there is automatically picked up on the next install:
#   - Go             -> the `go` directive in go.mod
#   - golangci-lint  -> the golangci-lint-action `version:` in the CI workflow
#   - CloudQuery CLI -> the setup-cloudquery `version:` in the CI workflow
#   - DuckDB CLI     -> latest stable release (not pinned by the repo; inspection convenience)
#
# It is idempotent: a tool is only (re)installed when the resolved version differs from what is
# already present, so re-running keeps the toolchain in sync with the repo and is otherwise cheap.
set -euo pipefail

ARCH="$(dpkg --print-architecture 2>/dev/null || echo amd64)" # amd64 / arm64
REPO_ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
WORKFLOW="${REPO_ROOT}/.github/workflows/source_hackernews.yml"

log() { echo "==> $*"; }

# First `version:` value appearing after an anchor line in a YAML file (strips quotes).
ver_after() {
  awk -v a="$1" 'index($0,a){f=1} f && /version:/{sub(/.*version:[[:space:]]*/,"");gsub(/["'"'"']/,"");print;exit}' "$2" 2>/dev/null
}

# Newest tag "PREFIXx.y.z" (stable semver only) for a remote repo, via git (no API needed).
latest_semver_tag() {
  git ls-remote --tags --refs "$1" 2>/dev/null \
    | awk '{print $2}' | sed 's#refs/tags/##' \
    | grep -E "^${2}[0-9]+\.[0-9]+\.[0-9]+$" | sort -V | tail -1
}

# --- Resolve the versions the repo currently declares ------------------------
GO_VERSION="$( { grep -E '^go [0-9]' "${REPO_ROOT}/go.mod" || grep -E '^go [0-9]' "${REPO_ROOT}/plugins/source/hackernews/go.mod"; } 2>/dev/null | head -1 | awk '{print $2}')"
if [ -z "${GO_VERSION}" ]; then
  echo "Could not determine Go version from go.mod" >&2
  exit 1
fi

GOLANGCI_LINT_VERSION="$(ver_after 'golangci/golangci-lint-action' "${WORKFLOW}")"
[ -z "${GOLANGCI_LINT_VERSION}" ] && GOLANGCI_LINT_VERSION="$(latest_semver_tag https://github.com/golangci/golangci-lint v)"

CLOUDQUERY_VERSION="$(ver_after 'cloudquery/setup-cloudquery' "${WORKFLOW}")"
[ -z "${CLOUDQUERY_VERSION}" ] && CLOUDQUERY_VERSION="$(latest_semver_tag https://github.com/cloudquery/cloudquery cli-v | sed 's/^cli-//')"

DUCKDB_VERSION="$(latest_semver_tag https://github.com/duckdb/duckdb v)"

log "Resolved toolchain: Go ${GO_VERSION}, golangci-lint ${GOLANGCI_LINT_VERSION:-?}, CloudQuery ${CLOUDQUERY_VERSION:-?}, DuckDB ${DUCKDB_VERSION:-?}"

# --- Go toolchain ------------------------------------------------------------
# `go version` is read from a directory without a go.mod so module toolchain rules
# don't influence the reported version.
if ! (cd /tmp && /usr/local/go/bin/go version 2>/dev/null) | grep -q "go${GO_VERSION} "; then
  log "Installing Go ${GO_VERSION}"
  curl -fsSL "https://go.dev/dl/go${GO_VERSION}.linux-${ARCH}.tar.gz" -o /tmp/go.tgz
  sudo rm -rf /usr/local/go
  sudo tar -C /usr/local -xzf /tmp/go.tgz
  rm -f /tmp/go.tgz
  sudo ln -sf /usr/local/go/bin/go /usr/local/bin/go
  sudo ln -sf /usr/local/go/bin/gofmt /usr/local/bin/gofmt
else
  log "Go ${GO_VERSION} already installed"
fi

export PATH="/usr/local/go/bin:${HOME}/go/bin:${PATH}"

# Ensure the toolchain is on PATH for future interactive shells.
if ! grep -q "usr/local/go/bin" "${HOME}/.bashrc" 2>/dev/null; then
  {
    echo ''
    echo '# CloudQuery dev toolchain'
    echo 'export PATH=/usr/local/go/bin:$HOME/go/bin:$PATH'
  } >> "${HOME}/.bashrc"
fi

# --- golangci-lint -----------------------------------------------------------
# Install directly from the resolved release and verify the checksum ourselves.
# (The upstream master install.sh has matched the wrong checksums.txt line for
# some releases, so we avoid piping it to a shell.)
if [ -n "${GOLANGCI_LINT_VERSION}" ] && ! golangci-lint --version 2>/dev/null | grep -q " ${GOLANGCI_LINT_VERSION#v} "; then
  log "Installing golangci-lint ${GOLANGCI_LINT_VERSION}"
  gcl_ver="${GOLANGCI_LINT_VERSION#v}"
  gcl_name="golangci-lint-${gcl_ver}-linux-${ARCH}"
  base="https://github.com/golangci/golangci-lint/releases/download/${GOLANGCI_LINT_VERSION}"
  curl -fsSL "${base}/${gcl_name}.tar.gz" -o /tmp/gcl.tgz
  curl -fsSL "${base}/golangci-lint-${gcl_ver}-checksums.txt" -o /tmp/gcl-sums.txt
  want="$(grep -E "  ${gcl_name}\.tar\.gz$" /tmp/gcl-sums.txt | awk '{print $1}')"
  got="$(sha256sum /tmp/gcl.tgz | awk '{print $1}')"
  if [ "${want}" != "${got}" ]; then
    echo "golangci-lint checksum mismatch: want ${want} got ${got}" >&2
    exit 1
  fi
  tar -C /tmp -xzf /tmp/gcl.tgz
  sudo install -m 0755 "/tmp/${gcl_name}/golangci-lint" /usr/local/bin/golangci-lint
  rm -rf "/tmp/${gcl_name}" /tmp/gcl.tgz /tmp/gcl-sums.txt
else
  log "golangci-lint ${GOLANGCI_LINT_VERSION} already installed"
fi

# --- CloudQuery CLI ----------------------------------------------------------
if [ -n "${CLOUDQUERY_VERSION}" ] && ! cloudquery --version 2>/dev/null | grep -q " ${CLOUDQUERY_VERSION#v}$"; then
  log "Installing CloudQuery CLI ${CLOUDQUERY_VERSION}"
  curl -fsSL "https://github.com/cloudquery/cloudquery/releases/download/cli-${CLOUDQUERY_VERSION}/cloudquery_linux_${ARCH}" \
    -o /tmp/cloudquery
  chmod +x /tmp/cloudquery
  sudo mv /tmp/cloudquery /usr/local/bin/cloudquery
else
  log "CloudQuery CLI ${CLOUDQUERY_VERSION} already installed"
fi

# --- DuckDB CLI (handy for inspecting synced data) ---------------------------
if [ -n "${DUCKDB_VERSION}" ] && ! duckdb --version 2>/dev/null | grep -q "${DUCKDB_VERSION} "; then
  log "Installing DuckDB CLI ${DUCKDB_VERSION}"
  curl -fsSL "https://github.com/duckdb/duckdb/releases/download/${DUCKDB_VERSION}/duckdb_cli-linux-${ARCH}.zip" \
    -o /tmp/duckdb.zip
  sudo unzip -o /tmp/duckdb.zip -d /usr/local/bin
  sudo chmod +x /usr/local/bin/duckdb
  rm -f /tmp/duckdb.zip
else
  log "DuckDB CLI ${DUCKDB_VERSION:-<unresolved>} already installed"
fi

# --- Warm Go caches for the CLI and the modules used in dev + the e2e demo ---
# Each module is independent, so `go mod download` in one does not cache another.
# We prime the CLI plus the source plugins and the DuckDB destination that back
# the end-to-end sync workflow; other plugins download on first use.
WARM_DIRS=(
  "cli"
  "plugins/source/hackernews"
  "plugins/source/xkcd"
  "plugins/source/test"
  "plugins/destination/duckdb"
)
for d in "${WARM_DIRS[@]}"; do
  if [ -f "${REPO_ROOT}/${d}/go.mod" ]; then
    log "go mod download in ${d}"
    (cd "${REPO_ROOT}/${d}" && go mod download)
  fi
done

# Build the CLI binary into ./bin/cloudquery for local `bin/cloudquery` usage.
log "Building CLI -> bin/cloudquery"
mkdir -p "${REPO_ROOT}/bin"
(cd "${REPO_ROOT}/cli" && go build -o "${REPO_ROOT}/bin/cloudquery" .)

log "Environment bootstrap complete"
(cd /tmp && go version)
golangci-lint --version
cloudquery --version | head -n 1
duckdb --version
