CREATE TABLE IF NOT EXISTS province (
    id int4 NOT NULL,
    code varchar(18) NOT NULL,
    "name" varchar NOT NULL,
    postal_codes _int4 NULL,
    CONSTRAINT province_pk PRIMARY KEY (id)
);