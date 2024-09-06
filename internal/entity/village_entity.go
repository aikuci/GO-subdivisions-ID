package entity

import "github.com/lib/pq"

type Village struct {
	Base
	DistrictID  int           `gorm:"column:id_district;primaryKey;autoIncrement:false"`
	CityID      int           `gorm:"column:id_city;primaryKey;autoIncrement:false"`
	ProvinceID  int           `gorm:"column:id_province;primaryKey;autoIncrement:false"`
	Code        string        `gorm:"column:code;size:18"`
	Name        string        `gorm:"column:name"`
	PostalCodes pq.Int64Array `gorm:"column:postal_codes;type:int4[]"`
	District    District
	City        City
	Province    Province
}

func (p *Village) TableName() string {
	return "villages"
}
