CREATE INDEX IF NOT EXISTS district_province_idx ON district (id_province);
CREATE INDEX IF NOT EXISTS district_city_idx ON district (id_city);
CREATE INDEX IF NOT EXISTS village_province_idx ON village (id_province);
CREATE INDEX IF NOT EXISTS village_city_idx ON village (id_city);
CREATE INDEX IF NOT EXISTS village_district_idx ON village (id_district);