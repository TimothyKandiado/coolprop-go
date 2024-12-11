package coolprop_go

import (
	"fmt"
	"github.com/TimothyKandiado/coolprop-go/internal"
	"strconv"
	"strings"
)

// PropsSI Return a value that depends on the thermodynamic state
// @param Output The output parameter, one of "T","D","H",etc.
// @param Name1 The first state variable name, one of "T","D","H",etc.
// @param Prop1 The first state variable value
// @param Name2 The second state variable name, one of "T","D","H",etc.
// @param Prop2 The second state variable value
// @param FluidName The fluid name
func PropsSI(output string, name1 string, property1 float64, name2 string, property2 float64, reference string) (float64, error) {
	//backend, fluid := extractBackend(reference)

	return 0, nil
}

func _propsSIMulti(outputs []string, name1 string, property1 []float64, name2 string, property2 []float64, backend string, fluids []string, fractions []float64) []float64 {
	return nil
}

/*
@brief Extract the backend from a string - something like "HEOS::Water" would split to "HEOS" and "Water".  If no backend is specified, the backend will be set to "?"

	@param fluid_string The input string
	@param backend The output backend, if none found, "?"
	@param fluid The output fluid string (minus the backend string)
*/
func extractBackend(fluidString string) (string, string) {
	var backend string
	var fluid string

	if index, found := hasBackendInString(fluidString); found {
		backend = fluidString[0:index]
		fluid = fluidString[index+2:]
	} else {
		backend = "?"
		fluid = fluidString
	}

	if internal.DEBUG_PRINT_ON {
		fmt.Printf("backend extracted. backend: '%v'. fluid: '%v'.\n", backend, fluid)
	}

	return backend, fluid
}

func hasBackendInString(fluidString string) (int, bool) {
	index := strings.Index(fluidString, "::")

	found := index != -1

	return index, found
}

/*
@brief Extract fractions (molar, mass, etc.) encoded in the string if any

	@param fluid_string The input string
	@param fractions The fractions
	@return The fluids, as a '&' delimited string
*/
func extractFractions(fluid string) (string, []float64, error) {
	var fractions []float64

	var fluidStr string

	if hasFractionsInString(fluid) {
		fractions = make([]float64, 0)
		names := make([]string, 0)

		// Break up into pairs - like "Ethane[0.5]&Methane[0.5]" -> ("Ethane[0.5]","Methane[0.5]")
		pairs := strings.Split(fluid, "&")

		for _, fluid := range pairs {
			if fluid[len(fluid)-1] != ']' {
				err := fmt.Errorf("fluid entry [%v] must end with ']' character", fluid)
				return "", nil, err
			}

			// Split at '[', but first remove the ']' from the end by taking a substring
			nameFraction := strings.Split(fluid[:len(fluid)-1], "[")

			if len(nameFraction) != 2 {
				err := fmt.Errorf("could not break [%v] into name/fraction", fluid[:len(fluid)-1])
				return "", nil, err
			}

			name := nameFraction[0]
			fractionStr := nameFraction[1]

			// convert fraction to float64
			fraction, err := strconv.ParseFloat(fractionStr, 64)

			if err != nil {
				return "", nil, fmt.Errorf("could not parse fraction [%v] to float", fluid)
			}

			fractions = append(fractions, fraction)
			names = append(names, name)
		}

		if internal.DEBUG_PRINT_ON {
			fmt.Printf("Detected fractions of %v for %v\n", fractions, names)
		}

		fluidStr = strings.Join(names, "&")
	} else if hasSolutionConcentration(fluid) {
		fractions = make([]float64, 0)
		var concentration float64

		fluidParts := strings.Split(fluid, "-")

		if len(fluidParts) != 2 {
			err := fmt.Errorf("Format of incompressible solution string [%v] is invalid, should be like \"EG-20%v\" or \"EG-0.2\" ", fluid, '%')
			return "", nil, err
		}

		concentrationStr := fluidParts[1]
		isPercentageConcentration := false
		if concentrationStr[len(concentrationStr)-1] == '%' {
			concentrationStr = concentrationStr[:len(concentrationStr)-1]
			isPercentageConcentration = true
		}

		concentration, err := strconv.ParseFloat(concentrationStr, 64)
		if err != nil {
			return "", nil, fmt.Errorf("could not parse concentration [%v] to float", fluid)
		}

		if isPercentageConcentration {
			concentration *= 0.01
		}

		fractions = append(fractions, concentration)
		fluidStr = fluidParts[0]

		if internal.DEBUG_PRINT_ON {
			fmt.Printf("Detected incompressible concentration of %v for %v.", concentration, fluidParts[0])
		}
	} else {
		fractions = make([]float64, 1)
		fractions[0] = 1.0

		fluidStr = fluid
	}

	return fluidStr, fractions, nil
}

func hasFractionsInString(fluidStr string) bool {
	// If can find both "[" and "]", it must have mole fractions encoded as string
	return strings.ContainsRune(fluidStr, '[') && strings.ContainsRune(fluidStr, ']')
}

func hasSolutionConcentration(fluidStr string) bool {
	// If can find "-", expect mass fractions encoded as string
	return strings.ContainsRune(fluidStr, '-')
}
