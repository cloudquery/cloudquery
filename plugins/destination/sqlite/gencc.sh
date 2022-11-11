#!/bin/bash
if [ "${GOOS}" = "windows" ]; then
  x86_64-w64-mingw32-gcc $@
elif [ "$GOOS" = "linux" ]; then
  gcc $@
elif [ "$GOOS" = "darwin" ]; then
  o64-clang $@
fi