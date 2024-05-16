package controller

import (
	"html/template"
	"net/http"
)

// ServerHandler 首页
func ServerHandler(w http.ResponseWriter, r *http.Request) {
	//t, err := template.ParseFiles("src/bookStore/views/pages/user/index.html")
	//
	//if err != nil {
	//	fmt.Println("err===", err)
	//	return
	//}

	t := template.Must(template.ParseFiles("src/bookStore/views/pages/user/index.html"))

	t.Execute(w, "")
}
