CREATE TABLE IF NOT EXISTS village (
    id int4 NOT NULL,
    id_province int4 NOT NULL,
    id_city int4 NOT NULL,
    id_district int4 NOT NULL,
    code varchar(18) NOT NULL,
    "name" varchar NOT NULL,
    postal_codes _int4 NULL,
    CONSTRAINT village_pk PRIMARY KEY (id, id_province, id_city, id_district),
    CONSTRAINT fk_village_district FOREIGN KEY (id_district, id_city, id_province) REFERENCES district (id, id_city, id_province),
    CONSTRAINT fk_village_city FOREIGN KEY (id_city, id_province) REFERENCES city (id, id_province)
);
CREATE INDEX IF NOT EXISTS village_province_idx ON district (id_province);
CREATE INDEX IF NOT EXISTS village_city_idx ON district (id_city);
CREATE INDEX IF NOT EXISTS village_district_idx ON village (id_district);