package calculator

import (
	"fmt"

	"github.com/vosuyak/montlybudget/expensemath"
	"github.com/vosuyak/montlybudget/taxmath"
)

func FederalIncomeTaxBracket(salary float64, filing string) {
	taxmath.Salary(salary)
	taxrate := taxmath.TaxRate(salary, filing)
	taxplus := taxmath.TaxPlus(salary, filing)
	taxowed := taxmath.FederalTax(salary, filing)
	aftertaxsalary := taxmath.AfterTaxSalary(salary, filing)
	monthlyBeforeTax := expensemath.MonthlyBeforeTax(salary)
	fmt.Println("Salary: ", salary)
	fmt.Println("taxrate: ", taxrate)
	fmt.Println("taxplus: ", taxplus)
	fmt.Println("Taxes: ", taxowed)
	fmt.Println("After Tax: ", aftertaxsalary)
	fmt.Println("monthlyBeforeTax: ", monthlyBeforeTax)
	fmt.Println("--------")
}

func AfterTaxIncome(salary float64, filing string) {
	taxmath.AfterTaxSalary(salary, filing)
}
