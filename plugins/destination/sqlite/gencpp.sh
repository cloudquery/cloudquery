#!/bin/bash
if [ "${GOOS}" = "windows" ]; then
  x86_64-w64-mingw32-g++ $@
elif [ "${GOOS}" = "linux" ]; then
  g++ $@
elif [ "${GOOS}" = "darwin" ]; then
  o64-clang++ $@
fi
