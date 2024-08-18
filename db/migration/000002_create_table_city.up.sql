CREATE TABLE IF NOT EXISTS city (
    id serial NOT NULL,
    id_province integer NOT NULL references province(id),
    code smallserial NOT NULL,
    name varchar NOT NULL,
    postal_codes integer ARRAY,
    CONSTRAINT city_pk PRIMARY KEY (id),
    CONSTRAINT city_unique UNIQUE (code)
);