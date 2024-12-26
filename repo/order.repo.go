package repo

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"

	"e-commerce/common/messages"
	"e-commerce/db"
	"e-commerce/models"
)

// Order repo object
type Order struct {
	repo *db.Database
}

// OrderRepo exposes order's methods to other packages
type OrderRepo interface {
	CreateOrder(ctx context.Context, order *models.Order) (*models.Order, error)
	GetOrderByFields(ctx context.Context, fields map[string]interface{}) (*models.Order, error)
	UpdateOrderById(ctx context.Context, id uuid.UUID, order *models.Order) error
	GetAllOrders(ctx context.Context, query *models.APIPagingDto, fields map[string]interface{}) (*models.OrdersResponse, error)
}

// NewOrderRepo instantiates the Order Repo object
func NewOrderRepo(db *db.Database) OrderRepo {
	order := &Order{
		repo: db,
	}
	return OrderRepo(order)
}

// CreateOrder stores a new order
func (o *Order) CreateOrder(ctx context.Context, order *models.Order) (*models.Order, error) {
	order.CreatedAt = time.Now().UTC()
	order.UpdatedAt = time.Now().UTC()

	db := o.repo.PostgresDb.WithContext(ctx).Create(order)
	if db.Error != nil {
		log.Err(db.Error).Msgf("Basic::CreateOrder error: %v, (%v)", "", db.Error)
		if strings.Contains(db.Error.Error(), "duplicate key value") {
			return nil, errors.New("an error occurred")
		}
		return nil, errors.New("an error occurred")
	}
	return order, nil
}

func (o *Order) GetOrderByFields(ctx context.Context, fields map[string]interface{}) (*models.Order, error) {
	var order models.Order
	db := o.repo.PostgresDb.WithContext(ctx).Where(fields).Preload("OrderRecords").Find(&order)
	if db.Error != nil {
		log.Err(db.Error).Msgf("Basic::GetOrderByFields error: %v, (%v)", "record not found", db.Error)
		return &order, errors.New("something went wrong")
	}

	// means no record was found
	if order.Id == uuid.Nil {
		return nil, messages.ErrOrderNotFound
	}
	return &order, nil
}

func (o *Order) UpdateOrderById(ctx context.Context, id uuid.UUID, order *models.Order) error {
	order.UpdatedAt = time.Now().UTC()
	db := o.repo.PostgresDb.WithContext(ctx).Model(&models.Order{
		Id: id,
	}).UpdateColumns(order)
	if db.Error != nil {
		log.Err(db.Error).Msgf("Basic::UpdateOrderById error: %v, (%v)", "update not successful", db.Error)
		return errors.New("update not successful")
	}

	return nil
}

func (o *Order) GetAllOrders(ctx context.Context, query *models.APIPagingDto, fields map[string]interface{}) (*models.OrdersResponse, error) {
	var orders []*models.Order
	var count, queryCount int64
	queryInfo, offset := getPaginationInfo(query)

	db := o.repo.PostgresDb.WithContext(ctx).Model(&models.Order{}).Preload("OrderRecords").Where(fields)
	filters := getFilterFromQuery(query.Filter)
	for _, filter := range filters {
		db = db.Where(fmt.Sprintf("%s %s ?", filter.field, filter.condition), filter.value)
	}
	// then do counting of all
	db.Count(&count)

	db = db.Offset(offset).Limit(queryInfo.Limit).
		Order(fmt.Sprintf("orders.%s %s", queryInfo.Sort, queryInfo.Direction)).
		Find(&orders)
	db.Count(&queryCount)

	if db.Error != nil {
		log.Err(db.Error).Msgf("Basic::GetAll error: %v, (%v)", "record not found", db.Error)
		return nil, errors.New("record not found")
	}

	pagingInfo := getPagingInfo(queryInfo, int(count))
	pagingInfo.Count = len(orders)
	return &models.OrdersResponse{
		Orders:     orders,
		PagingInfo: &pagingInfo,
	}, nil

}
