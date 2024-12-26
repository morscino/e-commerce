package repo

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"e-commerce/common/messages"
	"e-commerce/db"
	"e-commerce/models"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

// Product repo object
type Product struct {
	repo *db.Database
}

// ProductRepo exposes product's methods to other packages
type ProductRepo interface {
	CreateProduct(ctx context.Context, product *models.Product) (*models.Product, error)
	GetProductByFields(ctx context.Context, fields map[string]interface{}) (*models.Product, error)
	GetAllProducts(ctx context.Context, query *models.APIPagingDto) (*models.ProductsResponse, error)
	UpdateProductById(ctx context.Context, id uuid.UUID, product *models.Product) error
	DeleteProduct(ctx context.Context, product *models.Product) error
}

// NewProductsRepo instantiates the User Repo object
func NewProductRepo(db *db.Database) ProductRepo {
	product := &Product{
		repo: db,
	}
	return ProductRepo(product)

}

// Create stores a new product
func (p *Product) CreateProduct(ctx context.Context, product *models.Product) (*models.Product, error) {
	product.CreatedAt = time.Now().UTC()
	product.UpdatedAt = time.Now().UTC()

	db := p.repo.PostgresDb.WithContext(ctx).Create(product)
	if db.Error != nil {
		log.Err(db.Error).Msgf("Basic::CreateProduct error: %v, (%v)", "", db.Error)
		if strings.Contains(db.Error.Error(), "duplicate key value") {
			return nil, errors.New("an error occurred")
		}
		return nil, errors.New("an error occurred")
	}
	return product, nil
}

func (p *Product) GetProductByFields(ctx context.Context, fields map[string]interface{}) (*models.Product, error) {
	var product models.Product
	db := p.repo.PostgresDb.WithContext(ctx).Where(fields).Find(&product)
	if db.Error != nil {
		log.Err(db.Error).Msgf("Basic::GetProductByFields error: %v, (%v)", "record not found", db.Error)
		return &product, errors.New("something went wrong")
	}

	// means no record was found
	if product.Id == uuid.Nil {
		return nil, messages.ErrProductNotFound
	}
	return &product, nil
}

func (p *Product) UpdateProductById(ctx context.Context, id uuid.UUID, product *models.Product) error {
	product.UpdatedAt = time.Now().UTC()
	db := p.repo.PostgresDb.WithContext(ctx).Model(&models.Product{
		Id: id,
	}).UpdateColumns(product)
	if db.Error != nil {
		log.Err(db.Error).Msgf("Basic::UpdateProductByID error: %v, (%v)", "update not successful", db.Error)
		return errors.New("update not successful")
	}

	return nil
}

func (p *Product) DeleteProduct(ctx context.Context, product *models.Product) error {
	db := p.repo.PostgresDb.WithContext(ctx).Model(&models.Product{}).Delete(product)
	if db.Error != nil {
		log.Err(db.Error).Msgf("Basic::UpdateSavingsGoalByID error: %v, (%v)", "delete not successful", db.Error)
		return errors.New("delete not successful")
	}
	return nil
}

func (p *Product) GetAllProducts(ctx context.Context, query *models.APIPagingDto) (*models.ProductsResponse, error) {
	var products []*models.Product
	var count, queryCount int64
	queryInfo, offset := getPaginationInfo(query)

	db := p.repo.PostgresDb.WithContext(ctx).Model(&models.Product{})
	filters := getFilterFromQuery(query.Filter)
	for _, filter := range filters {
		db = db.Where(fmt.Sprintf("%s %s ?", filter.field, filter.condition), filter.value)
	}
	// then do counting of all
	db.Count(&count)

	db = db.Offset(offset).Limit(queryInfo.Limit).
		Order(fmt.Sprintf("products.%s %s", queryInfo.Sort, queryInfo.Direction)).
		Find(&products)
	db.Count(&queryCount)

	if db.Error != nil {
		log.Err(db.Error).Msgf("Basic::GetAll error: %v, (%v)", "record not found", db.Error)
		return nil, errors.New("record not found")
	}

	pagingInfo := getPagingInfo(queryInfo, int(count))
	pagingInfo.Count = len(products)
	return &models.ProductsResponse{
		Products:   products,
		PagingInfo: &pagingInfo,
	}, nil

}
