package main

import (
	"log"
	"net/http"
	"yap-arch-demo/internal/remote"
	handler "yap-arch-demo/internal/remote/handlers/employee"
	service "yap-arch-demo/internal/service/employee"
	"yap-arch-demo/internal/store"
)

func main() {
	empStore := store.NewEmployeeStore()
	empService := service.NewEmployeeService(empStore)
	empHandler := handler.NewEmployeeHandler(empService)
	r := remote.NewRouter(empHandler)
	log.Fatal(http.ListenAndServe(":8888", r))
}
