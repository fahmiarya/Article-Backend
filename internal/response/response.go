package response

import "github.com/gin-gonic/gin"

type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func JSON(c *gin.Context, statusCode int, status, message string, data interface{}) {
	if data == nil {
		data = gin.H{}
	}
	c.JSON(statusCode, Response{
		Status:  status,
		Message: message,
		Data:    data,
	})
}

func Success(c *gin.Context, statusCode int, message string, data interface{}) {
	JSON(c, statusCode, "success", message, data)
}

func Error(c *gin.Context, statusCode int, message string, data interface{}) {
	JSON(c, statusCode, "error", message, data)
}
