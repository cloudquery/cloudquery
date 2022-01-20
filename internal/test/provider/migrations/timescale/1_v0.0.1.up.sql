-- Resource: error_resource
CREATE TABLE IF NOT EXISTS "error_resource" (
	"cq_id" uuid NOT NULL,
	"cq_meta" jsonb,
	"cq_fetch_date" timestamp without time zone NOT NULL,
	CONSTRAINT error_resource_pk PRIMARY KEY(cq_fetch_date,cq_id),
	UNIQUE(cq_fetch_date,cq_id)
);

-- Resource: slow_resource
CREATE TABLE IF NOT EXISTS "slow_resource" (
	"cq_id" uuid NOT NULL,
	"cq_meta" jsonb,
	"cq_fetch_date" timestamp without time zone NOT NULL,
	"some_bool" boolean,
	CONSTRAINT slow_resource_pk PRIMARY KEY(cq_fetch_date,cq_id),
	UNIQUE(cq_fetch_date,cq_id)
);

-- Resource: very_slow_resource
CREATE TABLE IF NOT EXISTS "very_slow_resource" (
	"cq_id" uuid NOT NULL,
	"cq_meta" jsonb,
	"cq_fetch_date" timestamp without time zone NOT NULL,
	CONSTRAINT very_slow_resource_pk PRIMARY KEY(cq_fetch_date,cq_id),
	UNIQUE(cq_fetch_date,cq_id)
);
