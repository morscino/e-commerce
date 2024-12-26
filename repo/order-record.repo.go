package repo

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/rs/zerolog/log"

	"e-commerce/db"
	"e-commerce/models"
)

// OrderRecord repo object
type OrderRecord struct {
	repo *db.Database
}

// OrderRepo exposes order's methods to other packages
type OrderRecordRepo interface {
	CreateOrderRecord(ctx context.Context, order *models.OrderRecord) (*models.OrderRecord, error)
}

// NewOrderRecordRepo instantiates the Order Repo object
func NewOrderRecordRepo(db *db.Database) OrderRecordRepo {
	orderRecord := &OrderRecord{
		repo: db,
	}
	orderRecordRepo := OrderRecordRepo(orderRecord)
	return orderRecordRepo
}

// CreateOrderRecord stores a new order
func (o *OrderRecord) CreateOrderRecord(ctx context.Context, orderRecord *models.OrderRecord) (*models.OrderRecord, error) {
	orderRecord.CreatedAt = time.Now().UTC()
	orderRecord.UpdatedAt = time.Now().UTC()

	db := o.repo.PostgresDb.WithContext(ctx).Create(orderRecord)
	if db.Error != nil {
		log.Err(db.Error).Msgf("Basic::CreateOrderRecord error: %v, (%v)", "", db.Error)
		if strings.Contains(db.Error.Error(), "duplicate key value") {
			return nil, errors.New("an error occurred")
		}
		return nil, errors.New("an error occurred")
	}
	return orderRecord, nil
}
