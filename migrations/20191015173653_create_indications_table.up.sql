CREATE TABLE indications (
    id serial PRIMARY KEY,
    device_id bigint NOT NULL,
    created_at TIMESTAMP NOT NULL,
    data jsonb
);