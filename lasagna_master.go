package lasagna

func PreparationTime(layers []string, preparationTime int) int {
	if preparationTime == 0 {
		preparationTime = 2
	}

	return len(layers) * preparationTime
}

func Quantities(layers []string) (int, float64) {
	noodlesQuantityPerLayer := 50
	noodlesLayers := 0
	sauceQuantityPerLayer := 0.2
	sauceLayers := 0

	for _, layer := range layers {
		if layer == "noodles" {
			noodlesLayers++
		}
		if layer == "sauce" {
			sauceLayers++
		}
	}

	return noodlesQuantityPerLayer * noodlesLayers, sauceQuantityPerLayer * float64(sauceLayers)
}

func AddSecretIngredient(friendIngredientList []string, myIngredientList []string) {
	myIngredientList[len(myIngredientList)-1] = friendIngredientList[len(friendIngredientList)-1]
}

func ScaleRecipe(amount []float64, portions int) []float64 {
	amountCopy := make([]float64, len(amount))

	copy(amountCopy, amount)

	for i := 0; i < len(amountCopy); i++ {
		amountCopy[i] *= float64(portions) / 2
	}

	return amountCopy
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
