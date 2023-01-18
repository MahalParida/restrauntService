package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/mahal007/restrauntService/service"
	"github.com/mahal007/restrauntService/utils"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/jmoiron/sqlx"
)

func GetAllOrder(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := service.GetAllOrder(db)
		if err != nil {
			render.JSON(w, r, utils.ErrInternalServerError(err))
		}
		render.Status(r, 200)
		render.JSON(w, r, utils.Jsonify(200, "All pending orders", data))
	}
}

func UpdateOrder(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		customerStatus := make(map[string]string)
		orderId := chi.URLParam(r, "orderId")
		if orderId == "" {
			render.Status(r, 400)
			render.Render(w, r, utils.ErrBadRequest(errors.New("invalid orderId")))
			return
		}
		err := render.DecodeJSON(r.Body, customerStatus)
		if err != nil {
			render.Status(r, 400)
			render.Render(w, r, utils.ErrBadRequest(err))
			return
		}

		order, _ := strconv.Atoi(orderId)
		data, err := service.UpdateOrder(db, customerStatus, order)
		if err != nil {
			render.Status(r, 500)
			render.Render(w, r, utils.ErrInternalServerError(err))
			return
		}
		render.Status(r, 200)
		render.JSON(w, r, utils.Jsonify(200, "order status updated", data))
	}
}
