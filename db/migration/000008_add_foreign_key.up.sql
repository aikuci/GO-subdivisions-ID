ALTER TABLE district
ADD CONSTRAINT fk_district_city FOREIGN KEY (id_city, id_province) REFERENCES city (id, id_province);
ALTER TABLE village
ADD CONSTRAINT fk_village_city FOREIGN KEY (id_city, id_province) REFERENCES city (id, id_province);
ALTER TABLE village
ADD CONSTRAINT fk_village_district FOREIGN KEY (id_district, id_city, id_province) REFERENCES district (id, id_city, id_province);