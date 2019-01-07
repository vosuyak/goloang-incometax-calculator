package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/vosuyak/montlybudget/models"
)

var expensesReq []models.MonthlyExpensesReq

var expensesRes models.MonthlyExpensesRes

// create expense POST
func PostExpense(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var expense models.MonthlyExpensesReq

	err := json.NewDecoder(r.Body).Decode(&expense)
	if err != nil {
		fmt.Println(err)
	}
	expensesReq = append(expensesReq, expense)
	json.NewEncoder(w).Encode(expense)
}

// get expenses GET
func GetExpenses(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	expensesRes.Expenses = expensesReq
	models.UpdateExpense(&expensesRes)

	json.NewEncoder(w).Encode(expensesRes)

}

// delete expense DELETE
// func DeleteExpense(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")

// 	params := mux.Vars(r)
// 	// expenses := expensesRes.Expenses
// 	for index, expense := range expensesRes.Expenses {
// 		if expense.Expense == params["expense"] {
// 			expensesRes.Expenses = append(expensesRes.Expenses[:index], expensesRes.Expenses[index+1:]...)
// 		}
// 	}
// 	json.NewEncoder(w).Encode(expensesRes)
// }

// func InitRoutes() *mux.Router {
// func ExpensesRouter(router *mux.Router) *mux.Router {
// 	r := mux.NewRouter()

// 	r.HandleFunc("/api/expenses", GetExpenses).Methods("GET")
// 	r.HandleFunc("/api/expense", PostExpense).Methods("POST")
// 	// r.HandleFunc("/api/expense/{expense}", DeleteExpense).Methods("DELETE")

// 	return r
// }
