package contracts

import (
	model "kis/internal/models"
	"net/http"
)

type response struct {
	statusCode int
	body       map[string]any
}

func (r *response) Values() (int, map[string]any) {
	return r.statusCode, r.body
}

type responseOption func(*response)

func withBadRequest() responseOption {
	return withError(http.StatusBadRequest, "Неправильные параметры запроса")
}

func withInternalServerError() responseOption {
	return withError(http.StatusInternalServerError, "Ошибка на стороне сервера")
}

func withError(statusCode int, errorMsg string) responseOption {
	return func(r *response) {
		r.statusCode = statusCode
		r.body["status"] = statusCode
		r.body["error_msg"] = errorMsg
	}
}

func withSuccess(statusCode int, msg string) responseOption {
	return func(r *response) {
		r.statusCode = statusCode
		r.body["status"] = "success"
		r.body["message"] = msg
	}
}

func withResult(res []model.VacancyDTO, page int) responseOption {
	return func(r *response) {
		r.statusCode = 200
		r.body["page"] = page
		r.body["result"] = res
		r.body["status"] = "success"
	}
}

func newResponse(options ...responseOption) *response {
	res := response{body: make(map[string]any)}

	for _, opt := range options {
		opt(&res)
	}

	return &res
}
