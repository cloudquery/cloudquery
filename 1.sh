set -o pipefail

LIMIT=100

          echo Checking files with size over ${LIMIT}kB: \\n

find . \
-iname '*.png' -size +${LIMIT}k -o \
-iname '*.jpg' -size +${LIMIT}k -o \
-iname '*.jpeg' -size +${LIMIT}k \
| xargs du -sh
