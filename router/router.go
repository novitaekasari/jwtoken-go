package router

import (
	"github.com/gin-gonic/gin"
	"github.com/novitaekasari/JWToken/controllers"
	"github.com/novitaekasari/JWToken/middlewares"
)

// func StartApp() *gin.Engine {
// 	r := gin.Default()

// 	userRouter := r.Group("/users")
// 	{
// 		userRouter.POST("/register", controllers.UserRegister)

// 		userRouter.POST("/login", controllers.UserLogin)
// 	}

// 	productRouter := r.Group("/products")
// 	{
// 		productRouter.Use(middlewares.Authentication())

// 		productRouter.POST("/", controllers.CreateProduct)
// 	}
// 	return r
// }

func StartApp() *gin.Engine {
	r := gin.Default()

	userRouter := r.Group("/users")
	{
		userRouter.POST("/register", controllers.UserRegister)
		userRouter.POST("/login", controllers.UserLogin)
	}

	productRouter := r.Group("/products")
	{
		productRouter.Use(middlewares.Authentication())
		productRouter.POST("/", controllers.CreateProduct)
		productRouter.PUT("/:productId", middlewares.ProductAuthorization(), controllers.UpdateProduct)
	}

return r
}