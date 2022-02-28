ALTER TABLE IF EXISTS aws_route53_traffic_policy_versions
    DROP CONSTRAINT aws_route53_traffic_policy_versions_pk;
ALTER TABLE IF EXISTS aws_route53_traffic_policy_versions
    ADD CONSTRAINT aws_route53_traffic_policy_versions_pk PRIMARY KEY (cq_fetch_date, traffic_policy_cq_id, id, version);
