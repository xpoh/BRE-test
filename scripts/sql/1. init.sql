
CREATE TABLE service (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(256),
    content VARCHAR
);
