package models

import (
	"time"

	"github.com/google/uuid"
)

type ProductStatus string
type Currency string

const (
	IN_STOCK     ProductStatus = "in-stock"
	NOT_IN_STOCK ProductStatus = "not-in-stock"
	SOLD_OUT     ProductStatus = "sold-out"

	CURRENCY_NGN Currency = "NGN"
)

// Product is the product model
type Product struct {
	Id                uuid.UUID  `json:"id" gorm:"column:id;PRIMARY_KEY;type:uuid;default:gen_random_uuid()"`
	Slug              string     `json:"slug"`
	Name              string     `json:"name"`
	Description       string     `json:"description"`
	Price             int64      `json:"price"`
	Currency          string     `json:"currency"`
	Discount          int64      `json:"discount"`
	Status            string     `json:"status"`
	AvailableQuantity int64      `json:"available_quantity"`
	CreatedAt         time.Time  `json:"created_at"`
	UpdatedAt         time.Time  `json:"updated_at"`
	DeletedAt         *time.Time `json:"deleted_at"`
}

// CreateProductDto is the data transfer object to create new product
type CreateProductDto struct {
	Name        string   `json:"name" validate:"required,min=4,max=30"`
	Description string   `json:"description" validate:"required,min=4,max=100"`
	Quantity    int64    `json:"quantity" validate:"required,is_amount"`
	Price       int64    `json:"price" validate:"required,is_amount"`
	Discount    int64    `json:"discount" validate:"omitempty,is_amount"`
	Currency    Currency `json:"currency" validate:"required,is_enum"`
}

// UpdateProductDto is the data transfer object to update an existing product
type UpdateProductDto struct {
	Name        *string        `json:"name" validate:"omitempty,min=4,max=30"`
	Description *string        `json:"description" validate:"omitempty,min=4,max=100"`
	Quantity    *int64         `json:"quantity" validate:"omitempty,is_amount"`
	Status      *ProductStatus `json:"status" validate:"omitempty,is_enum"`
	Price       *int64         `json:"price" validate:"omitempty,is_amount"`
	Discount    *int64         `json:"discount" validate:"omitempty,is_amount"`
}

// ProductsResponse is the products data with pagination info
type ProductsResponse struct {
	Products   []*Product  `json:"products"`
	PagingInfo *PagingInfo `json:"paging_info"`
}

// IsValid checks if status is valid
func (p ProductStatus) IsValid() bool {
	switch p {
	case IN_STOCK, NOT_IN_STOCK, SOLD_OUT:
		return true
	}
	return false
}

// IsValid checks if status is valid
func (c Currency) IsValid() bool {
	switch c {
	case CURRENCY_NGN:
		return true
	}
	return false
}
