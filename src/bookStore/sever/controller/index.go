package controller

import (
	"bookStore/sever/service"
	"html/template"
	"net/http"
)

// ServerHandler 首页
func ServerHandler(w http.ResponseWriter, r *http.Request) {
	var t *template.Template
	cookie, _ := r.Cookie("user")
	if cookie != nil {
		cookieValue := cookie.Value
		sess, _ := service.GetSessionById(cookieValue)

		if sess != nil {
			t = template.Must(template.ParseFiles("src/bookStore/views/pages/user/index.html"))
		} else {
			t = template.Must(template.ParseFiles("src/bookStore/views/pages/user/login.html"))
		}
	} else {
		t = template.Must(template.ParseFiles("src/bookStore/views/pages/user/login.html"))

	}
	//t := template.Must(template.ParseFiles("src/bookStore/views/pages/user/index.html"))
	t.Execute(w, "")
}
