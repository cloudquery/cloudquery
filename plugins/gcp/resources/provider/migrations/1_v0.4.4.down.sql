ALTER TABLE "gcp_iam_service_accounts"
    RENAME COLUMN id TO unique_id;

ALTER TABLE IF EXISTS "gcp_iam_service_accounts" ADD CONSTRAINT "gcp_iam_service_accounts_pk" UNIQUE ("project_id", "name");