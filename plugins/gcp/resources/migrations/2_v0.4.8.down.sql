ALTER TABLE "gcp_compute_instances"
    ADD COLUMN IF NOT EXISTS post_key_revocation_action_type text;