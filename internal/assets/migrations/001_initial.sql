-- +migrate Up

CREATE TABLE person
(
    id        bigserial    not null,
    address      text         not null UNIQUE,
    PRIMARY KEY (id)
);

-- +migrate Down

DROP TABLE person;
