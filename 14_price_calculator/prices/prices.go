package prices

import (
	"fmt"

	"examples.com/price_calculator/conversions"
	"examples.com/price_calculator/filemanager"
)

type TaxincludedpriceJob struct {
	TaxRate           float64
	InputPrices       []float64
	TaxIncludedPrices map[string]float64
}

func (job *TaxincludedpriceJob) LoadData(filePath string) {
	lines, err := filemanager.ReadLines(filePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		panic(err)
	}

	prices, err := conversions.StringToFloat64(lines)
	if err != nil {
		fmt.Println("Error converting strings to float64:", err)
		panic(err)
	}
	job.InputPrices = prices
}

func (job *TaxincludedpriceJob) Process(filePath string) {
	job.LoadData(filePath)

	results := make(map[string]float64)

	for _, price := range job.InputPrices {
		results[fmt.Sprintf("price: %.2f", price)] = price * (1 + job.TaxRate)
	}

	// fmt.Printf("Tax Included Prices @ %.2f%%:\n", job.TaxRate*100)
	// for key, value := range results {
	// 	fmt.Printf("  %s: %.2f\n", key, value)
	// }

	job.TaxIncludedPrices = results

	err := filemanager.WriteJSONToFile(fmt.Sprintf("result_%.0f.json", job.TaxRate*100), job)
	if err != nil {
		fmt.Println("Error writing results to file:", err)
		panic(err)
	}
}

func NewTaxIncludedPriceJob(taxRate float64) *TaxincludedpriceJob {
	return &TaxincludedpriceJob{
		TaxRate:           taxRate,
		InputPrices:       []float64{},
		TaxIncludedPrices: make(map[string]float64),
	}
}
