package model

type Role struct {
	ID      int64  `gorm:"PRIMARY_KEY; AUTO_INCREMENT" json:"id"`
	Name    string `gorm:"NOT NULL; UNIQUE" json:"name"`
	Detail  string `gorm:"NOT NULL" json:"detail"`
	Enabled bool   `gorm:"NOT NULL; DEFAULT:true" json:"enabled"`
	DateModel
}
