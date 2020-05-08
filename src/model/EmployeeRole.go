package model

type EmployeeRole struct {
	EmployeeID int64 `gorm:"NOT NULL" json:"employee_id"`
	RoleID     int64 `gorm:"NOT NULL" json:"role_id"`
}
