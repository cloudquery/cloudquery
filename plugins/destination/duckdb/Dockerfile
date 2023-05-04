FROM ghcr.io/gythialy/golang-cross:1.20.2-0

# install goreleaser
# https://github.com/goreleaser/goreleaser-pro/releases/download/v1.17.2-pro/goreleaser-pro_Linux_x86_64.tar.gz
ARG GORELEASER_VERSION=1.17.2-pro
ARG GORELEASER_SHA=ca352fed092ed6ff2db50447fcd9137c266739763fa3a18e181984aee1d103e6
RUN  \
  GORELEASER_DOWNLOAD_FILE=goreleaser-pro_Linux_x86_64.tar.gz && \
  GORELEASER_DOWNLOAD_URL=https://github.com/goreleaser/goreleaser-pro/releases/download/v${GORELEASER_VERSION}/${GORELEASER_DOWNLOAD_FILE} && \
  wget ${GORELEASER_DOWNLOAD_URL}; \
  echo "$GORELEASER_SHA $GORELEASER_DOWNLOAD_FILE" | sha256sum -c - || exit 1; \
  tar -xzf $GORELEASER_DOWNLOAD_FILE -C /usr/bin/ goreleaser; \
  rm $GORELEASER_DOWNLOAD_FILE;

# Copy generic compiler wrappers
COPY gencc.sh /usr/bin/gencc.sh
COPY gencpp.sh /usr/bin/gencpp.sh