ALTER TABLE "village" DROP CONSTRAINT village_pk CASCADE,
    ADD PRIMARY KEY (id, id_district, id_city, id_province);
ALTER TABLE "district" DROP CONSTRAINT district_pk CASCADE,
    ADD PRIMARY KEY (id, id_city, id_province);
ALTER TABLE "city" DROP CONSTRAINT city_pk CASCADE,
    ADD PRIMARY KEY (id, id_province);