package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"e-commerce/common/messages"
	"e-commerce/helpers"
	"e-commerce/models"
)

// @Tags Product
// @Summary Create new Product
// @Schemes
// @Description Creates a new product
// @Param   request   body     models.CreateProductDto   true  "data to create new product"
// @Accept json
// @Produce json
// @Success 200 {object} models.ResponseObject "desc"
// @Router /products [post]
func (h *Handler) CreateProduct(c *gin.Context) {
	var input models.CreateProductDto
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
	result := h.controller.CreateProduct(c, &input, user)
	c.JSON(result.Code, result)
}

// @Tags Product
// @Summary Get All Products
// @Description Gets All products
// @Accept  json
// @Produce  json
// @Param   request   body     models.APIPagingDto   true  "data to query for all "
// @Success 200 {string} {object} models.ResponseObject{data=models.ProductsResponse} "desc"
// @Failure 400 {object} map[string]interface{} "Invalid input"
// @Failure 401 {object} map[string]interface{} "Unauthorized"
// @Router /products [get]
func (h *Handler) GetAllProducts(c *gin.Context) {
	query := getPagingInfo(c)
	result := h.controller.GetAllProducts(c, query)
	c.JSON(result.Code, result)
}

// @Tags Product
// @Summary Get Single product
// @Description Get Single product by id
// @Accept  json
// @Produce  json
// @Param   id   path     string   true  "Product Id"
// @Success 200 {string} {object} models.ResponseObject{data=models.Product} "desc"
// @Failure 400 {object} map[string]interface{} "Invalid input"
// @Failure 401 {object} map[string]interface{} "Unauthorized"
// @Router /products/{id} [get]
func (h *Handler) GetSingleProduct(c *gin.Context) {
	id, _ := uuid.Parse(c.Param("id"))
	result := h.controller.GetSingleProduct(c, id)
	c.JSON(result.Code, result)
}

// @Tags Product
// @Summary Update Product
// @Description Updates Product with a given Id
// @Accept  json
// @Produce  json
// @Param   id   path     string   true  "Product Id"
// @Param   request   body     models.UpdateProductDto   true  "data to update product with"
// @Success 200 {string} {object} models.ResponseObject{} "desc"
// @Failure 400 {object} map[string]interface{} "Invalid input"
// @Failure 401 {object} map[string]interface{} "Unauthorized"
// @Router /products/{id} [put]
func (h *Handler) UpdateProduct(c *gin.Context) {
	id, _ := uuid.Parse(c.Param("id"))
	var input models.UpdateProductDto
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
	result := h.controller.UpdateProduct(c, &input, id)
	c.JSON(result.Code, result)
}

// @Tags Product
// @Summary Delete Product
// @Description Delete product by id
// @Accept  json
// @Produce  json
// @Param   id   path     string   true  "Product Id"
// @Success 200 {string} {object} models.ResponseObject{} "desc"
// @Failure 400 {object} map[string]interface{} "Invalid input"
// @Failure 401 {object} map[string]interface{} "Unauthorized"
// @Router /products/{id} [delete]
func (h *Handler) DeleteProduct(c *gin.Context) {
	id, _ := uuid.Parse(c.Param("id"))
	result := h.controller.DeleteProduct(c, id)
	c.JSON(result.Code, result)
}
