package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, times int) int {
	if times == 0 {
		return len(layers) * 2;
	}	
	return len(layers) * times; 
}

// TODO: define the 'Quantities()' function
func Quantitites(layers []string) (int, float64) {
	noodleAmount := 0;
	sauceAmount := 0.0;

	for _, layer := range layers {
		if layer == "noodles" {
			noodleAmount += 1;
		}

		if layer == "sauce" {
			sauceAmount += 1.0; 
		}
	}

	noodleQuant := noodleAmount * 50;
	sauceQuant := sauceAmount * 0.2;
	return noodleQuant, sauceQuant; 
}

// TODO: define the 'AddSecretIngredient()' function

// TODO: define the 'ScaleRecipe()' function
