CREATE TABLE indications (
    id serial PRIMARY KEY,
    device_id bigint NOT NULL REFERENCES devices(id) ON DELETE CASCADE,
    created_at TIMESTAMP NOT NULL,
    data jsonb
);