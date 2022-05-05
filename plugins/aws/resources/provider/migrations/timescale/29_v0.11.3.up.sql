-- aws_redshift_clusters: update ARN to correct values and change PK to arn field
ALTER TABLE IF EXISTS aws_redshift_clusters DROP CONSTRAINT aws_redshift_clusters_pk;
UPDATE aws_redshift_clusters SET arn = format('arn:aws:redshift:%s:%s:cluster:%s', region, account_id, id);
ALTER TABLE IF EXISTS aws_redshift_clusters ADD CONSTRAINT aws_redshift_clusters_pk PRIMARY KEY (cq_fetch_date,arn);

-- aws_redshift_snapshots: add cluster_cq_id and a corresponding FK constraint
ALTER TABLE IF EXISTS aws_redshift_snapshots ADD COLUMN IF NOT EXISTS cluster_cq_id uuid;

UPDATE aws_redshift_snapshots s
SET cluster_cq_id = cl.cq_id
FROM aws_redshift_clusters cl
WHERE cl.id = s.cluster_identifier AND cl.cluster_create_time = s.cluster_create_time AND cl.cq_fetch_date = s.cq_fetch_date;

ALTER TABLE IF EXISTS aws_redshift_snapshots
    ADD CONSTRAINT aws_redshift_snapshots_cluster_cq_id_fkey
    FOREIGN KEY (cq_fetch_date, cluster_cq_id)
    REFERENCES aws_redshift_clusters(cq_fetch_date, cq_id);

-- aws_redshift_snapshots: add ARN, change PK to ARN and fill values
ALTER TABLE IF EXISTS aws_redshift_snapshots DROP CONSTRAINT aws_redshift_snapshots_pk;
ALTER TABLE IF EXISTS aws_redshift_snapshots ADD COLUMN IF NOT EXISTS arn text;

UPDATE aws_redshift_snapshots s
SET arn = format('arn:aws:redshift:%s:%s:snapshot:%s/%s', cl.region, cl.account_id, cl.id, s.snapshot_identifier)
FROM aws_redshift_clusters cl
WHERE cl.cq_id = s.cluster_cq_id AND cl.cq_fetch_date = s.cq_fetch_date;

ALTER TABLE IF EXISTS aws_redshift_snapshots ADD CONSTRAINT aws_redshift_snapshots_pk PRIMARY KEY (cq_fetch_date,arn);

-- Resource: redshift.event_subscriptions
ALTER TABLE IF EXISTS aws_redshift_event_subscriptions DROP CONSTRAINT aws_redshift_event_subscriptions_pk;
ALTER TABLE IF EXISTS aws_redshift_event_subscriptions ADD COLUMN IF NOT EXISTS arn text;
UPDATE aws_redshift_event_subscriptions SET arn = format('arn:aws:redshift:%s:%s:eventsubscription:%s', region, account_id, id);
ALTER TABLE IF EXISTS aws_redshift_event_subscriptions ADD CONSTRAINT aws_redshift_event_subscriptions_pk PRIMARY KEY (cq_fetch_date,arn);

-- aws_redshift_subnet_groups: change PK to ARN and update values
ALTER TABLE IF EXISTS aws_redshift_subnet_groups DROP CONSTRAINT aws_redshift_subnet_groups_pk;
UPDATE aws_redshift_subnet_groups SET arn = format('arn:aws:redshift:%s:%s:subnetgroup:%s', region, account_id, cluster_subnet_group_name);
ALTER TABLE IF EXISTS aws_redshift_subnet_groups ADD CONSTRAINT aws_redshift_subnet_groups_pk PRIMARY KEY (cq_fetch_date, arn);
