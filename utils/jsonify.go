package utils

import (
	"net/http"

	"github.com/go-chi/render"
)

type JsonResponse struct {
	HTTPStatusCode int         `json:"statusCode"` // http response status code
	StatusText     string      `json:"status"`
	Data           interface{} `json:"data"`
}

func (j *JsonResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, j.HTTPStatusCode)
	return nil
}

func Jsonify(httpStatus int, statusText string, data interface{}) render.Renderer {
	return &JsonResponse{
		HTTPStatusCode: httpStatus,
		StatusText:     statusText,
		Data:           data,
	}
}
