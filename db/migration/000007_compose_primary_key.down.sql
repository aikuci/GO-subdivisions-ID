ALTER TABLE "village" DROP CONSTRAINT village_pkey CASCADE,
    ADD CONSTRAINT village_pk PRIMARY KEY (id, id_district, id_city, id_province);
ALTER TABLE "district" DROP CONSTRAINT district_pkey CASCADE,
    ADD CONSTRAINT district_pk PRIMARY KEY (id, id_city, id_province);
ALTER TABLE "city" DROP CONSTRAINT city_pkey CASCADE,
    ADD CONSTRAINT city_pk PRIMARY KEY (id, id_province);