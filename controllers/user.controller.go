package controllers

import (
	"context"
	"net/http"
	"strings"

	"github.com/google/uuid"

	"e-commerce/common/messages"
	"e-commerce/helpers"
	"e-commerce/models"
)

// RegisterUser signs up users
func (c *Controller) RegisterUser(ctx context.Context, data *models.SignUpDto) *models.ResponseObject {
	// check if user with email exists
	existingUser, err := c.userRepo.GetUserByFields(ctx, helpers.Map{"email": data.Email})
	if err != nil && err != messages.ErrUserNotFound {
		return handleError(err, "server-error", http.StatusInternalServerError)
	}
	if existingUser != nil {
		return handleError(messages.ErrUserWithEmailAlreadyExists, "bad-request", http.StatusBadRequest)
	}
	role := string(models.USER_ROLE_USER)
	if data.Role != nil {
		role = string(*data.Role)
	}
	newUser := &models.User{
		Id:           uuid.New(),
		FirstName:    data.FirstName,
		LastName:     data.LastName,
		Role:         role,
		Email:        strings.ToLower(data.Email),
		PasswordHash: helpers.Hash(data.Password),
	}

	user, err := c.userRepo.CreateUser(ctx, newUser)
	if err != nil {
		return handleError(err, "server-error", http.StatusInternalServerError)
	}
	return &models.ResponseObject{Code: http.StatusCreated, Data: user, Status: "success", Message: "user signed up successfully"}
}

func handleError(err error, status string, code int) *models.ResponseObject {
	return &models.ResponseObject{
		Code:    code,
		Status:  status,
		Error:   err,
		Message: err.Error(),
	}
}

func handleSuccess(data interface{}, status, message string, code int) *models.ResponseObject {
	return &models.ResponseObject{
		Code:    code,
		Status:  status,
		Data:    data,
		Message: message,
	}
}

// Login logs user in
func (c *Controller) Login(ctx context.Context, data *models.SignInDto) *models.ResponseObject {
	// get user
	user, err := c.userRepo.GetUserByFields(ctx, helpers.Map{"email": data.Email})
	if err != nil {
		return &models.ResponseObject{
			Code:    http.StatusOK,
			Status:  "no-data-found",
			Message: err.Error(),
		}
	}

	// verify password
	if isValid := helpers.CompareHash(user.PasswordHash, data.Password); !isValid {
		return &models.ResponseObject{
			Code:    http.StatusBadRequest,
			Status:  "no-data-found",
			Message: messages.ErrWrongPassword.Error(),
		}
	}

	// generate jwt tokens
	token, err := c.middleware.Jwt.CreateAuthToken(user)
	if err != nil {
		return &models.ResponseObject{
			Code:    http.StatusInternalServerError,
			Status:  "server-errors",
			Message: err.Error(),
			Error:   err,
		}
	}
	authUser := &models.AuthenticatedUser{
		User:        user,
		AccessToken: token,
	}

	return &models.ResponseObject{Code: http.StatusOK, Data: authUser, Status: "success", Message: "user logged in successfully"}

}
