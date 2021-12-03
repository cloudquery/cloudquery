DROP TABLE IF EXISTS public.aws_emr_clusters;

CREATE TABLE public.aws_emr_clusters (
    cq_id uuid,
    meta jsonb,
    account_id text NOT NULL,
    region text,
    arn text,
    id text NOT NULL,
    name text,
    normalized_instance_hours integer,
    outpost_arn text,
    status_state text,
    status_state_change_reason_code text,
    status_state_change_reason_message text,
    status_timeline_creation_date_time timestamp without time zone,
    status_timeline_end_date_time timestamp without time zone,
    status_timeline_ready_date_time timestamp without time zone,
    vpc_id text
);

ALTER TABLE ONLY public.aws_emr_clusters
    ADD CONSTRAINT aws_emr_clusters_cq_id_key UNIQUE (cq_id);

ALTER TABLE ONLY public.aws_emr_clusters
    ADD CONSTRAINT aws_emr_clusters_pk PRIMARY KEY (account_id, id);
