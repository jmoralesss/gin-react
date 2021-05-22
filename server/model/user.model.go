package model

import "time"

type User struct {
	ID         uint      `db:"ID"`
	Email      string    `db:"Email"`
	Password   string    `db:"Password"`
	InsertDate time.Time `gorm:"column:InsertDate"`
}

func (User) TableName() string {
	return "User"
}
