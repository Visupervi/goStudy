package main

import (
	"bookStore/sever/db"
	"bookStore/sever/service"
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	db.InitDB()
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("src/bookStore/views/static"))))
	http.Handle("/pages/", http.StripPrefix("/pages/", http.FileServer(http.Dir("src/bookStore/views/pages"))))
	http.HandleFunc("/index", serverHandler)
	http.HandleFunc("/login", service.LoginService)
	http.HandleFunc("/registry", service.RegistryService)
	http.ListenAndServe(":9998", nil)
}

// 首页
func serverHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("src/bookStore/views/pages/user/index.html")

	if err != nil {
		fmt.Println("err===", err)
		return
	}

	t.Execute(w, "")
}
