package domain

import "time"

type Organization struct {
	Id          int64      `db:"id" json:"id"`
	UserId      int64      `db:"user_id" json:"user_id"`
	Name        string     `db:"name" json:"name"`
	Description string     `db:"description" json:"description"`
	City        string     `db:"city" json:"city"`
	Address     string     `db:"address" json:"address"`
	Lat         float64    `db:"lat" json:"lat"`
	Lon         float64    `db:"lon" json:"lon"`
	CreatedDate time.Time  `db:"created_date" json:"created_date"`
	UpdatedDate time.Time  `db:"updated_date" json:"updated_date"`
	DeletedDate *time.Time `db:"deleted_date" json:"deleted_date"`
}
