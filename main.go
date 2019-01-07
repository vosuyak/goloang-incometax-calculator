package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/vosuyak/montlybudget/expensemath"
	router "github.com/vosuyak/montlybudget/routers"
	"github.com/vosuyak/montlybudget/taxmath"
)

type User struct {
	ID                   string           `json:"id"`
	Firstname            string           `json:"firstname"`
	Lastname             string           `json:"lastname"`
	Fullname             string           `json:"fullname"`
	Salary               float64          `json:"salary"`
	AfterTaxSalary       float64          `json:"aftertaxsalary"`
	FederalTax           float64          `json:"federaltax"`
	MonthlyBeforeTax     float64          `json:"monthlybeforetax"`
	MonthlyAfterTax      float64          `json:"monthlyaftertax"`
	MonthlyNetIncome     float64          `json:"montlynetincome"`
	TotalMonthlyExpenses float64          `json:"totalmonthlyexpenses"`
	Taxes                *Taxes           `json:"taxes"`
	Expenses             []MonthlyExpense `json:"expenses"`
}

type Taxes struct {
	FederalTax        float64 `json:"federaltax"`
	TaxRate           float64 `json:"taxrate"`
	Fica              float64 `json:"fica"`
	Filing            string  `json:"filing"`
	Plus              float64 `json:"plus"`
	StandardDeduction int     `json:"standarddeduction"`
}

type MonthlyExpense struct {
	Expense       string  `json:"expense"`
	ExpenseAmount float64 `json:"expenseamount"`
}

type SetUserInfo interface {
	UserFullName() string
	UserSalary() float64
	UserAfterTaxSalary() float64
	UserFederalTax() float64
	UserTotalMonthlyExpense() float64
	UserMonthlyBeforeTax() float64
	UserMonthlyAfterTax() float64
	UserMonthlyNetIncome() float64
	TaxesFederal() float64
	TaxesTaxRate() float64
	TaxesFica() float64
	TaxesPlus() float64
	TaxesStandardDeduction() int
	TaxesFiling() string
}

func (p *User) UserFullName() string {
	p.Fullname = p.Firstname + " " + p.Lastname
	return p.Fullname
}
func (p *User) UserSalary() float64 {
	p.Salary = taxmath.Salary(p.Salary)
	return p.Salary
}
func (p *User) UserAfterTaxSalary() float64 {
	p.AfterTaxSalary = taxmath.AfterTaxSalary(p.Salary, p.Taxes.Filing)
	return p.AfterTaxSalary
}
func (p *User) UserFederalTax() float64 {
	p.FederalTax = taxmath.FederalTax(p.Salary, p.Taxes.Filing)
	return p.FederalTax
}

func (p *User) UserTotalMonthlyExpense() float64 {

	jsonParser := p.Expenses
	var total float64 = 0.00
	for _, expense := range jsonParser {
		total += expense.ExpenseAmount
	}
	p.TotalMonthlyExpenses = total
	return total
}

func (p *User) UserMonthlyBeforeTax() float64 {
	p.MonthlyBeforeTax = expensemath.MonthlyBeforeTax(p.Salary)
	return p.MonthlyBeforeTax
}

func (p *User) UserMonthlyAfterTax() float64 {
	p.MonthlyAfterTax = expensemath.MonthlyAfterTax(p.AfterTaxSalary)
	return p.MonthlyAfterTax
}

func (p *User) UserMonthlyNetIncome() float64 {
	p.MonthlyNetIncome = expensemath.MonthlyNetIncome(p.MonthlyAfterTax, p.TotalMonthlyExpenses)
	return p.MonthlyNetIncome
}
func UpdateUserInfo(p SetUserInfo) {
	p.UserFullName()
	p.UserSalary()
	p.UserAfterTaxSalary()
	p.UserFederalTax()
	p.UserTotalMonthlyExpense()
	p.UserMonthlyBeforeTax()
	p.UserMonthlyAfterTax()
	p.UserMonthlyNetIncome()
	fmt.Println("Full Name: ", p.UserFullName())
	fmt.Println("Salary: ", p.UserSalary())
	fmt.Println("After Tax Salary: ", p.UserAfterTaxSalary())
	fmt.Println("Federal Tax: ", p.UserFederalTax())
	fmt.Println("Total Monthly Expense: ", p.UserTotalMonthlyExpense())
	fmt.Println("Monthly Before Tax: ", p.UserMonthlyBeforeTax())
	fmt.Println("Monthly After Tax:", p.UserMonthlyAfterTax())
	fmt.Println(":", p.UserMonthlyNetIncome())
}
func (p *User) TaxesFederal() float64 {
	p.Taxes.FederalTax = taxmath.FederalTax(p.Salary, p.Taxes.Filing)
	return p.Taxes.FederalTax
}

func (p *User) TaxesTaxRate() float64 {
	p.Taxes.TaxRate = taxmath.TaxRate(p.Salary, p.Taxes.Filing)
	return p.Taxes.TaxRate
}

func (p *User) TaxesFica() float64 {
	p.Taxes.Fica = taxmath.Fica(p.Salary, p.Taxes.Filing)
	return p.Taxes.Fica
}
func (p *User) TaxesFiling() string {
	p.Taxes.Filing = p.Taxes.Filing
	return p.Taxes.Filing
}
func (p *User) TaxesPlus() float64 {
	p.Taxes.Plus = taxmath.TaxPlus(p.Salary, p.Taxes.Filing)
	return p.Taxes.Plus
}
func (p *User) TaxesStandardDeduction() int {
	p.Taxes.StandardDeduction = taxmath.StandardDeduction(p.Salary, p.Taxes.Filing)
	return p.Taxes.StandardDeduction
}
func UpdateTaxes(p SetUserInfo) {
	p.TaxesFederal()
	p.TaxesTaxRate()
	p.TaxesFica()
	p.TaxesStandardDeduction()
	p.TaxesPlus()
	fmt.Println("Federal Tax: ", p.TaxesFederal())
	fmt.Println("Tax Rate: ", p.TaxesTaxRate())
	fmt.Println("Fica: ", p.TaxesFica())
	fmt.Println("Standard Deduction: ", p.TaxesStandardDeduction())
	fmt.Println("Tax Plus: ", p.TaxesPlus())
}

var users []User

func GetSalaries(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}
func GetSalary(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, user := range users {
		if user.ID == params["id"] {
			UpdateUserInfo(&user)
			UpdateTaxes(&user)
			json.NewEncoder(w).Encode(&user)
			return
		}
	}
	json.NewEncoder(w).Encode(users)
}
func CreateSalary(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user User
	_ = json.NewDecoder(r.Body).Decode(&user)
	user.ID = strconv.Itoa(rand.Intn(100000))

	users = append(users, user)
	json.NewEncoder(w).Encode(user)
}
func UpdateSalary(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for index, user := range users {

		if user.ID == params["id"] {
			users = append(users[:index], users[index+1:]...)

			var user User
			_ = json.NewDecoder(r.Body).Decode(&user)
			UpdateUserInfo(&user)
			UpdateTaxes(&user)
			fmt.Println(&user)
			users = append(users, user)
			json.NewEncoder(w).Encode(users)
			return
		}
	}
	json.NewEncoder(w).Encode(users)
}
func DeleteSalary(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for _, user := range users {
		if user.ID == params["id"] {
			users = append(users, user)
			break
		}
	}
	json.NewEncoder(w).Encode(users)
}

func main() {
	taxmath.Bracket("single2018")
	users = append(users, User{
		ID:                   "54321",
		Firstname:            "Victor",
		Lastname:             "Osuyak",
		Fullname:             "",
		Salary:               600000.00,
		AfterTaxSalary:       00.00,
		FederalTax:           00.00,
		MonthlyBeforeTax:     00.00,
		MonthlyAfterTax:      00.00,
		MonthlyNetIncome:     00.00,
		TotalMonthlyExpenses: 00.00,
		Taxes: &Taxes{
			FederalTax:        00.00,
			TaxRate:           00.00,
			Fica:              00.00,
			Filing:            "single2018",
			Plus:              00.00,
			StandardDeduction: 00,
		},
		Expenses: []MonthlyExpense{
			MonthlyExpense{
				Expense:       "car",
				ExpenseAmount: 300.00,
			},
			MonthlyExpense{
				Expense:       "car",
				ExpenseAmount: 800.00,
			},
			MonthlyExpense{
				Expense:       "car",
				ExpenseAmount: 1800.00,
			},
			MonthlyExpense{
				Expense:       "car",
				ExpenseAmount: 800.00,
			},
		},
	})
	users = append(users, User{
		ID:                   "65432",
		Firstname:            "Ayana",
		Lastname:             "Heard",
		Fullname:             "",
		Salary:               160000.00,
		AfterTaxSalary:       00.00,
		FederalTax:           00.00,
		MonthlyBeforeTax:     00.00,
		MonthlyAfterTax:      00.00,
		MonthlyNetIncome:     00.00,
		TotalMonthlyExpenses: 00.00,
		Taxes: &Taxes{
			FederalTax:        00.00,
			TaxRate:           00.00,
			Fica:              00.00,
			Filing:            "marriedjointly2018",
			Plus:              00.00,
			StandardDeduction: 00,
		},
		Expenses: []MonthlyExpense{
			MonthlyExpense{
				Expense:       "car",
				ExpenseAmount: 300.00,
			},
			MonthlyExpense{
				Expense:       "car",
				ExpenseAmount: 800.00,
			},
			MonthlyExpense{
				Expense:       "car",
				ExpenseAmount: 1800.00,
			},
			MonthlyExpense{
				Expense:       "car",
				ExpenseAmount: 800.00,
			},
		},
	})

	// victor := User{
	// 	ID:                   "65432",
	// 	Firstname:            "Victor",
	// 	Lastname:             "Osuyak",
	// 	Fullname:             "",
	// 	Salary:               600000.00,
	// 	AfterTaxSalary:       00.00,
	// 	FederalTax:           00.00,
	// 	MonthlyBeforeTax:     00.00,
	// 	MonthlyAfterTax:      00.00,
	// 	MonthlyNetIncome:     00.00,
	// 	TotalMonthlyExpenses: 00.00,
	// 	Taxes: &Taxes{
	// 		FederalTax:        00.00,
	// 		TaxRate:           00.00,
	// 		Fica:              00.00,
	// 		Filing:            "single2018",
	// 		Plus:              00.00,
	// 		StandardDeduction: 00,
	// 	},
	// 	Expenses: []MonthlyExpense{
	// 		MonthlyExpense{
	// 			Expense:       "car",
	// 			ExpenseAmount: 300.00,
	// 		},
	// 		MonthlyExpense{
	// 			Expense:       "car",
	// 			ExpenseAmount: 800.00,
	// 		},
	// 		MonthlyExpense{
	// 			Expense:       "car",
	// 			ExpenseAmount: 1800.00,
	// 		},
	// 		MonthlyExpense{
	// 			Expense:       "car",
	// 			ExpenseAmount: 800.00,
	// 		},
	// 	},
	// }
	// UpdateUserInfo(&victor)
	// UpdateTaxes(&victor)

	// users = append(users, victor)
	// r := mux.NewRouter()

	// r.HandleFunc("/api/persons", GetSalaries).Methods("GET")
	// r.HandleFunc("/api/persons/{id}", GetSalary).Methods("GET")
	// r.HandleFunc("/api/persons", CreateSalary).Methods("POST")
	// r.HandleFunc("/api/persons/{id}", UpdateSalary).Methods("PUT")
	// r.HandleFunc("/api/persons/{id}", DeleteSalary).Methods("DELETE")

	// fmt.Println(victor)
	router := router.InitRoutes()
	server := &http.Server{
		Addr:    ":9000",
		Handler: router,
	}
	// http.ListenAndServe(":9000", router)
	server.ListenAndServe()
}
