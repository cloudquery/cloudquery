ALTER TABLE IF EXISTS "aws_elasticsearch_domains" DROP CONSTRAINT IF EXISTS "aws_elasticsearch_domains_pk";

ALTER TABLE "aws_elasticsearch_domains"
    ADD CONSTRAINT "aws_elasticsearch_domains_pk"
        PRIMARY KEY(account_id, region, id);

