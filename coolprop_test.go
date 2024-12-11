package coolprop_go

import (
	"slices"
	"testing"
)

func TestExtractBackend(t *testing.T) {
	fluidStr := "Air"

	backend, fluid := extractBackend(fluidStr)

	if backend != "?" && fluid != "Air" {
		t.Errorf("Backend Extraction failed, (backend, fluid) = (%v, %v), expected (%v, %v)\n", backend, fluid, "?", "Air")
	} else {
		t.Logf("In fluidStr: '%v'. Backend: '%v'. Fluid: '%v'.\n", fluidStr, backend, fluid)
	}

	fluidStr = "HEOS::Air"

	backend, fluid = extractBackend(fluidStr)
	if backend != "HEOS" && fluid != "Air" {
		t.Errorf("Backend Extraction failed, (backend, fluid) = (%v, %v), expected (%v, %v)\n", backend, fluid, "HEOS", "Air")
		t.Logf("In fluidStr: '%v'. Backend: '%v'. Fluid: '%v'.\n", fluidStr, backend, fluid)
	} else {
		t.Logf("In fluidStr: '%v'. Backend: '%v'. Fluid: '%v'.\n", fluidStr, backend, fluid)
	}

}

func TestHasBackendString(t *testing.T) {
	fluidStr := "Air"
	if _, found := hasBackendInString(fluidStr); found {
		t.Errorf("hasBackendString() failed. Fluid string '%v' should have no backend", fluidStr)
	} else {
		t.Logf("In fluidStr: '%v'. Backend was not found", fluidStr)
	}

	fluidStr = "HEOS::Air"
	if index, found := hasBackendInString(fluidStr); !found || index != 4 {
		t.Errorf("hasBackendString() failed. Fluid string '%v' should have backend at index 4, but index %d was returned", fluidStr, index)
	} else {
		t.Logf("In fluidStr: '%v'. Backend was found on position: '%d'.", fluidStr, index)
	}

}

func TestHasSolutionConcentration(t *testing.T) {
	fluidStr := "Air"

	if hasSolutionConcentration(fluidStr) {
		t.Errorf("hasSolutionConcentration() failed. fluidStr: '%s' should not have solution concentration", fluidStr)
	} else {
		t.Logf("fluidStr: '%s' does not have solution concentration", fluidStr)
	}

	fluidStr = "Air-20%"
	if !hasSolutionConcentration(fluidStr) {
		t.Errorf("hasSolutionConcentration() failed. fluidStr: '%s' should have solution concentration", fluidStr)
	} else {
		t.Logf("fluidStr: '%s' does have solution concentration", fluidStr)
	}
}

func TestExtractFractions(t *testing.T) {
	fluidStr := "Methane[0.5]&Ethane[0.5]"

	if names, fractions, err := extractFractions(fluidStr); err != nil {
		t.Errorf("extractFractions() Failed. fluidStr: '%v' should have fractions", fluidStr)
	} else {
		expectedFractions := []float64{0.5, 0.5}
		expectedNames := "Methane&Ethane"
		if names != expectedNames {
			t.Errorf("extractFractions() Failed. Expected names = [%v] but found [%v]", expectedNames, names)
		}

		if slices.Compare(fractions, expectedFractions) != 0 {
			t.Errorf("extractFractions() Failed. Expected fractions = %v but found %v", expectedFractions, fractions)
		}

		t.Logf("In fluidStr: '%v'. fluid fractions found: %v for %v", fluidStr, fractions, names)
	}

	fluidStr = "Water-30%"

	if names, fractions, err := extractFractions(fluidStr); err != nil {
		t.Errorf("extractFractionsFailed(). fluidStr: '%v' should have fractions", fluidStr)
	} else {
		expectedFractions := []float64{0.3}
		expectedNames := "Water"
		if names != expectedNames {
			t.Errorf("extractFractions() Failed. Expected names = [%v] but found [%v]", expectedNames, names)
		}

		if slices.Compare(fractions, expectedFractions) != 0 {
			t.Errorf("extractFractions() Failed. Expected fractions = %v but found %v", expectedFractions, fractions)
		}

		t.Logf("In fluidStr: '%v'. fluid fractions found: %v for %v", fluidStr, fractions, names)
	}

	fluidStr = "Water-0.6"

	if names, fractions, err := extractFractions(fluidStr); err != nil {
		t.Errorf("extractFractionsFailed(). fluidStr: '%v' should have fractions", fluidStr)
	} else {
		expectedFractions := []float64{0.6}
		expectedNames := "Water"
		if names != expectedNames {
			t.Errorf("extractFractions() Failed. Expected names = [%v] but found [%v]", expectedNames, names)
		}

		if slices.Compare(fractions, expectedFractions) != 0 {
			t.Errorf("extractFractions() Failed. Expected fractions = %v but found %v", expectedFractions, fractions)
		}

		t.Logf("In fluidStr: '%v'. fluid fractions found: %v for %v", fluidStr, fractions, names)
	}
}
