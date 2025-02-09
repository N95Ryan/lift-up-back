package main

import (
	"testing"
)

// Teste la conversion de lbs vers kg
func TestConvertWeight_LbsToKg(t *testing.T) {
	result, unit, err := convertWeight(100, "lbs")
	expected := 45.3592
	expectedUnit := "kg"

	if err != nil {
		t.Errorf("Erreur inattendue : %v", err)
	}

	if unit != expectedUnit {
		t.Errorf("Unité incorrecte : attendu %s, obtenu %s", expectedUnit, unit)
	}

	if result != expected {
		t.Errorf("Résultat incorrect : attendu %.4f, obtenu %.4f", expected, result)
	}
}

// Teste la conversion de kg vers lbs
func TestConvertWeight_KgToLbs(t *testing.T) {
	result, unit, err := convertWeight(50, "kg")
	expected := 110.231
	expectedUnit := "lbs"

	if err != nil {
		t.Errorf("Erreur inattendue : %v", err)
	}

	if unit != expectedUnit {
		t.Errorf("Unité incorrecte : attendu %s, obtenu %s", expectedUnit, unit)
	}

	if result != expected {
		t.Errorf("Résultat incorrect : attendu %.3f, obtenu %.3f", expected, result)
	}
}

// Teste une unité invalide
func TestConvertWeight_InvalidUnit(t *testing.T) {
	_, _, err := convertWeight(100, "invalid")

	if err == nil {
		t.Errorf("Une erreur était attendue pour une unité invalide, mais aucune erreur n'a été retournée.")
	}
}
