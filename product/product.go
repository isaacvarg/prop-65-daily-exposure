// Package product provides usage and rentention factor for various product categories
package product

import "fmt"

type Product struct {
	AmountAppliedDaily float64
	RententionFactor   float64
}

// TODO turn this into a slice and then generate the
// menu options dynamically,

// these values are derived from the
// scientific committee on consumer safety (SCCS)
// THE SCCS NOTES OF GUIDANCE FOR THE TESTING OF
// COSMETIC INGREDIENTS AND THEIR SAFETY EVALUATION
// 12TH REVISION page 27 Table 3A.
// url: https://health.ec.europa.eu/document/download/32a999f7-d820-496a-b659-d8c296cc99c1_en?filename=sccs_o_273_final.pdf
var products = map[string]Product{
	"shower-gel":          {AmountAppliedDaily: 18.64, RententionFactor: 0.01},
	"shampoo":             {AmountAppliedDaily: 10.46, RententionFactor: 0.01},
	"hair-styling":        {AmountAppliedDaily: 4.00, RententionFactor: 0.10},
	"body-lotion":         {AmountAppliedDaily: 7.82, RententionFactor: 1.00},
	"face-cream":          {AmountAppliedDaily: 1.54, RententionFactor: 1.00},
	"hand-cream":          {AmountAppliedDaily: 2.16, RententionFactor: 1.00},
	"liquid-foundation":   {AmountAppliedDaily: 0.51, RententionFactor: 1.00},
	"lip-products":        {AmountAppliedDaily: 0.057, RententionFactor: 1.00},
	"deodorant-non-spray": {AmountAppliedDaily: 6.54, RententionFactor: 1.00},
	"deodorant-spray":     {AmountAppliedDaily: 6.54, RententionFactor: 1.00},
}

func GetProductUsage(productType string) (Product, bool) {
	prod, exists := products[productType]
	return prod, exists
}

func PromptProductCategoryChoice() (Product, error) {
	fmt.Println("\nChoose a product category:")
	fmt.Println("1. Shower gel")
	fmt.Println("2. Shampoo")
	fmt.Println("3. Hair styling products")
	fmt.Println("4. Body lotion")
	fmt.Println("5. Face cream")
	fmt.Println("6. Hand cream")
	fmt.Println("7. Liquid foundation")
	fmt.Println("8. Lip products")
	fmt.Println("9. Deodorant (non-spray)")
	fmt.Println("10. Deodorant Dpray")
	fmt.Print("\nWhat is your choice?: ")

	var choice int
	_, err := fmt.Scanln(&choice)
	if err != nil {
		return Product{}, fmt.Errorf("invalid input format")
	}

	var key string
	switch choice {
	case 1:
		key = "shower-gel"
	case 2:
		key = "shampoo"
	case 3:
		key = "hair-styling"
	case 4:
		key = "body-lotion"
	case 5:
		key = "face-cream"
	case 6:
		key = "hand-cream"
	case 7:
		key = "liquid-foundation"
	case 8:
		key = "lip-products"
	case 9:
		key = "deodorant-non-spray"
	case 10:
		key = "deodorant-spray"
	default:
		return Product{}, fmt.Errorf("invalid option selected: %d", choice)
	}

	prod, _ := GetProductUsage(key)

	return prod, nil
}
