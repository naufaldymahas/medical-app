package model

type Human struct {
	Firstname string `gorm:"NOT NULL" json:"firstname"`
	Lastname  string `gorm:"NOT NULL" json:"lastname"`
}
