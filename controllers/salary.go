package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/vosuyak/montlybudget/models"
)

// model variable
var salaryReq []models.SalaryReq
var salaryRes models.SalaryRes

// create expense POST
func PostSalary(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var salary models.SalaryReq

	err := json.NewDecoder(r.Body).Decode(&salary)
	if err != nil {
		fmt.Println(err)
	}
	salaryReq = append(salaryReq, salary)
	json.NewEncoder(w).Encode(salary)
}

// get expenses GET
func GetSalary(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	models.UpdateTaxes(&salaryRes)
	json.NewEncoder(w).Encode(salaryRes)
}

// update expense PUT
// delete expense DELETE
