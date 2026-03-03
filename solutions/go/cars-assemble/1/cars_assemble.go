package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
    CalculateWorkingCarsPerHour := successRate * (float64(1) / float64(100)) * float64(productionRate)
	return CalculateWorkingCarsPerHour
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	CalculateWorkingCarsPerMinute := ((productionRate) * int(successRate + 0.5)) / int(6000)
    return CalculateWorkingCarsPerMinute
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	CalculateCost := ((carsCount / 10) * 95000) + ((carsCount % 10) * 10000)
    return uint(CalculateCost)
}
