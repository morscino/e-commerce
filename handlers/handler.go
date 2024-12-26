package handlers

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"

	"e-commerce/common/middleware"
	"e-commerce/config"
	"e-commerce/controllers"
	"e-commerce/db"
	"e-commerce/models"
)

type Handler struct {
	controller controllers.Operations
}

type Operations interface {
	// middleware
	AuthenticatedUserMiddleware() gin.HandlerFunc
	UserPermissionMiddleware() gin.HandlerFunc
	AdminPermissionMiddleware() gin.HandlerFunc

	// product
	CreateProduct(c *gin.Context)
	GetAllProducts(c *gin.Context)
	GetSingleProduct(c *gin.Context)
	UpdateProduct(c *gin.Context)
	DeleteProduct(c *gin.Context)
	// order
	PlaceOrder(c *gin.Context)
	GetAllOrders(c *gin.Context)
	UpdateOrderStatus(c *gin.Context)
	CancelOrder(c *gin.Context)
	GetSingleOrder(c *gin.Context)

	// users
	Login(c *gin.Context)
	SignUp(c *gin.Context)
}

func NewHandler(config *config.ConfigType, db *db.Database) Operations {
	middleware, err := middleware.NewMiddleware(db, config)
	if err != nil {
		log.Logger.Fatal().Msg(fmt.Sprintf("Create middleware error : %s", err.Error()))
	}
	h := &Handler{
		controller: *controllers.NewController(middleware, db),
	}
	return Operations(h)
}

func getPagingInfo(c *gin.Context) *models.APIPagingDto {
	var paging models.APIPagingDto

	limit, _ := strconv.Atoi(c.Query("limit"))
	page, _ := strconv.Atoi(c.Query("page"))
	paging.Filter = c.Query("filter")

	// default limit is 10
	if limit < 1 {
		limit = 10
	}
	if page < 1 {
		page = 1
	}
	paging.Limit = limit
	paging.Page = page
	return &paging
}
