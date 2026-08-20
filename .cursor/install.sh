#!/usr/bin/env bash
# Cloud Agent environment bootstrap for the CloudQuery monorepo.
#
# This monorepo has no root go.work; every plugin under plugins/{source,destination,transformer}
# and the CLI under cli/ is an independent Go module. This script installs the pinned toolchain
# (Go, golangci-lint, CloudQuery CLI, DuckDB CLI) and warms Go module + build caches for the
# components used in day-to-day development and the end-to-end sync demo.
#
# It is idempotent: re-running it is a no-op once the pinned tools are present.
set -euo pipefail

# --- Pinned versions (kept in sync with plugins/source/hackernews/go.mod and
# --- .github/workflows/source_hackernews.yml) --------------------------------
GO_VERSION="1.26.5"
GOLANGCI_LINT_VERSION="v2.12.2"
CLOUDQUERY_VERSION="v6.41.1"
DUCKDB_VERSION="v1.5.5"

ARCH="$(dpkg --print-architecture 2>/dev/null || echo amd64)" # amd64 / arm64
REPO_ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"

log() { echo "==> $*"; }

# --- Go toolchain ------------------------------------------------------------
if ! /usr/local/go/bin/go version 2>/dev/null | grep -q "go${GO_VERSION} "; then
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
export GOTOOLCHAIN="local" # we install the exact version the modules pin

# Ensure the toolchain is on PATH for future interactive shells.
if ! grep -q "usr/local/go/bin" "${HOME}/.bashrc" 2>/dev/null; then
  {
    echo ''
    echo '# CloudQuery dev toolchain'
    echo 'export PATH=/usr/local/go/bin:$HOME/go/bin:$PATH'
  } >> "${HOME}/.bashrc"
fi

# --- golangci-lint -----------------------------------------------------------
# Install directly from the pinned release and verify the checksum ourselves.
# (The upstream master install.sh has matched the wrong checksums.txt line for
# some releases, so we avoid piping it to a shell.)
if ! golangci-lint --version 2>/dev/null | grep -q "${GOLANGCI_LINT_VERSION#v}"; then
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
if ! cloudquery --version 2>/dev/null | grep -q "${CLOUDQUERY_VERSION#v}"; then
  log "Installing CloudQuery CLI ${CLOUDQUERY_VERSION}"
  curl -fsSL "https://github.com/cloudquery/cloudquery/releases/download/cli-${CLOUDQUERY_VERSION}/cloudquery_linux_${ARCH}" \
    -o /tmp/cloudquery
  chmod +x /tmp/cloudquery
  sudo mv /tmp/cloudquery /usr/local/bin/cloudquery
else
  log "CloudQuery CLI ${CLOUDQUERY_VERSION} already installed"
fi

# --- DuckDB CLI (handy for inspecting synced data) ---------------------------
if ! duckdb --version 2>/dev/null | grep -q "${DUCKDB_VERSION}"; then
  log "Installing DuckDB CLI ${DUCKDB_VERSION}"
  curl -fsSL "https://github.com/duckdb/duckdb/releases/download/${DUCKDB_VERSION}/duckdb_cli-linux-${ARCH}.zip" \
    -o /tmp/duckdb.zip
  sudo unzip -o /tmp/duckdb.zip -d /usr/local/bin
  sudo chmod +x /usr/local/bin/duckdb
  rm -f /tmp/duckdb.zip
else
  log "DuckDB CLI ${DUCKDB_VERSION} already installed"
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
go version
golangci-lint --version
cloudquery --version | head -n 1
duckdb --version
