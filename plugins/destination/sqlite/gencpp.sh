#!/bin/bash
if [ "${GOOS}" = "windows" ]; then
  x86_64-w64-mingw32-g++ $@
elif [ "$GOOS" = "linux" ]; then
  if [ "$GOARCH" = "arm64" ]; then
    aarch64-linux-gnu-g++ $@
  else
    g++ $@
  fi
elif [ "$GOOS" = "darwin" ]; then
  o64-clang++ $@
fi