package model

type Employee struct {
	ID       int64  `gorm:"PRIMARY_KEY; AUTO_INCREMENT" json:"id"`
	Username string `gorm:"NOT NULL; UNIQUE" json:"username"`
	Human
	Email    string `gorm:"NOT NULL; UNIQUE" json:"email"`
	Password string `gorm:"NOT NULL" json:"password"`
	Enabled  bool   `gorm:"NOT NULL; DEFAULT:true" json:"enabled"`
	DateModel
	Roles []Role `gorm:"many2many:employee_roles"`
}
