package service

import (
	"database/sql"
	"errors"
	"log"
	"time"

	"github.com/mahal007/restrauntService/models"

	"github.com/jmoiron/sqlx"
)

func GetMenuService(db *sqlx.DB, menuData *models.Menu) (menu interface{}, err error) {

	//Todo : Dynamic search can be added
	day := time.Now().Weekday().String()
	menuItems := []models.Menu{}
	sqlStrMenu := "select * from menu where Available_on like ?"
	err = db.Select(&menuItems, sqlStrMenu, "%"+day+"%")
	if err != nil {
		log.Panicf("Failed to get menu items %s", err)
		return nil, err
	}
	return menuItems, err
}

func AddItemToMenuService(db *sqlx.DB, menuData *models.Menu) (menu interface{}, err error) {
	menuItem := models.Menu{}
	sqlStrMenuItemSelect := `select * from menu where name=?`
	sqlStrMenuItemInsert := `insert into menu(name,
										type,
										spice_level,
										available_on,
										is_vegan,
										is_available,
										cost,
										preparation_time
										 values(?,?,?,?,?,?,?,?))`
	err = db.Get(&menuItem, sqlStrMenuItemSelect, menuData.Name)
	if err != sql.ErrNoRows && err != nil {
		log.Printf("Query failed %s err = %s", menuData.Name, err.Error())
		return nil, err
	}
	if err == nil {
		log.Printf("item already exits in table %s", menuData.Name)
		return nil, errors.New("item already exists")
	}
	ret, err := db.Exec(sqlStrMenuItemInsert, menuData)
	if err != nil {
		log.Printf("Cannot insert data into customer table for %s, err = %s", menuData.Name, err.Error())
		return nil, err
	}
	menu, err = ret.LastInsertId()
	if err != nil {
		log.Printf("Failed to get the last inserted id from menu table %s err = %s", menuData.Name, err.Error())
	}
	return menu, nil
}
