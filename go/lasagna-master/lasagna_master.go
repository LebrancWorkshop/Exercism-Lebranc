package lasagna

import (
	"fmt" 
)

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, times int) int {
	if times == 0 {
		return len(layers) * 2;
	}	
	return len(layers) * times; 
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (int, float64) {
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
func AddSecretIngredient(friendsList []string, myList []string) {
	myList[len(myList) - 1] = friendsList[len(friendsList) - 1]; 
	fmt.Println(myList); 
} 

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(inputList []float64, portions int) []float64 {
	newInputList := make([]float64, 0);  

	for _, v := range inputList {
		newInputList = append(newInputList, (v / 2.0) * float64(portions));  
	}

	return newInputList; 
}
