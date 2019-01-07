package routers

import (
	"github.com/gorilla/mux"
	"github.com/vosuyak/montlybudget/controllers"
)

func ExpensesRouter(router *mux.Router) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/api/expenses", controllers.GetExpenses).Methods("GET")
	r.HandleFunc("/api/expense", controllers.PostExpense).Methods("POST")
	// r.HandleFunc("/api/expense/{expense}", DeleteExpense).Methods("DELETE")

	return r
}
