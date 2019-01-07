package routers

import (
	"github.com/gorilla/mux"
	"github.com/vosuyak/montlybudget/controllers"
)

func SalaryRouter(router *mux.Router) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/api/get/salary", controllers.GetSalary).Methods("GET")
	r.HandleFunc("/api/post/salary", controllers.PostSalary).Methods("POST")
	// r.HandleFunc("/api/expense/{expense}", DeleteExpense).Methods("DELETE")

	return r
}
