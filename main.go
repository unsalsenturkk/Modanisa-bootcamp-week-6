package main

import (
	"fmt"
	"net/http"
	"week-6-assignment/controller"
	"week-6-assignment/database"
	"week-6-assignment/models"
	"week-6-assignment/service"
)

func main() {
	in_memory := make(models.DataResponse)
	db := database.NewDatabase(in_memory)
	svc := service.NewWalletService(db)
	handler := controller.NewWalletController(svc)

	http.HandleFunc("/", handler.WalletHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}

}
