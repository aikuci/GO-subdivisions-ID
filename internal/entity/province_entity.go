package entity

import "github.com/lib/pq"

type Province struct {
	Base
	Code        string        `gorm:"column:code;size:18"`
	Name        string        `gorm:"column:name"`
	PostalCodes pq.Int64Array `gorm:"column:postal_codes;type:int4[]"`
}

func (p *Province) TableName() string {
	return "province"
}
