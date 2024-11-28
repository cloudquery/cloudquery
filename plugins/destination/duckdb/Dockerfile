FROM ghcr.io/gythialy/golang-cross:1.23.3-0

# install goreleaser
# https://github.com/goreleaser/goreleaser-pro/releases/download/v1.12.3-pro/goreleaser-pro_Linux_x86_64.tar.gz
ARG GORELEASER_VERSION=1.12.3-pro
ARG GORELEASER_SHA=af9288138078ec3291656b70f2c3a08476b0a7bafdeb43c7ddf778fbb761caba
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