CREATE TABLE IF NOT EXISTS providers
(
    source TEXT NOT NULL,
    name TEXT NOT NULL,
    version TEXT NOT NULL,
    v_major INT NOT NULL,
    v_minor INT NOT NULL,
    v_patch INT NOT NULL,
    v_pre TEXT NOT NULL,
    v_meta TEXT NOT NULL,
    CONSTRAINT "providers_id" PRIMARY KEY (source, name)
);
