package entity

import "github.com/lib/pq"

type City struct {
	Base
	ProvinceID  int           `gorm:"column:id_province;primaryKey;autoIncrement:false"`
	Code        string        `gorm:"column:code;size:18"`
	Name        string        `gorm:"column:name"`
	PostalCodes pq.Int64Array `gorm:"column:postal_codes;type:int4[]"`
	Province    Province
	Districts   []District `gorm:"foreignKey:id_city,id_province"`
	Villages    []Village  `gorm:"foreignKey:id_city,id_province"`
}

func (p *City) TableName() string {
	return "cities"
}
