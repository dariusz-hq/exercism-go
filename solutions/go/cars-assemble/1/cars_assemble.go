package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	return float64(productionRate) * (successRate / 100)
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
    perHour := CalculateWorkingCarsPerHour(productionRate, successRate)
    perMinute := perHour / 60
	return int(perMinute)
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	unitCost := 10000
    groupCost := 95000
    modulo := carsCount % 10
    overallCost := groupCost * ((carsCount - modulo)/10) + unitCost * modulo
    return uint(overallCost)
    
}
