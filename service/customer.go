package service

import (
	"database/sql"
	"errors"
	"log"
	"strconv"
	"time"

	"github.com/mahal007/restrauntService/models"

	"github.com/jmoiron/sqlx"
)

func CreateNewCustomerService(db *sqlx.DB, customerData *models.Customer) (customerId interface{}, err error) {
	sqlStrSelect := "select * from customer where email=?"
	sqlStrForInsert := "insert into customer(name, email) values(?,?)"
	customer := models.Customer{}
	err = db.Get(&customer, sqlStrSelect, customerData.Email)
	log.Println(customer.Id, err)
	if err != sql.ErrNoRows && err != nil {
		log.Printf("Query failed %s err = %s", customerData.Email, err.Error())
		return nil, err
	}
	if err == nil {
		return nil, errors.New("user already exists")
	}
	ret, err := db.Exec(sqlStrForInsert, customerData.Name, customerData.Email)
	if err != nil {
		log.Printf("Cannot insert data into customer table for %s, err = %s", customerData.Email, err.Error())
		return nil, err
	}
	customerId, err = ret.LastInsertId()
	if err != nil {
		log.Printf("Failed to get last inserted id %s, err = %s", customerData.Email, err.Error())
	}
	return customerId, nil
}

func CreateNewOrderService(db *sqlx.DB, orderData *models.Orders) (order interface{}, err error) {
	orderItemData := []models.OrderItem{}
	sqlStrOrderInsert := `insert into orders(customer_id, order_date) values(?,?)`
	sqlStrOrderItemInsert := `insert into orderitem(order_id, item_id) values(:order_id, :item_id)`
	ret, err := db.Exec(sqlStrOrderInsert, orderData.Customer_id, time.Now())
	if err != nil {
		log.Printf("failed to insert data into itemTable %d err = %s", orderData.Customer_id, err.Error())
		return nil, err
	}
	orderId, err := ret.LastInsertId()
	if err != nil {
		log.Printf("failed to get the inserted id %s", err.Error())
		return nil, err
	}
	for _, v := range orderData.Orderitem {
		if v != 0 {
			temp := models.OrderItem{}
			temp.Order_id = int(orderId)
			temp.Item_id = v
			orderItemData = append(orderItemData, temp)
		}
	}
	retI, err := db.NamedExec(sqlStrOrderItemInsert, orderItemData)
	if err != nil {
		log.Printf("failed to insert data into itemTable %d err = %s", orderId, err.Error())
		return nil, err
	}
	_, err = retI.LastInsertId()
	if err != nil {
		log.Printf("failed to get the inserted id %s", err.Error())
		return nil, err
	}
	return orderItemData, nil
}

func GetAllOrderByCustomerIdService(db *sqlx.DB, customerId int) (orderData interface{}, err error) {
	sqlStrOrder := "Select distinct o.Order_date, o.status, m.name, m.is_available  from orderitem  join orders o join menu m where o.customer_id=?"
	err = db.Select(&orderData, sqlStrOrder, customerId)
	if err != sql.ErrNoRows && err != nil {
		log.Printf("Query failed %s err = %s", strconv.Itoa(customerId), err.Error())
		return nil, err
	}
	return orderData, nil
}

func GetOrderByIdService(db *sqlx.DB, orderId int) (orderData interface{}, err error) {
	sqlStrOrder := "Select o.id,o.status, m.name, m.is_available  from orderitem join orders o join menu m where order_id=?"
	err = db.Get(&orderData, sqlStrOrder, orderId)
	if err != sql.ErrNoRows && err != nil {
		log.Printf("Query failed %s err = %s", strconv.Itoa(orderId), err.Error())
		return nil, err
	}
	return orderData, nil
}
