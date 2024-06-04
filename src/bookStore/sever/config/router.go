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
	http.HandleFunc("/addCart", controller.AddCart)
	http.HandleFunc("/clearCart", controller.ClearCart)
	http.HandleFunc("/deleteCartItem", controller.DeleteCartItemById)
	http.HandleFunc("/updateCartItem", controller.UpdateCartItemById)
	http.HandleFunc("/checkout", controller.Checkout)
	http.HandleFunc("/getOrderList", controller.GetOrderList)
	http.HandleFunc("/orderDetail", controller.OrderDetail)
	http.HandleFunc("/updateState", controller.UpdateOrderState)
	http.HandleFunc("/getMyOrders", controller.GetMyOrders)
	http.ListenAndServe(":9998", nil)
}
