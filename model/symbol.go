package model

type Symbol struct {
	ID   uint 	`gorm:"primary_key"`
	Code string
}
