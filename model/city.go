package model

type City struct {
	Id        int    `gorm:"primarykey"`
	Name      string `gorm:"index;not null"`
	Country   string `gorm:"index;not null"`
	Region    string
	Geolat    float32
	Geolng    float32
	Landmarks []Landmark `gorm:"foreignkey:CityId;constraint:OnDelete:CASCADE"`
	Airports  []Airport  `gorm:"foreignkey:CityId;constraint:OnDelete:CASCADE"`
}
