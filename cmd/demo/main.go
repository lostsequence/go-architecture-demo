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
	s := store.NewEmployeeStore()
	es := service.NewEmployeeService(s)
	eh := handler.NewEmployeeHandler(es)
	r := remote.NewRouter(eh)
	log.Fatal(http.ListenAndServe(":8888", r))
}
