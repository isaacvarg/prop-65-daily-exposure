package main

import (
	"flag"
	"fmt"
)

func main() {
	mc, cc := getOptions()

	if mc == -1.00 {
		mc = getUserInput("Please enter material concentration: ")
	}

	if cc == -1.00 {
		cc = getUserInput("Please enter the concentration of the Proposition 65 chemical in the raw material: ")
	}

	result := getChemicalConcFinishedProduct(mc, cc)
	fmt.Printf("Finished product chemical concentration: %.4f%%\n", result)
}

// getChemicalConcFinishedProduct gets the concentration of
// the propisition 65 risidual chemical, originating from the
// raw material, in the final product
func getChemicalConcFinishedProduct(mc, cc float64) float64 {
	return (mc / 100) * cc
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
