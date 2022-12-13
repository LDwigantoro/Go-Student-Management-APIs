package utils

import (
	"net/http"
	"strconv"

	entities "github.com/LDwigantoro/go-clean-arch/entities/response"
	"github.com/gin-gonic/gin"
)

func HandleSuccess(c *gin.Context, data interface{}) {
	responeData := entities.Response{
		Status:  "200",
		Message: "Perintah berhasil dilakukan",
		Data:    data,
	}

	c.JSON(http.StatusOK, responeData)

}

func HandleError(c *gin.Context, status int, message string) {
	responseData := entities.Response{
		Status:  strconv.Itoa(status),
		Message: message,
	}
	c.JSON(status, responseData)
}
