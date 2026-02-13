package repo

import (
	"context"
	"errors"

	"go.mod/models"
	"gorm.io/gorm"
)

type AuthRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) *AuthRepository {
	return &AuthRepository{
		db: db,
	}
}

//
// --------------------- User ---------------------
//

// CreateUser inserts a new user
func (r *AuthRepository) CreateUser(ctx context.Context, user *models.User) error {
	return r.db.WithContext(ctx).Create(user).Error
}

// FindUserByUsername returns user by username
func (r *AuthRepository) FindUserByUsername(ctx context.Context, username string) (*models.User, error) {
	var user models.User

	err := r.db.WithContext(ctx).
		Where("username = ?", username).
		First(&user).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

// IncrementTokenVersion increments token_version column
func (r *AuthRepository) IncrementTokenVersion(ctx context.Context, userID uint64) error {
	return r.db.WithContext(ctx).
		Model(&models.User{}).
		Where("id = ?", userID).
		UpdateColumn("token_version", gorm.Expr("token_version + ?", 1)).
		Error
}

//
// --------------------- RefreshToken ---------------------
//

// SaveRefreshToken inserts new refresh token
func (r *AuthRepository) SaveRefreshToken(ctx context.Context, token *models.RefreshToken) error {
	return r.db.WithContext(ctx).Create(token).Error
}

// GetRefreshToken finds refresh token by JTI
func (r *AuthRepository) GetRefreshToken(ctx context.Context, jti string) (*models.RefreshToken, error) {
	var token models.RefreshToken

	err := r.db.WithContext(ctx).
		Where("jti = ?", jti).
		First(&token).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &token, nil
}

// RevokeRefreshToken sets revoked = true
func (r *AuthRepository) RevokeRefreshToken(ctx context.Context, jti string) error {
	result := r.db.WithContext(ctx).
		Model(&models.RefreshToken{}).
		Where("jti = ?", jti).
		Update("revoked", true)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("refresh token not found")
	}

	return nil
}
