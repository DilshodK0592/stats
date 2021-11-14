package stats

import "github.com/DilshodK0592/bank/pkg/types"

// Avg рассчитывает среднюю сумму платежа.
func Avg(payments []types.Payment) types.Money {
	pay:= types.Money(0)
	for _, payment := range payments {
		pay += payment.Amount
	}
	averagePay := pay/types.Money(len(payments))
	return averagePay
}

// TotalInCategory находит сумму покупок в опредённой категории.
func TotalInCategory(payments []types.Payment, categoty types.Category) types.Money {
	pay:= types.Money(0)
	for _, payment := range payments {
		if categoty == payment.Category {
		pay += payment.Amount
		}
	}
	
	return pay
}