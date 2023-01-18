package models

import "time"

type Customer struct {
	Id         int        `json:"id,omitempty" db:"Id"`
	Name       string     `json:"name" db:"Name"`
	Email      string     `json:"email" db:"Email"`
	Created_at *time.Time `json:"-" db:"Created_at"`
	Updated_at *time.Time `json:"-" db:"Updated_at"`
}
