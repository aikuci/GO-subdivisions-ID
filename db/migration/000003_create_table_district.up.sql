CREATE TABLE IF NOT EXISTS district (
    id int4 NOT NULL,
    id_province int4 NOT NULL,
    id_city int4 NOT NULL,
    code varchar(18) NOT NULL,
    "name" varchar NOT NULL,
    postal_codes _int4 NULL,
    CONSTRAINT district_pk PRIMARY KEY (id, id_province, id_city),
    CONSTRAINT fk_district_city FOREIGN KEY (id_city, id_province) REFERENCES city (id, id_province)
);
CREATE INDEX IF NOT EXISTS district_province_idx ON district (id_province);
CREATE INDEX IF NOT EXISTS district_city_idx ON district (id_city);