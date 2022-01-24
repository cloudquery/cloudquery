-- Resource: error_resource
CREATE TABLE IF NOT EXISTS "error_resource" (
	"cq_id" uuid NOT NULL,
	"cq_meta" jsonb,
	CONSTRAINT error_resource_pk PRIMARY KEY(cq_id),
	UNIQUE(cq_id)
);

-- Resource: slow_resource
CREATE TABLE IF NOT EXISTS "slow_resource" (
	"cq_id" uuid NOT NULL,
	"cq_meta" jsonb,
	"some_bool" boolean,
	CONSTRAINT slow_resource_pk PRIMARY KEY(cq_id),
	UNIQUE(cq_id)
);

-- Resource: very_slow_resource
CREATE TABLE IF NOT EXISTS "very_slow_resource" (
	"cq_id" uuid NOT NULL,
	"cq_meta" jsonb,
	CONSTRAINT very_slow_resource_pk PRIMARY KEY(cq_id),
	UNIQUE(cq_id)
);
