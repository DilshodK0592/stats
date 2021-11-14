package stats

import "github.com/DilshodK0592/bank/v2/pkg/types"

// Avg рассчитывает среднюю сумму платежа.
func Avg(payments []types.Payment) types.Money {
	pay:= types.Money(0)
	indexPay := 0
	for _, payment := range payments {
		if payment.Status != types.StatusFail {
		pay += payment.Amount
		indexPay++
		}
	}
	averagePay := pay/types.Money(indexPay)
	return averagePay
}

// TotalInCategory находит сумму покупок в опредённой категории.
func TotalInCategory(payments []types.Payment, category types.Category) types.Money {
	pay:= types.Money(0)
	for _, payment := range payments {
		if payment.Status == types.StatusFail {
			continue
		}
		if category != payment.Category {
			continue
		}
		pay += payment.Amount
	}
	
	return pay
}