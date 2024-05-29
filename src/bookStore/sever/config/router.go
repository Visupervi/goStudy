package config

import (
	"bookStore/sever/controller"
	"net/http"
)

func RouterHandle() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("src/bookStore/views/static"))))
	http.Handle("/pages/", http.StripPrefix("/pages/", http.FileServer(http.Dir("src/bookStore/views/pages"))))
	//http.Handle("/books/", http.StripPrefix("/books/", http.FileServer(http.Dir("src/bookStore/views/pages"))))
	//http.HandleFunc("/setCookies", controller.SetCookie)
	http.HandleFunc("/index", controller.ServerHandler)
	http.HandleFunc("/login", controller.Login)
	http.HandleFunc("/registry", controller.Registry)
	http.HandleFunc("/getBooks", controller.GetBookSlice)
	http.HandleFunc("/addBook", controller.InsertBooks)
	http.HandleFunc("/getAllBooks", controller.GetBooks)
	http.HandleFunc("/deleteBook", controller.DeleteBook)
	http.HandleFunc("/updateBook", controller.UpdateBook)
	http.HandleFunc("/getCart", controller.GetCart)
	http.ListenAndServe(":9998", nil)
}
