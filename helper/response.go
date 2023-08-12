package helper

import "github.com/gin-gonic/gin"

func NewSuccessResponse(c *gin.Context, statusCode int, message string, data interface{}) {
	r := struct {
		StatusCode int         `json:"status_code"`
		Message    string      `json:"message"`
		Data       interface{} `json:"data"`
	}{
		StatusCode: statusCode,
		Message:    message,
		Data:       data,
	}

	c.JSON(statusCode, r)
}

func NewErrorResponse(c *gin.Context, statusCode int, message string, err interface{}) {
	r := struct {
		StatusCode int         `json:"status_code"`
		Message    string      `json:"message"`
		Error      interface{} `json:"errors"`
	}{
		StatusCode: statusCode,
		Message:    message,
		Error:      err,
	}

	c.JSON(statusCode, r)
}
