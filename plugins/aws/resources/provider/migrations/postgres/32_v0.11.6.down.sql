-- aws_redshift_snapshots
ALTER TABLE IF EXISTS aws_redshift_snapshots DROP CONSTRAINT aws_redshift_snapshots_cluster_cq_id_fkey;
ALTER TABLE IF EXISTS aws_redshift_snapshots
    ADD CONSTRAINT aws_redshift_snapshots_cluster_cq_id_fkey
    FOREIGN KEY (cluster_cq_id)
    REFERENCES aws_redshift_clusters(cq_id);
