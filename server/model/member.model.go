package model

import "time"

type Member struct {
	Email      string    `gorm:"email"`
	InsertDate time.Time `gorm:"column:insertDate"`
}
