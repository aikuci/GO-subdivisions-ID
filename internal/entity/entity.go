package entity

type Base struct {
	ID int `gorm:"column:id;primaryKey;autoIncrement:false"`
}
