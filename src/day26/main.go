package main

import (
	"day26/service"
	"day26/view"
)

func main() {
	cv := view.CustomerView{Key: "", Loop: true}
	cv.CustomerService = service.InitSplice()
	cv.GetMainMenu()
}
