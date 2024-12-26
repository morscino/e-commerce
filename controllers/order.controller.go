package controllers

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"

	"e-commerce/common/messages"
	"e-commerce/helpers"
	"e-commerce/models"
)

// place order
func (c *Controller) PlaceOrder(ctx context.Context, data *models.PlaceOrderDto, user *models.User) *models.ResponseObject {
	order := models.Order{
		Id:           uuid.New(),
		UserId:       user.Id,
		TrackingCode: helpers.GenerateUniqueReferenceId(12),
		Status:       string(models.PENDING),
		Fee:          models.ORDER_FEE,
		Currency:     string(data.Currency),
		History: models.OrderHistoryData{
			Data: []models.OrderHistory{
				{
					Note:      "order placed",
					CreatedAt: time.Now(),
					Status:    string(models.PENDING),
				},
			},
		},
	}

	var orderRecords []*models.OrderRecord
	for _, orderData := range data.Data {
		id, _ := uuid.Parse(orderData.ProductId)
		product, err := c.productRepo.GetProductByFields(ctx, helpers.Map{"id": id})
		if err != nil && err != messages.ErrProductNotFound {
			return handleError(err, "server-error", http.StatusInternalServerError)
		}

		if product == nil {
			return handleError(messages.ErrProductNotFound, "bad-request", http.StatusBadRequest)
		}

		amount := product.Price - product.Discount
		orderRecord := models.OrderRecord{
			ProductId: product.Id,
			Quantity:  orderData.Quantity,
			Amount:    amount,
			OrderId:   order.Id,
		}
		// create order record
		newOrderRecord, err := c.orderRecordRepo.CreateOrderRecord(ctx, &orderRecord)
		if err != nil {
			return handleError(err, "server-error", http.StatusInternalServerError)
		}
		orderRecords = append(orderRecords, newOrderRecord)
	}

	// create order
	newOrder, err := c.orderRepo.CreateOrder(ctx, &order)
	if err != nil {
		return handleError(err, "server-error", http.StatusInternalServerError)
	}

	newOrder.OrderRecords = orderRecords
	newOrder.TotalAmount = newOrder.GetTotalAmount()

	return handleSuccess(newOrder, "success", "order created successfully", http.StatusCreated)
}

// get all orders
func (c *Controller) GetAllOrders(ctx context.Context, user *models.User, query *models.APIPagingDto) *models.ResponseObject {
	response, err := c.orderRepo.GetAllOrders(ctx, query, helpers.Map{"user_id": user.Id})
	if err != nil {
		return handleError(err, "server-error", http.StatusInternalServerError)
	}
	return handleSuccess(response, "success", "orders successfully fetched", http.StatusOK)
}

// get single order
func (c *Controller) GetSingleOrder(ctx context.Context, orderId uuid.UUID) *models.ResponseObject {
	order, err := c.orderRepo.GetOrderByFields(ctx, helpers.Map{"id": orderId})
	if err != nil {
		return handleError(err, "server-error", http.StatusInternalServerError)
	}
	order.TotalAmount = order.GetTotalAmount()
	return handleSuccess(order, "success", "order successfully fetched", http.StatusOK)
}

// cancel order
func (c *Controller) CancelOrder(ctx context.Context, orderId uuid.UUID, user *models.User) *models.ResponseObject {
	order, err := c.orderRepo.GetOrderByFields(ctx, helpers.Map{"id": orderId, "user_id": user.Id})
	if err != nil {
		return handleError(err, "server-error", http.StatusInternalServerError)
	}
	if order.Status != string(models.PENDING) {
		return handleError(messages.ErrOrderCannotBeCancelled, "bad-request", http.StatusBadRequest)
	}
	history := models.OrderHistory{
		Note:      "order cancelled",
		Status:    string(models.CANCELLED),
		CreatedAt: time.Now().UTC(),
	}
	order.History.Data = append(order.History.Data, history)
	err = c.orderRepo.UpdateOrderById(ctx, orderId, &models.Order{
		Status:  string(models.CANCELLED),
		History: order.History,
	})
	if err != nil {
		return handleError(err, "server-error", http.StatusInternalServerError)
	}

	return handleSuccess(nil, "success", "order successfully cancelled", http.StatusOK)
}

// update order status
func (c *Controller) UpdateOrderStatus(ctx context.Context, orderId uuid.UUID, data *models.UpdateOrderStatusDto, user *models.User) *models.ResponseObject {
	order, err := c.orderRepo.GetOrderByFields(ctx, helpers.Map{"id": orderId})
	if err != nil {
		return handleError(err, "server-error", http.StatusInternalServerError)
	}

	if order.Status == string(models.CANCELLED) {
		return handleError(messages.ErrOrderCannotBeCancelled, "bad-request", http.StatusBadRequest)
	}
	history := models.OrderHistory{
		Note:      fmt.Sprintf("order %s", string(data.Status)),
		Status:    string(data.Status),
		CreatedAt: time.Now().UTC(),
	}
	order.History.Data = append(order.History.Data, history)
	err = c.orderRepo.UpdateOrderById(ctx, orderId, &models.Order{
		Status:  string(data.Status),
		History: order.History,
	})
	if err != nil {
		return handleError(err, "server-error", http.StatusInternalServerError)
	}
	return handleSuccess(nil, "success", "order successfully updated", http.StatusOK)
}
