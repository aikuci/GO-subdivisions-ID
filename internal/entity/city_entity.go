package entity

import "github.com/lib/pq"

type City struct {
	Base
	Code        string        `gorm:"column:code;size:18"`
	Name        string        `gorm:"column:name"`
	PostalCodes pq.Int64Array `gorm:"column:postal_codes;type:int4[]"`
}

func (p *City) TableName() string {
	return "city"
}
