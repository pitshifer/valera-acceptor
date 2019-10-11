CREATE TABLE devices (
    id serial PRIMARY KEY,
    mac_address macaddr NOT NULL,
    reg_at TIMESTAMP NOT NULL
);