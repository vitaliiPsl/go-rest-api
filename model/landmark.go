package model

type Landmark struct {
	Id          int    `gorm:"primarykey"`
	CityId      int    `gorm:"index"`
	Name        string `gorm:"index;not null"`
	Description string
	Address     string
}
