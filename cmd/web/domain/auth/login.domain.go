package authDomain

import (
	"GoBaby/internal/models"
	"GoBaby/internal/utils"
	"GoBaby/ui"
	"net/http"
)

func LoginView(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"html/base.html",
		"html/pages/auth/login.tmpl.html",
	}
	utils.CheckIfPath(w, r, models.RoutesInstance.LOGIN)
	utils.ParseTemplateFiles(w, "base", utils.EmptyStruct, utils.EmptyFuncMap, ui.Content, files...)
}
