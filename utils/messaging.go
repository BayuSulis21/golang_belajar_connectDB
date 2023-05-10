package utils

import (
	"service/models"
	"strconv"
)

func ResponseError(code int, msg string) (int, models.ErrorResponse) {
	return code, models.ErrorResponse{
		ResponseCode:    strconv.Itoa(code),
		ResponseMessage: msg,
	}
}
