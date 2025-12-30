package user

import (
	"context"
	"errors"
	errorCommon "user-service/common/error"
	errorConstant "user-service/constants/error"
	"user-service/domain/dto"
	"user-service/domain/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

type IUserRepository interface {
	Register(context.Context, *dto.RegisterRequest) (*models.User, error)
	Update(context.Context, *dto.UpdateRequest, string) (*models.User, error)
	FindByUsername(context.Context, string) (*models.User, error)
	FindByEmail(context.Context, string) (*models.User, error)
	FindByUUID(context.Context, string) (*models.User, error)
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Register(ctx context.Context, req *dto.RegisterRequest) (*models.User, error) {
	user := &models.User{
		UUID:        uuid.New(),
		Name:        req.Name,
		Email:       req.Email,
		Password:    req.Password,
		PhoneNumber: req.PhoneNumber,
		Username:    req.Username,
		RoleID:      req.RoleID,
	}

	err := r.db.WithContext(ctx).Create(user).Error
	if err != nil {
		return nil, errorCommon.WrapError(errorConstant.ErrSQLError)
	}

	return user, nil
}

func (r *UserRepository) Update(ctx context.Context, req *dto.UpdateRequest, uuid string) (*models.User, error) {
	user := &models.User{
		Name:        req.Name,
		Email:       req.Email,
		PhoneNumber: req.PhoneNumber,
		Password:    *req.Password,
		Username:    req.Username,
	}

	err := r.db.WithContext(ctx).
		Where("uuid = ?", uuid).
		Updates(user).Error
	if err != nil {
		return nil, errorCommon.WrapError(errorConstant.ErrSQLError)
	}

	return user, nil
}

func (r *UserRepository) FindByUsername(ctx context.Context, username string) (*models.User, error) {
	var user models.User

	err := r.db.WithContext(ctx).
		Preload("Role").
		Where("username = ?", username).
		First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errorConstant.ErrUserNotFound
		}
		return nil, errorCommon.WrapError(errorConstant.ErrSQLError)
	}

	return &user, nil
}

// FindByEmail implements [IUserRepository].
func (r *UserRepository) FindByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User

	err := r.db.WithContext(ctx).
		Preload("Role").
		Where("email = ?", email).
		First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errorConstant.ErrUserNotFound
		}
		return nil, errorCommon.WrapError(errorConstant.ErrSQLError)
	}

	return &user, nil
}

// FindByUUID implements [IUserRepository].
func (r *UserRepository) FindByUUID(ctx context.Context, uuid string) (*models.User, error) {
	var user models.User

	err := r.db.WithContext(ctx).
		Preload("Role").
		Where("uuid = ?", uuid).
		First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errorConstant.ErrUserNotFound
		}
		return nil, errorCommon.WrapError(errorConstant.ErrSQLError)
	}

	return &user, nil
}
