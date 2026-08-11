package main

import (
	"flag"
	"fmt"

	"github.com/isaacvarg/prop-65-daily-exposure/product"
)

func main() {
	mcg := "µg"
	unit := mcg + "/day"
	// this really should be only set to 1.00 for a conservative calculation
	// anything less would require rebust toxicological dermal penetration testing
	dermalAbsorbtionFactor := 1.00
	mc, cc := getOptions()

	if mc == -1.00 {
		mc = getUserInput("Please enter material concentration as w/w%: ")
	}

	if cc == -1.00 {
		cc = getUserInput("Please enter the concentration of the Proposition 65 chemical in the raw material" + "(" + mcg + "): ")
	}

	chemicalConcentration := getChemicalConcFinishedProduct(mc, cc)

	product, err := product.PromptProductCategoryChoice()
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	usageResult := getDailyUsage(chemicalConcentration, product, dermalAbsorbtionFactor)

	fmt.Printf("Daily exposure: %.4f %s\n", usageResult, unit)
}

func getDailyUsage(chemicalConcentration float64, product product.Product, dermalAbsorbtionFactor float64) float64 {
	return chemicalConcentration * product.AmountAppliedDaily * product.RententionFactor * dermalAbsorbtionFactor
}

// getChemicalConcFinishedProduct gets the concentration of
// the propisition 65 risidual chemical, originating from the
// raw material, in the final product
func getChemicalConcFinishedProduct(mc, cc float64) float64 {
	return (mc * 1e-2) * cc
}

func getUserInput(promptText string) float64 {
	fmt.Print(promptText)
	var value float64
	_, err := fmt.Scanln(&value)
	if err != nil {
		fmt.Println("error getting user input: ", err)
	}
	return value
}

func getOptions() (mc, cc float64) {
	var materialConcentration float64
	var chemicalConcentration float64

	flag.Float64Var(&materialConcentration, "m", -1.00, "The concentration of the raw material containing the Prop. 65 chemical")

	flag.Float64Var(&chemicalConcentration, "c", -1.00, "The conentration of the Prop 65 chemical in the raw material")

	flag.Parse()

	return materialConcentration, chemicalConcentration
}
