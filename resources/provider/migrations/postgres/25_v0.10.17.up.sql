-- Resource: redshift.event_subscriptions
CREATE TABLE IF NOT EXISTS "aws_redshift_event_subscriptions" (
	"cq_id" uuid NOT NULL,
	"cq_meta" jsonb,
	"account_id" text,
	"region" text,
	"id" text,
	"customer_aws_id" text,
	"enabled" boolean,
	"event_categories_list" text[],
	"severity" text,
	"sns_topic_arn" text,
	"source_ids_list" text[],
	"source_type" text,
	"status" text,
	"subscription_creation_time" timestamp without time zone,
	"tags" jsonb,
	CONSTRAINT aws_redshift_event_subscriptions_pk PRIMARY KEY(account_id,id),
	UNIQUE(cq_id)
);
