-- aws_redshift_clusters: change PK to arn field
TRUNCATE TABLE aws_redshift_clusters CASCADE;
ALTER TABLE IF EXISTS aws_redshift_clusters DROP CONSTRAINT aws_redshift_clusters_pk;
ALTER TABLE IF EXISTS aws_redshift_clusters ADD CONSTRAINT aws_redshift_clusters_pk PRIMARY KEY (arn);

-- aws_redshift_snapshots: add cluster_cq_id and a corresponding FK constraint
ALTER TABLE IF EXISTS aws_redshift_snapshots ADD COLUMN IF NOT EXISTS cluster_cq_id uuid;
ALTER TABLE IF EXISTS aws_redshift_snapshots
    ADD CONSTRAINT aws_redshift_snapshots_cluster_cq_id_fkey
    FOREIGN KEY (cluster_cq_id)
    REFERENCES aws_redshift_clusters(cq_id);

-- aws_redshift_snapshots: add ARN, change PK to ARN
ALTER TABLE IF EXISTS aws_redshift_snapshots DROP CONSTRAINT aws_redshift_snapshots_pk;
ALTER TABLE IF EXISTS aws_redshift_snapshots ADD COLUMN IF NOT EXISTS arn text;
ALTER TABLE IF EXISTS aws_redshift_snapshots ADD CONSTRAINT aws_redshift_snapshots_pk PRIMARY KEY (arn);

-- aws_redshift_event_subscriptions: add ARN, change PK to ARN
TRUNCATE TABLE aws_redshift_event_subscriptions CASCADE;
ALTER TABLE IF EXISTS aws_redshift_event_subscriptions DROP CONSTRAINT aws_redshift_event_subscriptions_pk;
ALTER TABLE IF EXISTS aws_redshift_event_subscriptions ADD COLUMN IF NOT EXISTS arn text;
ALTER TABLE IF EXISTS aws_redshift_event_subscriptions ADD CONSTRAINT aws_redshift_event_subscriptions_pk PRIMARY KEY (arn);

-- aws_redshift_subnet_groups: change PK to ARN
TRUNCATE TABLE aws_redshift_subnet_groups CASCADE;
ALTER TABLE IF EXISTS aws_redshift_subnet_groups DROP CONSTRAINT aws_redshift_subnet_groups_pk;
ALTER TABLE IF EXISTS aws_redshift_subnet_groups ADD CONSTRAINT aws_redshift_subnet_groups_pk PRIMARY KEY (arn);
