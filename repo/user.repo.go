package repo

import (
	"context"
	"errors"
	"strings"
	"time"

	"e-commerce/common/messages"
	"e-commerce/db"
	"e-commerce/models"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

// User repo object
type User struct {
	repo *db.Database
}

// TaskRepo exposes user's methods to other packages
type UserRepo interface {
	CreateUser(ctx context.Context, user *models.User) (*models.User, error)
	GetUserByFields(ctx context.Context, fields map[string]interface{}) (*models.User, error)
}

// NewUserRepo instantiates the User Repo object
func NewUserRepo(db *db.Database) UserRepo {
	user := &User{
		repo: db,
	}
	return UserRepo(user)

}

// Create stores a new user
func (u *User) CreateUser(ctx context.Context, user *models.User) (*models.User, error) {
	user.CreatedAt = time.Now().UTC()
	user.UpdatedAt = time.Now().UTC()

	db := u.repo.PostgresDb.WithContext(ctx).Create(user)
	if db.Error != nil {
		log.Err(db.Error).Msgf("Basic::CreateUser error: %v, (%v)", "", db.Error)
		if strings.Contains(db.Error.Error(), "duplicate key value") {
			return nil, errors.New("an error occurred")
		}
		return nil, errors.New("an error occurred")
	}
	return user, nil
}

func (u *User) GetUserByFields(ctx context.Context, fields map[string]interface{}) (*models.User, error) {
	var user models.User
	db := u.repo.PostgresDb.WithContext(ctx).Where(fields).Find(&user)
	if db.Error != nil {
		log.Err(db.Error).Msgf("Basic::GetUserByFields error: %v, (%v)", "record not found", db.Error)
		return &user, errors.New("something went wrong")
	}

	// means no record was found
	if user.Id == uuid.Nil {
		return nil, messages.ErrUserNotFound
	}
	return &user, nil
}
