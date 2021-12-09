package model

type Demo struct {
	ID *int `gorm:"id" json:"id"`
	Name string `gorm:"name" json:"name"`
}

func (Demo) TableName() string {
	return "springdemo"
}