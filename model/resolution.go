package model

type Resolution struct {
	ID   uint 	`gorm:"primary_key"`
	Code string
	Bar  []Bar
}
