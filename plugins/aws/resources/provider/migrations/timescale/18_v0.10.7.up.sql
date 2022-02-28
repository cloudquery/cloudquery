ALTER TABLE IF EXISTS aws_route53_traffic_policy_versions
    DROP CONSTRAINT aws_route53_traffic_policy_versions_pk;
-- a PK without cq_fetch_date is not valid if we have data in the table, and this was released with an invalid PK definition...
-- we're fixing this migration and releasing a new one at the same time, so this fix has to be the same both in 18 and 19.
ALTER TABLE IF EXISTS aws_route53_traffic_policy_versions
    ADD CONSTRAINT aws_route53_traffic_policy_versions_pk PRIMARY KEY (cq_fetch_date, traffic_policy_cq_id, id, version);