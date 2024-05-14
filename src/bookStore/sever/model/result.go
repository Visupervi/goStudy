package model

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type BookStoreResult struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
}

func (bsr *BookStoreResult) ResponseMsg(w http.ResponseWriter, data interface{}, status int) (str string, err error) {
	w.WriteHeader(status)
	bsr.Status = status
	bsr.Data = data
	res, err := json.Marshal(bsr)
	if err != nil {
		fmt.Println("ResponseMsg json.Marshal Error=", err)
		return "", err
	}
	return string(res), nil
}
