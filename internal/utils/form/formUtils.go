package formUtils

import (
	"GoBaby/internal/models"
	"net/http"
)

func GetFormValues(r *http.Request, formFields []string) (map[string]string, *models.AppError) {
	err := r.ParseForm()
	if err != nil {
		return nil, &models.AppError{
			Message: "Error parsing form",
			Code:    models.ErrInternalServer,
			Err:     err,
		}
	}
	formValues := make(map[string]string)
	for _, field := range formFields {
		formValues[field] = r.Form.Get(field)
	}
	return formValues, nil
}
