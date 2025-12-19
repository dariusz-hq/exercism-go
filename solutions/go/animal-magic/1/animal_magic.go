package chance

import "math/rand"
import "time"

// RollADie returns a random int d with 1 <= d <= 20.
func RollADie() int {
    result := 0
    for result < 1 {
        result = rand.Intn(20) 
    }
	return result
}

// GenerateWandEnergy returns a random float64 f with 0.0 <= f < 12.0.
func GenerateWandEnergy() float64 {
    rand.Seed(time.Now().UnixNano())
	return rand.Float64() * 12
}

// ShuffleAnimals returns a slice with all eight animal strings in random order.
func ShuffleAnimals() []string {
    slice := []string{"ant", "beaver", "cat", "dog", "elephant", "fox", "giraffe", "hedgehog"}
	rand.Shuffle(len(slice), func(i, j int){
        slice[i], slice[j] = slice[j], slice[i]
    })
    return slice
}
