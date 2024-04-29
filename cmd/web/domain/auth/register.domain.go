package authDomain

import (
	"GoBaby/internal/models"
	"GoBaby/internal/utils"
	"GoBaby/ui"
	"net/http"
)

func RegisterView(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"html/base.html",
		"html/pages/auth/register.tmpl.html",
	}
	utils.CheckIfPath(w, r, models.RoutesInstance.REGISTER)
	utils.ParseTemplateFiles(w, "base", utils.EmptyStruct, utils.EmptyFuncMap, ui.Content, files...)
}
