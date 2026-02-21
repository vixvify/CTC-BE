package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Data       interface{} `json:"data,omitempty"`
	Error      string      `json:"error,omitempty"`
	Status     string      `json:"status"`
	StatusCode int         `json:"statusCode"`
}

func JSON(c *gin.Context, code int, data interface{}, err string) {
	c.JSON(code, Response{
		Data:       data,
		Error:      err,
		Status:     http.StatusText(code),
		StatusCode: code,
	})
}

func OK(c *gin.Context, data interface{}) {
	JSON(c, http.StatusOK, data, "")
}

func Created(c *gin.Context, data interface{}) {
	JSON(c, http.StatusCreated, data, "")
}

func BadRequest(c *gin.Context, err string) {
	JSON(c, http.StatusBadRequest, nil, err)
}

func Unauthorized(c *gin.Context, err string) {
	JSON(c, http.StatusUnauthorized, nil, err)
}

func Internal(c *gin.Context, err string) {
	JSON(c, http.StatusInternalServerError, nil, err)
}

func NotFound(c *gin.Context, err string) {
	JSON(c, http.StatusNotFound, nil, err)
}

func Forbidden(c *gin.Context, err string) {
	JSON(c, http.StatusForbidden, nil, err)
}
