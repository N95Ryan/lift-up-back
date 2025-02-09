package main

import "fmt"

// convertWeight convertit un poids entre livres (lbs) et kilogrammes (kg)
func convertWeight(value float64, unit string) (float64, string, error) {
	if unit == "lbs" {
		return value * 0.453592, "kg", nil
	} else if unit == "kg" {
		return value * 2.20462, "lbs", nil
	} else {
		return 0, "", fmt.Errorf("unité non valide : %s. Utilisez 'lbs' ou 'kg'", unit)
	}
}

func main() {
	// Exemple d'utilisation
	weight, newUnit, err := convertWeight(100, "lbs")
	if err != nil {
		fmt.Println("Erreur :", err)
	} else {
		fmt.Printf("100 lbs = %.2f %s\n", weight, newUnit)
	}

	weight, newUnit, err = convertWeight(45.5, "kg")
	if err != nil {
		fmt.Println("Erreur :", err)
	} else {
		fmt.Printf("45.5 kg = %.2f %s\n", weight, newUnit)
	}
}
