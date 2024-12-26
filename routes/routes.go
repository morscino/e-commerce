package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"e-commerce/handlers"
)

type Routes struct {
	handler handlers.Operations
}

func NewRoutes(h handlers.Operations) Routes {
	return Routes{handler: h}
}

func (ro Routes) RegisterRoutes(r *gin.Engine, handler handlers.Operations) {
	CheckRoutes(r)

	// products
	products := r.Group("products", handler.AuthenticatedUserMiddleware())
	{
		products.POST("", handler.AdminPermissionMiddleware(), handler.CreateProduct)
		products.GET("", handler.GetAllProducts)
		products.GET("/:id", handler.GetSingleProduct)
		products.PUT("/:id", handler.AdminPermissionMiddleware(), handler.UpdateProduct)
		products.DELETE("/:id", handler.AdminPermissionMiddleware(), handler.DeleteProduct)
	}
	// orders
	orders := r.Group("orders", handler.AuthenticatedUserMiddleware())
	{
		orders.POST("", handler.UserPermissionMiddleware(), handler.PlaceOrder)
		orders.GET("", handler.UserPermissionMiddleware(), handler.GetAllOrders)
		orders.GET("/:id", handler.UserPermissionMiddleware(), handler.GetSingleOrder)
		orders.PUT("/:id/status", handler.AdminPermissionMiddleware(), handler.UpdateOrderStatus)
		orders.PUT("/:id/cancel", handler.UserPermissionMiddleware(), handler.CancelOrder)
	}

	// auth
	auth := r.Group("auth")
	{
		auth.POST("", handler.SignUp)
		auth.POST("/login", handler.Login)
	}

}

func CheckRoutes(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to Task App API",
		})
	})
}
