package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"

	"github.com/google/uuid"
)

type OrderStatus string

const (
	PENDING    OrderStatus = "pending"
	PROCESSING OrderStatus = "processing"
	DELIVERED  OrderStatus = "delivered"
	CANCELLED  OrderStatus = "cancelled"
	SHIPPED    OrderStatus = "shipped"

	ORDER_FEE int64 = 100000 // 1,000 per order
)

// Order is the order object
type Order struct {
	Id           uuid.UUID        `json:"id" gorm:"column:id;PRIMARY_KEY;type:uuid;default:gen_random_uuid()"`
	UserId       uuid.UUID        `json:"user_id"`
	TrackingCode string           `json:"tracking_code"`
	Status       string           `json:"status"`
	Currency     string           `json:"currency"`
	Fee          int64            `json:"fee"`
	History      OrderHistoryData `json:"history"`
	CreatedAt    time.Time        `json:"created_at"`
	UpdatedAt    time.Time        `json:"updated_at"`

	TotalAmount  int64          `json:"total_amount" gorm:"-"`
	OrderRecords []*OrderRecord `json:"order_records" gorm:"foreignkey:OrderId"`
}

// OrderRecord keeps the record for each product ordered
type OrderRecord struct {
	Id        uuid.UUID `json:"id" gorm:"column:id;PRIMARY_KEY;type:uuid;default:gen_random_uuid()"`
	ProductId uuid.UUID `json:"product_id"`
	Quantity  int64     `json:"quantity"`
	OrderId   uuid.UUID `json:"order_id"`
	Amount    int64     `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// OrderHistory is the order history object
type OrderHistory struct {
	Note      string    `json:"note"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}

// OrderHistoryData is the order history data object
type OrderHistoryData struct {
	Data []OrderHistory `json:"data"`
}

// PlaceOrderDto is the place order transfer object
type PlaceOrderDto struct {
	Data     []PlaceOrder `json:"data" validate:"gt=0,dive"`
	Currency Currency     `json:"currency" validate:"required,is_enum"`
}

// PlaceOrder is the place order object
type PlaceOrder struct {
	ProductId string `json:"product_id" validate:"required,is_uuid"`
	Quantity  int64  `json:"quantity" validate:"required,is_amount"`
}

// OrdersResponse is the products data with pagination info
type OrdersResponse struct {
	Orders     []*Order    `json:"orders"`
	PagingInfo *PagingInfo `json:"paging_info"`
}

// UpdateOrderStatusDto is the update data transfer object
type UpdateOrderStatusDto struct {
	Status OrderStatus `json:"status" validate:"required,is_enum"`
}

// GetTotalAmount gets total amount of an order
func (o *Order) GetTotalAmount() int64 {
	var totalAmount int64
	for _, orderRecord := range o.OrderRecords {
		totalAmount += (orderRecord.Amount * orderRecord.Quantity)
	}
	return totalAmount
}

// IsValid checks if status is valid
func (o OrderStatus) IsValid() bool {
	switch o {
	case PENDING, PROCESSING, DELIVERED, CANCELLED, SHIPPED:
		return true
	}
	return false
}

func (t OrderHistoryData) Value() (driver.Value, error) {
	return json.Marshal(t)
}

func (t *OrderHistoryData) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to byte failed")
	}

	return json.Unmarshal(b, &t)
}
