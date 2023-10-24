package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/iamvibs/golang-projects/jwt/controllers"
	"github.com/iamvibs/golang-projects/jwt/middleware"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/users", controllers.GetUsers())
	incomingRoutes.GET("/users/:user_id", controllers.GetUser())
}
