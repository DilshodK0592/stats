package stats

import (
	"fmt"

	"github.com/DilshodK0592/bank/v2/pkg/types"
)

func ExampleAvg() {
	pay := []types.Payment {
		{
			Amount: 1_000_00,
			Status: "OK",
		},
		{
			Amount: 2_000_00,
			Status: "FAIL",
		},
		{
			Amount: 5_000_00,
			Status: "INPROGRESS",
		},
	}
	result := Avg(pay)
	fmt.Println(result)
	//Output: 300000
}

func ExampleTotalInCategory() {
	pays := []types.Payment {
		{
			Category: "auto",
			Amount: 2_000_00,
			Status: "OK",
		},
		{
			Category: "auto",
			Amount: 3_000_00,
			Status: "INPROGRESS",
		},
		{
			Category: "auto",
			Amount: 1_000_00,
			Status: "FAIL",
		},
	}

	result := TotalInCategory(pays, "auto")
	
	fmt.Println(result)
	//Output: 200000
}