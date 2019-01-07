package models

type (
	// SEND
	MonthlyExpensesReq struct {
		Expense       string  `json:"expense"`
		ExpenseAmount float64 `json:"expenseamount"`
	}

	//GET
	MonthlyExpensesRes struct {
		TotalMonthlyExpenses float64              `json:"totalmonthlyexpenses"`
		Expenses             []MonthlyExpensesReq `json:"expenses"`
	}
)

func (p *MonthlyExpensesRes) TotalExpense() float64 {

	jsonParser := p.Expenses
	var total float64 = 0.00
	for _, expense := range jsonParser {
		total += expense.ExpenseAmount
	}
	p.TotalMonthlyExpenses = total
	return total
}

type SetExpenseInfo interface {
	TotalExpense() float64
}

func UpdateExpense(p SetExpenseInfo) {
	p.TotalExpense()
}
