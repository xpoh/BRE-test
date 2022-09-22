\c db

SET ROLE user;

create schema test;

alter schema test owner to user;

CREATE TABLE test.service (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(256),
    content VARCHAR
);
