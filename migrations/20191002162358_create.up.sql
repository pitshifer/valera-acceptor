CREATE TABLE devices (
    id serial PRIMARY KEY,
    mac_address macaddr UNIQUE NOT NULL,
    reg_at TIMESTAMP NOT NULL
);