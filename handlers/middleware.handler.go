package handlers

import (
	"net/http"

	"e-commerce/common/messages"
	"e-commerce/models"

	"github.com/gin-gonic/gin"
)

// AuthenticatedUserMiddleware converts a bearer token to an authenticated user
func (h *Handler) AuthenticatedUserMiddleware() gin.HandlerFunc {
	// add the middleware function
	return func(c *gin.Context) {
		user, err := h.controller.Middleware().JwtUserAuth(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, models.ResponseObject{Code: http.StatusBadRequest, Error: err.Error(), Status: "bad-request", Message: err.Error()})
			c.Abort()
		} else {
			c.Set("authUser", user)
		}
		c.Next()
	}
}

// UserPermissionMiddleware ensures a user has user role
func (h *Handler) UserPermissionMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		user := c.MustGet("authUser").(*models.User)
		if user.Role != string(models.USER_ROLE_USER) {
			c.JSON(http.StatusBadRequest, models.ResponseObject{Code: http.StatusUnauthorized, Error: messages.ErrAccessDenied, Status: "unauthorized", Message: messages.ErrAccessDenied.Error()})
			c.Abort()
		}
		c.Next()
	}
}

// AdminPermissionMiddleware ensures a user has an admin role
func (h *Handler) AdminPermissionMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		user := c.MustGet("authUser").(*models.User)
		if user.Role != string(models.USER_ROLE_ADMIN) {
			c.JSON(http.StatusBadRequest, models.ResponseObject{Code: http.StatusUnauthorized, Error: messages.ErrAccessDenied, Status: "unauthorized", Message: messages.ErrAccessDenied.Error()})
			c.Abort()
		}
		c.Next()
	}
}
