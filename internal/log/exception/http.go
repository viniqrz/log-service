package exception

import "net/http"

const NotFoundErrorMessage = "NOT FOUND"
const InternalServerErrorMessage = "INTERNAL SERVER ERROR"
const BadRequestErrorMessage = "BAD REQUEST"
const UnauthorizedErrorMessage = "UNAUTHORIZED"
const UnprocessableEntityErrorMessage = "UNPROCESSABLE ENTITY"

func GetStatusCode(err error) int {
	switch err.Error() {
	case NotFoundErrorMessage:
		return http.StatusNotFound
	case BadRequestErrorMessage:
		return http.StatusBadRequest
	case UnauthorizedErrorMessage:
		return http.StatusUnauthorized
	case UnprocessableEntityErrorMessage:
		return http.StatusUnprocessableEntity
	default:
		return http.StatusInternalServerError
	}
}
