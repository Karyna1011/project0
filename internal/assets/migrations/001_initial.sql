-- +migrate Up

CREATE TABLE person
(
    id        bigserial    not null,
    name      text not null,
    duration  bigint          not null,
    completed bool         not null
);

-- +migrate Down

DROP TABLE person;