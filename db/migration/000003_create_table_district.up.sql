CREATE TABLE IF NOT EXISTS district (
    id serial NOT NULL,
    id_province integer NOT NULL references province(id),
    id_city integer NOT NULL references city(id),
    code smallserial NOT NULL,
    name varchar NOT NULL,
    postal_codes integer ARRAY,
    CONSTRAINT district_pk PRIMARY KEY (id),
    CONSTRAINT district_unique UNIQUE (code)
);