package controllers

import (
	"context"

	"github.com/google/uuid"

	"e-commerce/common/middleware"
	"e-commerce/config"
	"e-commerce/db"
	"e-commerce/models"
	"e-commerce/repo"
)

// Cobtroller is the controller object
type Controller struct {
	middleware *middleware.Middleware
	Config     *config.ConfigType

	userRepo        repo.UserRepo
	productRepo     repo.ProductRepo
	orderRepo       repo.OrderRepo
	orderRecordRepo repo.OrderRecordRepo
}

// Operations registers all controllers method
type Operations interface {
	Middleware() *middleware.Middleware

	// users
	RegisterUser(ctx context.Context, data *models.SignUpDto) *models.ResponseObject
	Login(ctx context.Context, data *models.SignInDto) *models.ResponseObject

	// order
	PlaceOrder(ctx context.Context, data *models.PlaceOrderDto, user *models.User) *models.ResponseObject
	GetAllOrders(ctx context.Context, user *models.User, query *models.APIPagingDto) *models.ResponseObject
	GetSingleOrder(ctx context.Context, orderId uuid.UUID) *models.ResponseObject
	CancelOrder(ctx context.Context, orderId uuid.UUID, user *models.User) *models.ResponseObject
	UpdateOrderStatus(ctx context.Context, orderId uuid.UUID, data *models.UpdateOrderStatusDto, user *models.User) *models.ResponseObject

	// product
	CreateProduct(ctx context.Context, data *models.CreateProductDto, user *models.User) *models.ResponseObject
	GetSingleProduct(ctx context.Context, productId uuid.UUID) *models.ResponseObject
	UpdateProduct(ctx context.Context, data *models.UpdateProductDto, productId uuid.UUID) *models.ResponseObject
	GetAllProducts(ctx context.Context, query *models.APIPagingDto) *models.ResponseObject
	DeleteProduct(ctx context.Context, productId uuid.UUID) *models.ResponseObject
}

// NewController loads all controllers resources
func NewController(middleware *middleware.Middleware, db *db.Database) *Operations {
	c := &Controller{
		middleware: middleware,

		userRepo:        repo.NewUserRepo(db),
		productRepo:     repo.NewProductRepo(db),
		orderRepo:       repo.NewOrderRepo(db),
		orderRecordRepo: repo.NewOrderRecordRepo(db),
	}
	op := Operations(c)

	return &op
}
func (c *Controller) Middleware() *middleware.Middleware {
	return c.middleware
}
