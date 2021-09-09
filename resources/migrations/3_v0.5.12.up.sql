ALTER TABLE "aws_access_analyzer_analyzer_finding_sources" DROP CONSTRAINT IF EXISTS "aws_access_analyzer_analyzer_finding_sources_pk";

ALTER TABLE "aws_access_analyzer_analyzer_finding_sources" DROP CONSTRAINT IF EXISTS "aws_access_analyzer_analyzer_finding_sources_cq_id_key";

ALTER TABLE aws_access_analyzer_analyzer_finding_sources
    ADD CONSTRAINT "aws_access_analyzer_analyzer_finding_sources_pk"
        PRIMARY KEY(cq_id);

