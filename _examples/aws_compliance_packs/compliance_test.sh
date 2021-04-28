#!/bin/bash

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"

echo "ISO 27001 Section 6: Organization of information security"
echo "==========================================================="
echo "## VPCs without Flow Logs enabled:"
cat ${DIR}/vpc_flow_logs_disabled.sql | sqlite3 cloudquery.db -header -column

echo ""
echo "## VPCs without Flow Logs enabled:"
cat ${DIR}/rds_unencrypted.sql | sqlite3 cloudquery.db -header -column

echo ""
echo "## IAM Access Keys older than 90 days:"
cat ${DIR}/iam_access_keys_old.sql | sqlite3 cloudquery.db -header -column

echo ""
echo "## IAM Access Keys not used in the last 90 days:"
cat ${DIR}/iam_access_keys_not_used.sql | sqlite3 cloudquery.db -header -column

echo ""
echo "## Unencrypted S3 buckets"
cat ${DIR}/s3_unencrypted.sql | sqlite3 cloudquery.db -header -column

echo "ISO 27001 Section 12: Operational Security:"
echo "==========================================================="
echo "## RDS Instances with snapshot retention less than seven days"
cat ${DIR}/rds_clusters_less_than_seven_days_retention.sql | sqlite3 cloudquery.db -header -column

echo ""
echo "## RDS Instances without logging setup"
cat ${DIR}/rds_clusters_without_logging_setup.sql | sqlite3 cloudquery.db -header -column

echo ""
echo "## EC2 Instances without monitoring used"
cat ${DIR}/ec2_instances_without_monitoring.sql | sqlite3 cloudquery.db -header -column