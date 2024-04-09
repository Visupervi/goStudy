package service

import "day26/view"

type CustomerService struct {
	customers   []view.CustomerView `json:"customers"`
	customerNum int
}
