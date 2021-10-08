ALTER TABLE "gcp_compute_instance"
    ADD COLUMN IF NOT EXISTS post_key_revocation_action_type text;