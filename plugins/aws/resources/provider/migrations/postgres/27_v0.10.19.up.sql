-- Resource: wafregional.rate_based_rules
CREATE TABLE IF NOT EXISTS "aws_wafregional_rate_based_rules" (
	"cq_id" uuid NOT NULL,
	"cq_meta" jsonb,
	"account_id" text,
	"region" text,
	"arn" text,
	"tags" jsonb,
	"rate_key" text,
	"rate_limit" bigint,
	"id" text,
	"metric_name" text,
	"name" text,
	CONSTRAINT aws_wafregional_rate_based_rules_pk PRIMARY KEY(account_id,region,id),
	UNIQUE(cq_id)
);
CREATE TABLE IF NOT EXISTS "aws_wafregional_rate_based_rule_match_predicates" (
	"cq_id" uuid NOT NULL,
	"cq_meta" jsonb,
	"rate_based_rule_cq_id" uuid,
	"data_id" text,
	"negated" boolean,
	"type" text,
	CONSTRAINT aws_wafregional_rate_based_rule_match_predicates_pk PRIMARY KEY(cq_id),
	UNIQUE(cq_id),
	FOREIGN KEY (rate_based_rule_cq_id) REFERENCES aws_wafregional_rate_based_rules(cq_id) ON DELETE CASCADE
);

-- Resource: wafregional.rule_groups
CREATE TABLE IF NOT EXISTS "aws_wafregional_rule_groups" (
	"cq_id" uuid NOT NULL,
	"cq_meta" jsonb,
	"account_id" text,
	"region" text,
	"arn" text,
	"tags" jsonb,
	"id" text,
	"metric_name" text,
	"name" text,
	CONSTRAINT aws_wafregional_rule_groups_pk PRIMARY KEY(account_id,region,id),
	UNIQUE(cq_id)
);

-- Resource: wafregional.rules
CREATE TABLE IF NOT EXISTS "aws_wafregional_rules" (
	"cq_id" uuid NOT NULL,
	"cq_meta" jsonb,
	"account_id" text,
	"region" text,
	"arn" text,
	"tags" jsonb,
	"id" text,
	"metric_name" text,
	"name" text,
	CONSTRAINT aws_wafregional_rules_pk PRIMARY KEY(account_id,region,id),
	UNIQUE(cq_id)
);
CREATE TABLE IF NOT EXISTS "aws_wafregional_rule_predicates" (
	"cq_id" uuid NOT NULL,
	"cq_meta" jsonb,
	"rule_cq_id" uuid,
	"data_id" text,
	"negated" boolean,
	"type" text,
	CONSTRAINT aws_wafregional_rule_predicates_pk PRIMARY KEY(cq_id),
	UNIQUE(cq_id),
	FOREIGN KEY (rule_cq_id) REFERENCES aws_wafregional_rules(cq_id) ON DELETE CASCADE
);

-- Resource: wafregional.web_acls
CREATE TABLE IF NOT EXISTS "aws_wafregional_web_acls" (
	"cq_id" uuid NOT NULL,
	"cq_meta" jsonb,
	"account_id" text,
	"region" text,
	"tags" jsonb,
	"default_action" text,
	"id" text,
	"metric_name" text,
	"name" text,
	"arn" text,
	CONSTRAINT aws_wafregional_web_acls_pk PRIMARY KEY(account_id,region,id),
	UNIQUE(cq_id)
);
CREATE TABLE IF NOT EXISTS "aws_wafregional_web_acl_rules" (
	"cq_id" uuid NOT NULL,
	"cq_meta" jsonb,
	"web_acl_cq_id" uuid,
	"priority" integer,
	"rule_id" text,
	"action" text,
	"excluded_rules" text[],
	"override_action" text,
	"type" text,
	CONSTRAINT aws_wafregional_web_acl_rules_pk PRIMARY KEY(cq_id),
	UNIQUE(cq_id),
	FOREIGN KEY (web_acl_cq_id) REFERENCES aws_wafregional_web_acls(cq_id) ON DELETE CASCADE
);
