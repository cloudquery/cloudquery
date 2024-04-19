#!/bin/bash
if [ "${GOOS}" = "windows" ]; then
  x86_64-w64-mingw32-gcc $@
elif [ "$GOOS" = "linux" ]; then
  if [ "$GOARCH" = "arm64" ]; then
    aarch64-linux-gnu-gcc $@
  else
    gcc $@
  fi
elif [ "$GOOS" = "darwin" ]; then
  o64-clang $@
fi