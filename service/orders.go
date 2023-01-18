package service

import (
	"log"
	"strconv"

	"github.com/jmoiron/sqlx"
)

func GetAllOrder(db *sqlx.DB) (orderdata interface{}, err error) {
	sqlStrSelect := `select * from orders where status like=` // TODO : Should add pagination
	orders := make([]interface{}, 0)
	err = db.Select(&orders, sqlStrSelect, "%"+"pending"+"%")
	if err != nil {
		log.Printf("error while fetching orders from the table %s", err.Error())
		return nil, err
	}
	return orders, nil
}

func UpdateOrder(db *sqlx.DB, status map[string]string, orderId int) (order interface{}, err error) {
	sqlStrUpdate := `update orders set status=? where orderId=?`
	ret, err := db.Exec(sqlStrUpdate, status["status"], orderId)
	if err != nil {
		log.Printf("error occur while updating orders status %s err = %s", strconv.Itoa(orderId), err.Error())
		return nil, err
	}
	order, err = ret.RowsAffected()
	if err != nil {
		log.Printf("error occured while geting affected %s", err.Error())
		return nil, err
	}
	return order, err
}
