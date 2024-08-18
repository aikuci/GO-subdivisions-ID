CREATE TABLE IF NOT EXISTS village (
    id serial NOT NULL,
    id_province integer NOT NULL references province(id),
    id_city integer NOT NULL references city(id),
    id_district integer NOT NULL references district(id),
    code smallserial NOT NULL,
    name varchar NOT NULL,
    postal_codes integer ARRAY,
    CONSTRAINT village_pk PRIMARY KEY (id),
    CONSTRAINT village_unique UNIQUE (code)
);