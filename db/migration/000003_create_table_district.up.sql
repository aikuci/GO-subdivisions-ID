CREATE TABLE IF NOT EXISTS districts (
    id int4 NOT NULL,
    id_province int4 NOT NULL,
    id_city int4 NOT NULL,
    code varchar(18) NOT NULL,
    "name" varchar NOT NULL,
    postal_codes _int4 NULL,
    CONSTRAINT districts_pk PRIMARY KEY (id, id_province, id_city),
    CONSTRAINT fk_districts_cities FOREIGN KEY (id_city, id_province) REFERENCES cities (id, id_province)
);
CREATE INDEX IF NOT EXISTS districts_cities_idx ON districts (id_city);
CREATE INDEX IF NOT EXISTS districts_provinces_idx ON districts (id_province);