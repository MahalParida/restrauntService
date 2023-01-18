package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/mahal007/restrauntService/models"
	"github.com/mahal007/restrauntService/service"
	"github.com/mahal007/restrauntService/utils"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/jmoiron/sqlx"
)

func CreateNewCustomer(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		c := &models.Customer{}
		if err := render.DecodeJSON(r.Body, c); err != nil {
			render.Status(r, 400)
			render.Render(w, r, utils.ErrBadRequest(err))
			return
		}
		data, err := service.CreateNewCustomerService(db, c)
		if err != nil {
			if err.Error() == "user already exists" {
				render.Status(r, 400)
				render.Render(w, r, utils.ErrBadRequest(err))
				return
			}
			render.Status(r, 500)
			render.Render(w, r, utils.ErrInternalServerError(err))
			return
		}
		render.Status(r, 201)
		render.JSON(w, r, utils.Jsonify(201, "user created", data))
	}
}

func CreateNewOrder(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		o := &models.Orders{}
		if err := render.DecodeJSON(r.Body, o); err != nil {
			render.Status(r, 400)
			render.Render(w, r, utils.ErrBadRequest(err))
			return
		}
		data, err := service.CreateNewOrderService(db, o)
		if err != nil {
			render.Status(r, 500)
			render.Render(w, r, utils.ErrInternalServerError(err))
			return
		}
		render.JSON(w, r, utils.Jsonify(201, "New order created", data))
	}
}

func GetAllOrderByCustomerId(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		customerId := chi.URLParam(r, "customerId")
		if customerId == "" {
			render.Status(r, 400)
			render.Render(w, r, utils.ErrBadRequest(errors.New("invalid customerId")))
			return
		}
		customer, _ := strconv.Atoi(customerId)
		data, err := service.GetAllOrderByCustomerIdService(db, customer)
		if err != nil {
			render.Status(r, 500)
			render.Render(w, r, utils.ErrInternalServerError(err))
			return
		}
		render.Status(r, 200)
		render.Render(w, r, utils.Jsonify(200, "list of order", data))
	}
}

func GetOrderById(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		orderId := chi.URLParam(r, "orderId")
		if orderId == "" {
			render.Status(r, 400)
			render.Render(w, r, utils.ErrBadRequest(errors.New("invalid orderId")))
			return
		}
		order, _ := strconv.Atoi(orderId)
		data, err := service.GetOrderByIdService(db, order)
		if err != nil {
			render.Status(r, 500)
			render.Render(w, r, utils.ErrInternalServerError(err))
			return
		}
		render.Status(r, 200)
		render.Render(w, r, utils.Jsonify(200, "list of order", data))

	}
}

func UpdateOrderById(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}
