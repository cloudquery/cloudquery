ALTER TABLE IF EXISTS "gcp_iam_service_accounts" RENAME COLUMN unique_id TO id;

ALTER TABLE IF EXISTS "gcp_iam_service_accounts" DROP CONSTRAINT IF EXISTS "gcp_iam_service_accounts_pk";