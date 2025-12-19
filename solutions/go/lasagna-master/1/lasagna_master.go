package lasagna

// PreparationTime returns estimated prep
func PreparationTime(layers []string, minutes int) int {
    prepTimePerLayer := 2
    if minutes != 0 {
        prepTimePerLayer = minutes
    }
    return len(layers) * prepTimePerLayer
}

// Quantities function returns number of gram of noodes and liters of sauce
func Quantities(layers []string) (noodles int, sauce float64) {
    noodles = 0
    sauce = 0.0
    noodleCount := 0
    sauceCount := 0.0
    for _, v := range layers {
        if v == "noodles" {
            noodleCount++
            noodles = noodleCount * 50
        }
        if v == "sauce" {
          	sauceCount++
            sauce = sauceCount * 0.2
        }
    }
    return
}

// AddSecretIngredient() function replaces last item in myList with last item in friendsList
func AddSecretIngredient(friendsList []string, myList []string) {
    if len(myList) == 0 || len(friendsList) == 0 {
        panic("list is empty")
    }
    myList[len(myList) - 1] = friendsList[len(friendsList) - 1]
}

// ScaleRecipe() function does something
func ScaleRecipe(quantities []float64, noOfPortions int) []float64 {
    result := make([]float64, len(quantities))
    for i := range quantities {
        result[i] = quantities[i] * float64(noOfPortions) / 2
    }
    return result
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
