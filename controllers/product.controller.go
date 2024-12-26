package controllers

import (
	"context"
	"net/http"

	"e-commerce/common/messages"
	"e-commerce/helpers"
	"e-commerce/models"

	"github.com/google/uuid"
)

// CreateProduct creates a new product
func (c *Controller) CreateProduct(ctx context.Context, data *models.CreateProductDto, user *models.User) *models.ResponseObject {
	slug := helpers.ToSlug(data.Name)
	existingProduct, err := c.productRepo.GetProductByFields(ctx, helpers.Map{"slug": slug})
	if err != nil && err != messages.ErrProductNotFound {
		return handleError(err, "server-error", http.StatusInternalServerError)
	}
	if existingProduct != nil {
		return handleError(messages.ErrProductWithNameAlreadyExists, "bad-request", http.StatusBadRequest)
	}

	newProduct := &models.Product{
		Name:              data.Name,
		Description:       data.Description,
		Slug:              slug,
		Price:             data.Price,
		Currency:          string(data.Currency),
		AvailableQuantity: data.Quantity,
		Discount:          data.Discount,
		Status:            string(models.IN_STOCK),
	}
	product, err := c.productRepo.CreateProduct(ctx, newProduct)
	if err != nil {
		return handleError(err, "server-error", http.StatusInternalServerError)
	}

	return handleSuccess(product, "success", "product created successfully", http.StatusCreated)
}

func (c *Controller) GetSingleProduct(ctx context.Context, productId uuid.UUID) *models.ResponseObject {
	product, err := c.productRepo.GetProductByFields(ctx, helpers.Map{"id": productId})
	if err != nil {
		return handleError(err, "server-error", http.StatusInternalServerError)
	}
	return handleSuccess(product, "success", "product fetched successfully", http.StatusOK)
}

func (c *Controller) UpdateProduct(ctx context.Context, data *models.UpdateProductDto, productId uuid.UUID) *models.ResponseObject {
	var update models.Product
	if data.Name != nil {
		slug := helpers.ToSlug(*data.Name)
		_, err := c.productRepo.GetProductByFields(ctx, helpers.Map{"slug": slug})
		if err != nil && err != messages.ErrProductNotFound {
			return handleError(err, "server-error", http.StatusInternalServerError)
		}

		update.Name = *data.Name
		update.Slug = slug
	}

	if data.Description != nil {
		update.Description = *data.Description
	}

	if data.Quantity != nil {
		update.AvailableQuantity = *data.Quantity
	}

	if data.Price != nil {
		update.Price = *data.Price
	}

	if data.Discount != nil {
		update.Discount = *data.Discount
	}

	if data.Status != nil {
		update.Status = string(*data.Status)
	}

	err := c.productRepo.UpdateProductById(ctx, productId, &update)
	if err != nil {
		return handleError(err, "server-error", http.StatusInternalServerError)
	}

	return handleSuccess(nil, "success", "product updated successfully", http.StatusOK)
}

func (c *Controller) GetAllProducts(ctx context.Context, query *models.APIPagingDto) *models.ResponseObject {
	result, err := c.productRepo.GetAllProducts(ctx, query)
	if err != nil {
		return handleError(err, "server-error", http.StatusInternalServerError)
	}
	return handleSuccess(result, "success", "products fetched successfully", http.StatusOK)
}

func (c *Controller) DeleteProduct(ctx context.Context, productId uuid.UUID) *models.ResponseObject {
	err := c.productRepo.DeleteProduct(ctx, &models.Product{Id: productId})
	if err != nil {
		return handleError(err, "server-error", http.StatusInternalServerError)
	}
	return handleSuccess(nil, "success", "product deleted successfully", http.StatusOK)
}
