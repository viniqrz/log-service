package exception

import "net/http"

const NotFoundErrorMessage = "NOT FOUND"
const InternalServerErrorMessage = "INTERNAL SERVER ERROR"
const BadRequestErrorMessage = "BAD REQUEST"

func GetStatusCode(err error) int {
	switch err.Error() {
	case NotFoundErrorMessage:
		return http.StatusNotFound
	case BadRequestErrorMessage:
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}
