package entity

type Province struct {
	ID          int    `gorm:"column:id;primaryKey;autoIncrement:false"`
	Code        string `gorm:"column:code;size:18"`
	Name        string `gorm:"column:name"`
	PostalCodes []int  `gorm:"column:postal_codes;type:int4[]"`
}

func (p *Province) TableName() string {
	return "province"
}
