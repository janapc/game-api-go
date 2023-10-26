CREATE DATABASE
    IF NOT EXISTS `gameapi` DEFAULT CHARACTER SET utf8 COLLATE utf8_general_ci;

CREATE TABLE
    IF NOT EXISTS games(
        id varchar(36) PRIMARY KEY,
        name VARCHAR(175) NOT NULL,
        description VARCHAR(255) NOT NULL,
        release_date VARCHAR(55) NOT NULL,
        created_at DATETIME NOT NULL
    );