package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("src/bookStore/views/static"))))
	http.HandleFunc("/login", serverHandler)
	http.ListenAndServe(":9998", nil)
}

func serverHandler(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintln(w,"front-page")

	t, err := template.ParseFiles("src/bookStore/views/pages/user/login.html")

	if err != nil {
		fmt.Println("err===", err)
		return
	}

	t.Execute(w, "")
}

//func staticHandler(w http.ResponseWriter, r *http.Request)  {
//	http.Handl
//}
