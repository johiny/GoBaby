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

func LoginAction(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Error al parsear el formulario", http.StatusInternalServerError)
		return
	}
	Email := r.Form.Get("Email")
	Password := r.Form.Get("Password")
	println(Email, Password)
}
