package expensemath

func MonthlyBeforeTax(salary float64) float64 {
	const months = float64(12)

	monthlyBeforeTax := salary / months
	return monthlyBeforeTax
}
func MonthlyAfterTax(afterTaxSalary float64) float64 {
	const months = float64(12)

	monthlyAfterTax := afterTaxSalary / months
	return monthlyAfterTax
}
func TotalMonthlyExpenses(arr []float64, size int) float64 {
	var i int
	total := 0.00

	for i = 0; i < size; i++ {
		total += arr[i]
	}

	totalMonthlyExpenses := total
	return float64(totalMonthlyExpenses)
}

func MonthlyNetIncome(monthlyAfterTax, totalMonthlyExpenses float64) float64 {
	return monthlyAfterTax - totalMonthlyExpenses
}
