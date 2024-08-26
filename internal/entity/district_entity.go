package entity

import "github.com/lib/pq"

type District struct {
	ID          int           `gorm:"primaryKey;autoIncrement:false"`
	IDProvince  int           `gorm:"column:id_province;primaryKey;autoIncrement:false"`
	IDCity      int           `gorm:"column:id_city;primaryKey;autoIncrement:false"`
	Code        string        `gorm:"column:code;size:18"`
	Name        string        `gorm:"column:name"`
	PostalCodes pq.Int64Array `gorm:"column:postal_codes;type:int4[]"`
}

func (p *District) TableName() string {
	return "district"
}
