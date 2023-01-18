package models

import "time"

type Menu struct {
	Id               string      `json:"id" db:"Id"`
	Name             string      `json:"name" db:"Name"`
	Type             string      `json:"type" db:"Type"`
	Spice_level      string      `json:"spiceLevel" db:"Spice_level"`
	Available_on     interface{} `json:"-" db:"Available_on"`
	Is_vegan         bool        `json:"isVegan" db:"Is_vegan"`
	Is_available     bool        `json:"isAvailable" db:"Is_available"`
	Cost             int         `json:"cost" db:"Cost"`
	Preparation_time int         `json:"preparationTime" db:"Preparation_time"`
	Created_at       *time.Time  `json:"-" db:"Created_at"`
	Updated_at       *time.Time  `json:"-" db:"Updated_at"`
}
