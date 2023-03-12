package model

type City struct {
	Id          int    `gorm:"primarykey"`
	Name        string `gorm:"index;not null"`
	Country     string `gorm:"index;not null"`
	Region      string
	Geolat      float32
	Geolng      float32
}
