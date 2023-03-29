#!/bin/bash
# Original at https://github.com/gvenzl/oci-oracle-xe/blob/c96c28fc41dd1735abf03d508d4ea3f2cf76d14e/healthcheck.sh
# 
# Since: January, 2021
# Author: gvenzl
# Name: healthcheck.sh
# Description: Checks the health of the database
#
# Copyright 2021 Gerald Venzl
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

# Exit on errors
# Great explanation on https://vaneyckt.io/posts/safer_bash_scripts_with_set_euxo_pipefail/
set -Eeuo pipefail

db_status=$(sqlplus -s / << EOF
   set heading off;
   set pagesize 0;
   SELECT status FROM v\$instance;
   exit;
EOF
)

if [ "${db_status}" == "OPEN" ]; then
   exit 0;
else
   exit 1;
fi;