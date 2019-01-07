package taxmath

import (
	"encoding/json"
	"fmt"
	"os"
)

type TaxBracket struct {
	Rate              float64   `json:"rate"`
	Filing            string    `json:"filing"`
	Range             []float64 `json:"range"`
	Plus              float64   `json:"plus"`
	StandardDeduction int       `json:"standarddeduction"`
	Fica              []float64 `json:"fica"`
	State             string    `json:"state"`
	Local             string    `json:"local"`
}

func Bracket(filing string) []TaxBracket {
	var bracket []TaxBracket
	filejson, err := os.Open("./common/" + filing + ".json")

	defer filejson.Close()
	if err != nil {
		fmt.Println(err)
	}
	json.NewDecoder(filejson).Decode(&bracket)
	return bracket
}

func Salary(salary float64) float64 {
	return salary
}
func TaxRate(salary float64, filing string) float64 {
	percentage := 00.00

	brackets := Bracket(filing)
	for _, v := range brackets {
		if v.Filing == filing {
			if salary > v.Range[0] && salary < v.Range[1] {
				percentage = v.Rate
			}
		}
	}
	// if filing == "single2018" {
	// 	if salary < 9525.00 {
	// 		percentage = 00.10
	// 	}
	// 	if salary > 9526.00 {
	// 		percentage = 00.12
	// 	}
	// 	if salary > 38701.00 {
	// 		percentage = 00.22
	// 	}
	// 	if salary > 82501.00 {
	// 		percentage = 00.24
	// 	}
	// 	if salary > 157501.00 {
	// 		percentage = 00.32
	// 	}
	// 	if salary > 200001.00 {
	// 		percentage = 00.35
	// 	}
	// 	if salary > 500001.00 {
	// 		percentage = 00.37
	// 	}
	// }
	// if filing == "marriedjointly2018" {
	// 	if salary < 19050.00 {
	// 		percentage = 00.10
	// 	}
	// 	if salary > 19051.00 {
	// 		percentage = 00.12
	// 	}
	// 	if salary > 77401.00 {
	// 		percentage = 00.22
	// 	}
	// 	if salary > 165001.00 {
	// 		percentage = 00.24
	// 	}
	// 	if salary > 315001.00 {
	// 		percentage = 00.32
	// 	}
	// 	if salary > 400001.00 {
	// 		percentage = 00.35
	// 	}
	// 	if salary > 600001.00 {
	// 		percentage = 00.37
	// 	}
	// }

	return percentage
}
func TaxPlus(salary float64, filing string) float64 {
	plus := 00.00

	brackets := Bracket(filing)
	for _, v := range brackets {
		if v.Filing == filing {
			if salary > v.Range[0] && salary < v.Range[1] {
				plus = v.Plus
			}
		}
	}
	// if filing == "single2018" {
	// 	if salary > 0 && salary < 9525.00 {
	// 		plus = 00.00
	// 	}
	// 	if salary > 9526.00 && salary < 38700.00 {
	// 		plus = 952.50
	// 	}
	// 	if salary > 38701.00 && salary < 82500.00 {
	// 		plus = 4453.50
	// 	}
	// 	if salary > 82501.00 && salary < 157500.00 {
	// 		plus = 14089.50
	// 	}
	// 	if salary > 157501.00 && salary < 200000.00 {
	// 		plus = 32089.50
	// 	}
	// 	if salary > 200001.00 && salary < 500000.00 {
	// 		plus = 45689.50
	// 	}
	// 	if salary > 500001.00 {
	// 		plus = 150689.50
	// 	}
	// }
	// if filing == "marriedjointly2018" {
	// 	if salary > 0 && salary < 19050.00 {
	// 		plus = 00.00
	// 	}
	// 	if salary > 19051.00 && salary < 77400.00 {
	// 		plus = 1905.50
	// 	}
	// 	if salary > 77401.00 && salary < 165000.00 {
	// 		plus = 8907.50
	// 	}
	// 	if salary > 165001.00 && salary < 315000.00 {
	// 		plus = 28179.50
	// 	}
	// 	if salary > 315001.00 && salary < 400000.00 {
	// 		plus = 64179.50
	// 	}
	// 	if salary > 400001.00 && salary < 600000.00 {
	// 		plus = 91379.50
	// 	}
	// 	if salary > 600001.00 {
	// 		plus = 161379.50
	// 	}
	// }
	return plus
}
func FederalTax(salary float64, filing string) float64 {
	output := 00.00
	standarddeduction := float64(StandardDeduction(salary, filing))
	salary = salary - standarddeduction
	brackets := Bracket(filing)
	for _, v := range brackets {
		if v.Filing == filing {
			if salary > v.Range[0] && salary < v.Range[1] {
				output = ((salary - v.Range[0]) * TaxRate(salary, filing)) + TaxPlus(salary, filing)
			}
		}
	}

	// if filing == "single2018" {
	// 	salary = salary - standarddeduction
	// 	if salary > 0 && salary < 9525.00 {
	// 		output = 9525.00 * TaxRate(salary, filing)
	// 	}
	// 	if salary > 9526.00 && salary < 38700.00 {
	// 		output = ((salary - 9526.00) * TaxRate(salary, filing)) + TaxPlus(salary, filing)
	// 	}
	// 	if salary > 38701.00 && salary < 82500.00 {
	// 		output = ((salary - 38701.00) * TaxRate(salary, filing)) + TaxPlus(salary, filing)
	// 	}
	// 	if salary > 82501.00 && salary < 157500.00 {
	// 		output = ((salary - 82501.00) * TaxRate(salary, filing)) + TaxPlus(salary, filing)
	// 	}
	// 	if salary > 157501.00 && salary < 200000.00 {
	// 		output = ((salary - 157501.00) * TaxRate(salary, filing)) + TaxPlus(salary, filing)
	// 	}
	// 	if salary > 200001.00 && salary < 500000.00 {
	// 		output = ((salary - 200001.00) * TaxRate(salary, filing)) + TaxPlus(salary, filing)
	// 	}
	// 	if salary > 500001.00 {
	// 		output = ((salary - 500001.00) * .37) + TaxPlus(salary, filing)
	// 	}
	// }
	// if filing == "marriedjointly2018" {
	// 	salary = salary - standarddeduction
	// 	if salary > 0 && salary < 19050.00 {
	// 		output = 19050.00 * TaxRate(salary, filing)
	// 	}
	// 	if salary > 19051.00 && salary < 77400.00 {
	// 		output = ((salary - 19051.00) * TaxRate(salary, filing)) + TaxPlus(salary, filing)
	// 	}
	// 	if salary > 77401.00 && salary < 165000.00 {
	// 		output = ((salary - 77401.00) * TaxRate(salary, filing)) + TaxPlus(salary, filing)
	// 	}
	// 	if salary > 165001.00 && salary < 315000.00 {
	// 		output = ((salary - 165001.00) * TaxRate(salary, filing)) + TaxPlus(salary, filing)
	// 	}
	// 	if salary > 315001.00 && salary < 400000.00 {
	// 		output = ((salary - 315001.00) * TaxRate(salary, filing)) + TaxPlus(salary, filing)
	// 	}
	// 	if salary > 400001.00 && salary < 500000.00 {
	// 		output = ((salary - 400001.00) * TaxRate(salary, filing)) + TaxPlus(salary, filing)
	// 	}
	// 	if salary > 600000.01 {
	// 		output = ((salary - 600000.01) * .37) + TaxPlus(salary, filing)
	// 	}
	// }

	return output

}
func StandardDeduction(salary float64, filing string) int {
	var standarddeduction int
	brackets := Bracket(filing)
	for _, v := range brackets {
		if salary > v.Range[0] && salary < v.Range[1] {
			if filing == "single2018" {
				standarddeduction = v.StandardDeduction
			}
			if filing == "marriedjointly2018" {
				standarddeduction = v.StandardDeduction
			}
			if filing == "single2017" {
				standarddeduction = v.StandardDeduction
			}
			if filing == "marriedjointly2017" {
				standarddeduction = v.StandardDeduction
			}
		}
	}

	return int(standarddeduction)
}
func Fica(salary float64, filing string) float64 {
	var ficaamount = 00.00
	var ss = 00.00
	var hi = 00.00
	brackets := Bracket(filing)
	for _, v := range brackets {
		if v.Filing == filing {
			if salary > v.Range[0] && salary < v.Range[1] {
				if salary > v.Fica[0] {
					ss = v.Fica[0] * v.Fica[1]
					hi = salary * v.Fica[3]
					ficaamount = hi + ss
				} else {
					ss = salary * v.Fica[1]
					hi = salary * v.Fica[3]
					ficaamount = hi + ss
				}
				if salary > v.Fica[2] {
					diff := salary - v.Fica[2]
					diff = diff * v.Fica[4]
					hi = salary * v.Fica[3]
					ficaamount = diff + hi + ss
				}
			}
		}
	}
	return ficaamount
}

func State() {

}
func Local() {

}
func AfterTaxSalary(salary float64, filing string) float64 {
	fica := Fica(salary, filing)

	brackets := Bracket(filing)
	for _, v := range brackets {
		if v.Filing == filing {
			if salary > v.Range[0] && salary < v.Range[1] {
				federaltax := FederalTax(salary, filing)
				salary = salary - (float64(federaltax) + fica)
			}
		}
	}

	// federaltax := FederalTax(salary, filing)
	// salary = salary - (float64(federaltax) + fica)
	return salary
}
