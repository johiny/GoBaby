package authUtils

import (
	"GoBaby/internal/models"

	"golang.org/x/crypto/bcrypt"
)

func EncryptPassword(password string) (string, *models.AppError) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", &models.AppError{
			Message: "Failed To Hash Password",
			Code:    models.ErrInternalServer,
			Err:     err,
		}
	}
	return string(hash), nil
}

func CheckPassword(password, hash string) *models.AppError {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return &models.AppError{
			Message: "Failed To Check Password",
			Code:    models.ErrInternalServer,
			Err:     err,
		}
	}
	return nil
}
