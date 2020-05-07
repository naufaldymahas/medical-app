package model

import "time"

type DateModel struct {
	CreatedDate time.Time `gorm:"DEFAULT:CURRENT_TIMESTAMP;NOT NULL" json:"created_date"`
	UpdatedDate time.Time `gorm:"DEFAULT:CURRENT_TIMESTAMP;NOT NULL" json:"updated_date"`
}
