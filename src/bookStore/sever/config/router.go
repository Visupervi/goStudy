package config

import (
	"bookStore/sever/controller"
	"net/http"
)

func RouterHandle() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("src/bookStore/views/static"))))
	http.Handle("/pages/", http.StripPrefix("/pages/", http.FileServer(http.Dir("src/bookStore/views/pages"))))
	//http.Handle("/books/", http.StripPrefix("/books/", http.FileServer(http.Dir("src/bookStore/views/pages"))))
	http.HandleFunc("/index", controller.ServerHandler)
	http.HandleFunc("/login", controller.Login)
	http.HandleFunc("/registry", controller.Registry)
	http.HandleFunc("/getBooks", controller.GetBookSlice)
	http.HandleFunc("/addBook", controller.InsertBooks)
	http.ListenAndServe(":9998", nil)
}
