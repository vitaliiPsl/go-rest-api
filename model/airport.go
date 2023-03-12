package model

type Airport struct {
	Id      int    `gorm:"primarykey"`
	CityId  int    `gorm:"index"`
	Name    string `gorm:"index;not null"`
	Code    string `gorm:"index;not null"`
	Opened  int
	Address string
}
