package dbUtils

import (
	"GoBaby/internal/models"

	"github.com/google/uuid"
)

func GetRandomUUID() (string, *models.AppError) {
	UUID, err := uuid.NewRandom()
	if err != nil {
		return "", &models.AppError{
			Message: "failed to create UUID",
			Code:    500,
			Err:     err,
		}
	}
	return UUID.String(), nil
}
