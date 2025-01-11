package routes

import (
	"github.com/Goland-backend-loja-dejogos/src/controller"
	"github.com/gin-gonic/gin"
)

// manda as rotas para o reteador
func InitRoutes(r *gin.RouterGroup){
	// no gin /: é quando ele vai receber um parametro atraves da url
	//da para criar controller aqui r.GET("/...., func(c *gin.Context) {}"), pode se ser até mais de um mas fica muito caótico
	r.GET("/getUserById/:userId", controller.FindUserById)
	r.GET("/getUserByEmail/:userEmail", controller.FindUserByEmail)
	r.POST("/createUser", controller.CreateUser)
	r.PUT("/updateUser/:userId", controller.UpdateUser)
	r.DELETE("/deleteUser/:userId", controller.DeleteUser)
}