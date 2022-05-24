CREATE EXTENSION IF NOT EXISTS pg_trgm;

CREATE TABLE IF NOT EXISTS "policy_executions" (
  "id" uuid NOT NULL,
  "timestamp" timestamp,
  "scheme" text,
  "location" text,
  "policy_name" text,
  "selector" text,
  "sha256_hash" text,
  "version" text,
  "checks_total" int,
  "checks_failed" int,
  "checks_passed" int,
  PRIMARY KEY ("id")
);
CREATE INDEX IF NOT EXISTS "policy_executions_timestamp_idx" ON "policy_executions" USING btree ("timestamp");
CREATE INDEX IF NOT EXISTS "policy_executions_version_idx" ON "policy_executions" USING brin ("version");

CREATE TABLE IF NOT EXISTS "check_results" (
  "execution_id" uuid NOT NULL,
  "execution_timestamp" timestamp,
  "name" text,
  "selector" text NOT NULL,
  "description" text,
  "status" text,
  "raw_results" jsonb,
  "error" text,
  PRIMARY KEY ("execution_id", "selector"),
  FOREIGN KEY ("execution_id") REFERENCES "policy_executions" ("id") ON DELETE CASCADE
);
CREATE INDEX IF NOT EXISTS "check_results_execution_timestamp_idx" ON "check_results" USING btree ("execution_timestamp");
CREATE INDEX IF NOT EXISTS "check_results_selector_idx" ON "check_results" USING gin ("selector" gin_trgm_ops);

CREATE OR REPLACE FUNCTION "cloudquery"."calculate_policy_executions_stats"()
RETURNS TRIGGER AS $$
BEGIN
		UPDATE cloudquery.policy_executions SET
			checks_total = checks_total + 1,
			checks_failed = checks_failed + CASE WHEN NEW.status = 'failed' THEN 1 ELSE 0 END,
			checks_passed = checks_passed + CASE WHEN NEW.status = 'passed' THEN 1 ELSE 0 END
			WHERE id = NEW.execution_id;
    RETURN NEW;
END;
$$ LANGUAGE "plpgsql";

CREATE TRIGGER "CalculatePolicyExecutionsStats"
AFTER INSERT ON "cloudquery"."check_results"
FOR EACH ROW
EXECUTE PROCEDURE "cloudquery"."calculate_policy_executions_stats"();
