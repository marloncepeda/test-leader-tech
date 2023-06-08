package utils

import "api/pkg/models"

func Response(codeHttp int, data interface{}, message string) models.Response {
	var response models.Response

	response.Code = codeHttp
	response.Message = message
	response.Data = data

	return response
}
