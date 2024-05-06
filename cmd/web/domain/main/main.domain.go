package mainDomain

import (
	authDomain "GoBaby/cmd/web/domain/auth"
	"GoBaby/internal/models"
	"GoBaby/internal/utils"
	"GoBaby/ui"
	"net/http"
)

func MainView(w http.ResponseWriter, r *http.Request) {
	err := authDomain.CheckSession(r)
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	utils.CheckIfPath(w, r, models.RoutesInstance.MAIN)

	files := []string{
		"html/base.html",
		"html/pages/main/main.tmpl.html",
	}

	utils.ParseTemplateFiles(w, "base", utils.EmptyStruct, utils.EmptyFuncMap, ui.Content, files...)
}
