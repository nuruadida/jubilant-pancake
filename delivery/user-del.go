package delivery

import (
	"Nurul-trial/models"
	"Nurul-trial/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserDelivery struct {
	UserUcase usecase.UcaseUserInterface
}

func NewImplUserDel(r *gin.RouterGroup, Ucase usecase.UcaseUserInterface) {
	handler := &UserDelivery{UserUcase: Ucase}

	r.POST("/register", handler.AddUser)
}

func (u UserDelivery) AddUser(c *gin.Context) {
	user := models.User{}

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			Status:  http.StatusBadRequest,
			Message: "SALAH GAES!",
		})
		return
	}

	user, err := u.UserUcase.AddUser(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			Status:  http.StatusBadRequest,
			Message: "error",
		})
		return
	}

	Response := models.Response{
		Status:  http.StatusAccepted,
		Message: "SUKSES GAES!",
	}
	c.JSON(http.StatusAccepted, Response)

}
