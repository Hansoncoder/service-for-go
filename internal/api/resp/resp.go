package resp

import (
	"errors"
	"net/http"
	CustomError "veo/pkg/errors"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func JSONResponse(c *gin.Context, data interface{}, err error) {
	if err != nil {
		var be *CustomError.CustomError
		if errors.As(err, &be) {
			c.JSON(http.StatusOK, Response{
				Code:    int(be.Code),
				Message: be.Message,
				Data:    data,
			})
		} else {
			c.JSON(http.StatusInternalServerError, Response{
				Code:    http.StatusInternalServerError,
				Message: "service fail",
			})
		}
		return
	}

	c.JSON(http.StatusOK, Response{
		Code: http.StatusOK,
		Data: data,
	})

}
