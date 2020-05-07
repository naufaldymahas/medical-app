package model

type User struct {
	ID       int64  `gorm:"PRIMARY_KEY; AUTO_INCREMENT" json:"id"`
	Username string `gorm:"NOT NULL; UNIQUE" json:"username"`
	Email    string `gorm:"NOT NULL; UNIQUE" json:"email"`
	Password string `gorm:"NOT NULL" json:"password"`
	Human
	Phone   string `gorm:"NOT NULL; UNIQUE" json:"phone"`
	Enabled bool   `gorm:"NOT NULL; DEFAULT:true" json:"enabled"`
	DateModel
}
