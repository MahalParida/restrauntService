package handlers

import (
	"net/http"

	"github.com/mahal007/restrauntService/models"
	"github.com/mahal007/restrauntService/service"
	"github.com/mahal007/restrauntService/utils"

	"github.com/go-chi/render"
	"github.com/jmoiron/sqlx"
)

func GetMenu(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		menuData := &models.Menu{}
		if err := render.DecodeJSON(r.Body, menuData); err != nil {
			render.Render(w, r, utils.ErrBadRequest(err))
			return
		}
		data, err := service.GetMenuService(db, menuData)
		if err != nil {
			render.Render(w, r, utils.ErrInternalServerError(err))
			return
		}
		render.Status(r, 200)
		render.JSON(w, r, utils.Jsonify(200, "Menu items", data))
	}
}

func AddItemToMenu(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		menuData := &models.Menu{}
		if err := render.DecodeJSON(r.Body, menuData); err != nil {
			render.Render(w, r, utils.ErrBadRequest(err))
			return
		}
		data, err := service.AddItemToMenuService(db, menuData)
		if err != nil {
			if err.Error() == "item already exists" {
				render.Render(w, r, utils.ErrBadRequest(err))
			}
			render.Render(w, r, utils.ErrInternalServerError(err))
		}
		render.Status(r, 201)
		render.JSON(w, r, utils.Jsonify(201, "item added to menu", data))
	}
}
