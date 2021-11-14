package stats

import (
	"fmt"

	"github.com/DilshodK0592/bank/pkg/types"
)

func ExampleAvg() {
	pay := []types.Payment {
		{
			Amount: 1_000_00,
		},
		{
			Amount: 2_000_00,
		},
		{
			Amount: 3_000_00,
		},
	}
	result := Avg(pay)
	fmt.Println(result)
	//Output: 200000
}

func ExampleTotalInCategory() {
	pays := []types.Payment {
		{
			Category: "auto",
			Amount: 1_000_00,
		},
		{
			Category: "auto",
			Amount: 1_000_00,
		},
		{
			Category: "food",
			Amount: 1_000_00,
		},
	}

	result := TotalInCategory(pays, "auto")
	
	fmt.Println(result)
	//Output: 200000
}