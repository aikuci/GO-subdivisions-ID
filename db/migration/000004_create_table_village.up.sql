CREATE TABLE IF NOT EXISTS villages (
    id int4 NOT NULL,
    id_province int4 NOT NULL,
    id_city int4 NOT NULL,
    id_district int4 NOT NULL,
    code varchar(18) NOT NULL,
    "name" varchar NOT NULL,
    postal_codes _int4 NULL,
    CONSTRAINT villages_pk PRIMARY KEY (id, id_province, id_city, id_district),
    CONSTRAINT fk_villages_districts FOREIGN KEY (id_district, id_city, id_province) REFERENCES districts (id, id_city, id_province),
    CONSTRAINT fk_villages_cities FOREIGN KEY (id_city, id_province) REFERENCES cities (id, id_province)
);
CREATE INDEX IF NOT EXISTS villages_districts_idx ON villages (id_district);
CREATE INDEX IF NOT EXISTS villages_cities_idx ON villages (id_city);
CREATE INDEX IF NOT EXISTS villages_provinces_idx ON villages (id_province);