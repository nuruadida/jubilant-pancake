package delivery

import (
	"Nurul-trial/models"
	"Nurul-trial/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
)

type DelivStruct struct {
	UCase usecase.UcaseLoginInterface
}

func NewImplDelivery(r *gin.RouterGroup, UCase usecase.UcaseLoginInterface) {
	handler := DelivStruct{UCase: UCase}
	r.POST("/login", handler.Login)
}

func (a DelivStruct) Login(c *gin.Context) {
	var user models.GetItGuys

	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			Status:  http.StatusBadRequest,
			Message: err.Error() + "Wrong Username, Password, or Email",
		})
		return
	}

	Response := models.Response{
		Status:  http.StatusOK,
		Message: "Ok",
	}
	c.JSON(http.StatusOK, Response)
}
