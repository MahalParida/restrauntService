package models

import "time"

type Orders struct {
	Id          int        `json:"id" db:"Id"`
	Customer_id int        `json:"customerId" db:"Customer_id"`
	Order_date  time.Time  `json:"orderDate" db:"Order_id"`
	Status      string     `json:"status" db:"Status"`
	Orderitem   []int      `json:"-" db:"-"`
	Created_at  *time.Time `json:"-" db:"Created_at"`
	Updated_at  *time.Time `json:"-" db:"Updated_at"`
}

type OrderItem struct {
	Order_id int
	Item_id  int
}
