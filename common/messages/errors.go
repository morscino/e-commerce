package messages

import "errors"

var (
	ErrNoDataFound                   = errors.New("no data found")
	ErrUserNotFound                  = errors.New("user not found")
	ErrOrderNotFound                 = errors.New("order not found")
	ErrOrderCannotBeCancelled        = errors.New("order cannot be cancelled")
	ErrCancelledOrderCannotBeUpdated = errors.New("cancelled order cannot be updated")
	ErrProductNotFound               = errors.New("product not found")
	ErrUserWithEmailAlreadyExists    = errors.New("user with email already exists")
	ErrProductWithNameAlreadyExists  = errors.New("product with name already exists")
	ErrInvalidToken                  = errors.New("invalid token")
	ErrAccessDenied                  = errors.New("access denied")
	ErrWrongPassword                 = errors.New("wrong password")
	ErrCouldNotGenerateToken         = errors.New("could not generate user token")
	ErrTaskWithSlugAlreadyExists     = errors.New("task with slug already exists")
	ErrInvalidInput                  = errors.New("invalid input")
	ErrServerError                   = errors.New("server error")
)
