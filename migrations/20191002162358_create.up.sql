CREATE TABLE devices (
    `uuid` CHAR(36) NOT NULL,
    `last_activity` DATETIME NOT NULL DEFAULT NOW(),
    PRIMARY KEY(`uuid`)
) ENGINE = InnoDB;