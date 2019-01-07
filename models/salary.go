package models

import (
	"fmt"

	"github.com/vosuyak/montlybudget/taxmath"
)

type (
	SalaryReq struct {
		Salary float64 `json:"salary"`
		Filing string  `json:"filing"`
	}
	SalaryRes struct {
		Finance          *SalaryReq `json:"finance"`
		AfterTaxSalary   float64    `json:"aftertaxsalary"`
		FederalTax       float64    `json:"federaltax"`
		MonthlyBeforeTax float64    `json:"monthlybeforetax"`
		MonthlyAfterTax  float64    `json:"monthlyaftertax"`
		// MonthlyNetIncome     float64 `json:"montlynetincome"`
		// TotalMonthlyExpenses float64 `json:"totalmonthlyexpenses"`
		Taxes *TaxesRes `json:"taxes"`
		// Expenses             []MonthlyExpense `json:"expenses"`
	}

	TaxesRes struct {
		FederalTax        float64 `json:"federaltax"`
		TaxRate           float64 `json:"taxrate"`
		Fica              float64 `json:"fica"`
		Filing            string  `json:"filing"`
		Plus              float64 `json:"plus"`
		StandardDeduction int     `json:"standarddeduction"`
	}
)

// type MonthlyExpense struct {
// 	Expense       string  `json:"expense"`
// 	ExpenseAmount float64 `json:"expenseamount"`
// }
func (o *SalaryRes) Salary() (float64, string) {
	o.Finance.Salary = 68000.00
	o.Finance.Filing = "single2018"
	return o.Finance.Salary, o.Finance.Filing
}
func (o *SalaryRes) TaxesFederal() float64 {
	o.FederalTax = taxmath.FederalTax(68000.00, "single2018")
	fmt.Println("fed: ", o.FederalTax)

	return o.FederalTax
}

func (o *SalaryRes) TaxesTaxRate() float64 {
	o.Taxes.TaxRate = taxmath.TaxRate(o.Finance.Salary, o.Finance.Filing)
	return o.Taxes.TaxRate
}

func (o *SalaryRes) TaxesFica() float64 {
	o.Taxes.Fica = taxmath.Fica(o.Finance.Salary, o.Finance.Filing)
	return o.Taxes.Fica
}
func (o *SalaryRes) TaxesFiling() string {
	o.Taxes.Filing = o.Taxes.Filing
	return o.Taxes.Filing
}
func (o *SalaryRes) TaxesPlus() float64 {
	o.Taxes.Plus = taxmath.TaxPlus(o.Finance.Salary, o.Finance.Filing)
	return o.Taxes.Plus
}
func (o *SalaryRes) TaxesStandardDeduction() int {
	o.Taxes.StandardDeduction = taxmath.StandardDeduction(o.Finance.Salary, o.Finance.Filing)
	return o.Taxes.StandardDeduction
}

type SetTaxesInfo interface {
	Salary() (float64, string)
	TaxesFederal() float64
	TaxesTaxRate() float64
	TaxesFica() float64
	TaxesFiling() string
	TaxesPlus() float64
	TaxesStandardDeduction() int
}

func UpdateTaxes(p SetTaxesInfo) {
	p.Salary()
	p.TaxesFederal()
	p.TaxesTaxRate()
	p.TaxesFica()
	p.TaxesFiling()
	p.TaxesPlus()
	p.TaxesStandardDeduction()
	fmt.Println("Federal Tax: ", p.TaxesFederal())
	// fmt.Println("Tax Rate: ", p.TaxesTaxRate())
	// fmt.Println("Fica: ", p.TaxesFica())
	// fmt.Println("Standard Deduction: ", p.TaxesStandardDeduction())
	// fmt.Println("Tax Plus: ", p.TaxesPlus())
}
