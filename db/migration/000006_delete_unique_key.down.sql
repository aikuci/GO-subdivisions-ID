ALTER TABLE village
ADD CONSTRAINT village_unique UNIQUE (code);
ALTER TABLE district
ADD CONSTRAINT district_unique UNIQUE (code);
ALTER TABLE city
ADD CONSTRAINT city_unique UNIQUE (code);