package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"e-commerce/common/messages"
	"e-commerce/helpers"
	"e-commerce/models"
)

// @Tags Order
// @Summary Place Order
// @Schemes
// @Description Places a new Order
// @Param   request   body     models.PlaceOrderDto   true  "data to place new order"
// @Accept json
// @Produce json
// @Success 200 {object} models.ResponseObject "desc"
// @Router /orders [post]
func (h *Handler) PlaceOrder(c *gin.Context) {
	var input models.PlaceOrderDto
	// bind input
	err := c.BindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseObject{Code: http.StatusBadRequest, Error: err, Status: "bad-request", Message: messages.ErrInvalidInput.Error()})
		return
	}
	inputErrors := helpers.ValidateInput(input)
	if inputErrors != nil {

		c.JSON(http.StatusBadRequest, models.ResponseObject{Code: http.StatusBadRequest, Error: inputErrors, Status: "bad-request", Message: messages.ErrInvalidInput.Error()})
		return
	}
	user := c.MustGet("authUser").(*models.User) // auth user
	// send to controller
	result := h.controller.PlaceOrder(c, &input, user)
	c.JSON(result.Code, result)
}

// @Tags Order
// @Summary Get All Orders
// @Description Gets All Orders
// @Accept  json
// @Produce  json
// @Param   request   body     models.APIPagingDto   true  "data to query for all "
// @Success 200 {string} {object} models.ResponseObject{data=models.OrdersResponse} "desc"
// @Failure 400 {object} map[string]interface{} "Invalid input"
// @Failure 401 {object} map[string]interface{} "Unauthorized"
// @Router /orders [get]
func (h *Handler) GetAllOrders(c *gin.Context) {
	query := getPagingInfo(c)
	user := c.MustGet("authUser").(*models.User) // auth user
	result := h.controller.GetAllOrders(c, user, query)
	c.JSON(result.Code, result)
}

// @Tags Order
// @Summary Get Single Order
// @Description Get Single Order by id
// @Accept  json
// @Produce  json
// @Param   id   path     string   true  "Order Id"
// @Success 200 {string} {object} models.ResponseObject{data=models.Order} "desc"
// @Failure 400 {object} map[string]interface{} "Invalid input"
// @Failure 401 {object} map[string]interface{} "Unauthorized"
// @Router /orders/{id} [get]
func (h *Handler) GetSingleOrder(c *gin.Context) {
	id, _ := uuid.Parse(c.Param("id"))
	result := h.controller.GetSingleOrder(c, id)
	c.JSON(result.Code, result)
}

// @Tags Order
// @Summary Update Order Status
// @Description Update Order Status with a given Id
// @Accept  json
// @Produce  json
// @Param   id   path     string   true  "Order Id"
// @Param   request   body     models.UpdateOrderStatusDto   true  "data to update order status"
// @Success 200 {string} {object} models.ResponseObject{} "desc"
// @Failure 400 {object} map[string]interface{} "Invalid input"
// @Failure 401 {object} map[string]interface{} "Unauthorized"
// @Router /orders/{id}/status [put]
func (h *Handler) UpdateOrderStatus(c *gin.Context) {
	id, _ := uuid.Parse(c.Param("id"))
	var input models.UpdateOrderStatusDto
	// bind input
	err := c.BindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseObject{Code: http.StatusBadRequest, Error: err, Status: "bad-request", Message: messages.ErrInvalidInput.Error()})
		return
	}
	inputErrors := helpers.ValidateInput(input)
	if inputErrors != nil {
		c.JSON(http.StatusBadRequest, models.ResponseObject{Code: http.StatusBadRequest, Error: inputErrors, Status: "bad-request", Message: messages.ErrInvalidInput.Error()})
		return
	}
	user := c.MustGet("authUser").(*models.User) // auth user
	result := h.controller.UpdateOrderStatus(c, id, &input, user)
	c.JSON(result.Code, result)
}

// @Tags Order
// @Summary Cancel Order
// @Description Cancel Order with a given Id
// @Accept  json
// @Produce  json
// @Param   id   path     string   true  "Order Id"
// @Success 200 {string} {object} models.ResponseObject{} "desc"
// @Failure 400 {object} map[string]interface{} "Invalid input"
// @Failure 401 {object} map[string]interface{} "Unauthorized"
// @Router /orders/{id}/cancel [put]
func (h *Handler) CancelOrder(c *gin.Context) {
	id, _ := uuid.Parse(c.Param("id"))
	user := c.MustGet("authUser").(*models.User) // auth user
	result := h.controller.CancelOrder(c, id, user)
	c.JSON(result.Code, result)
}
