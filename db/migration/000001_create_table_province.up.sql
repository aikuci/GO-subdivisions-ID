CREATE TABLE IF NOT EXISTS province (
    id serial NOT NULL,
    code smallserial NOT NULL,
    name varchar NOT NULL,
    postal_codes integer ARRAY,
    CONSTRAINT province_pk PRIMARY KEY (id),
    CONSTRAINT province_unique UNIQUE (code)
);