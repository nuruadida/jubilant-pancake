package router

import (
	"Nurul-trial/config/db"
	"Nurul-trial/delivery"
	"Nurul-trial/repository"
	"Nurul-trial/usecase"
	"fmt"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	Router := gin.New()

	DB := db.DB
	Router.Use(gin.Logger())
	Router.Use(gin.Recovery())

	fmt.Println(DB, "<<<<")

	// TODO : CALL DEPENDENCY REPOSITORY
	LoginRepo := repository.NewImplRepo(DB)
	UserRepo := repository.NewImplUserRepo(DB)

	// TODO : CALL DEPENDENCY USECASE
	LoginUcase := usecase.NewImplUcase(LoginRepo)
	UserUcase := usecase.NewImplUserUcase(UserRepo)

	NewRouteRole := Router.Group("")

	// TODO : CALL DEPENDENCY DELIVERY
	delivery.NewImplDelivery(NewRouteRole, LoginUcase)
	delivery.NewImplUserDel(NewRouteRole, UserUcase)

	return Router
}
