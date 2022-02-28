ALTER TABLE IF EXISTS aws_route53_traffic_policy_versions
    DROP CONSTRAINT aws_route53_traffic_policy_versions_pk;
-- a PK without cq_fetch_date is not valid if we have data in the table, so we're keeping the down migration same as the up one
ALTER TABLE IF EXISTS aws_route53_traffic_policy_versions
    ADD CONSTRAINT aws_route53_traffic_policy_versions_pk PRIMARY KEY (cq_fetch_date, traffic_policy_cq_id, id, version);
