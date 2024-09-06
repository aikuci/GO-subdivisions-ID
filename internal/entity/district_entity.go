package entity

import "github.com/lib/pq"

type District struct {
	Base
	CityID      int           `gorm:"column:id_city;primaryKey;autoIncrement:false"`
	ProvinceID  int           `gorm:"column:id_province;primaryKey;autoIncrement:false"`
	Code        string        `gorm:"column:code;size:18"`
	Name        string        `gorm:"column:name"`
	PostalCodes pq.Int64Array `gorm:"column:postal_codes;type:int4[]"`
	City        City
	Province    Province
	Villages    []Village `gorm:"foreignKey:id_district,id_city,id_province"`
}

func (p *District) TableName() string {
	return "districts"
}
