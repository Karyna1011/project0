-- +migrate Up

CREATE TABLE debtor
(
    id        bigserial        not null ,
    address      text         not null   UNIQUE,
    PRIMARY KEY (id)
);

-- +migrate Down

DROP TABLE debtor;
