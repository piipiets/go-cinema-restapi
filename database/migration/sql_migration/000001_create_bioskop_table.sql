-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE cinema (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    location VARCHAR(255) NOT NULL,
    rate FLOAT NOT NULL
);

-- +migrate StatementEnd

-- +migrate Down
-- +migrate StatementBegin

DROP TABLE IF EXISTS cinema;

-- +migrate StatementEnd