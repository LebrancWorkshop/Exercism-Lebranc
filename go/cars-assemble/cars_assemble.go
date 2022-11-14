package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	return (float64(productionRate) * successRate) / 100  

}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	workingCarPerHour := (float64(productionRate) * successRate) / 100; 
	return int(workingCarPerHour / 60) 
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	groupOfTen := (carsCount / 10) * 95000
	individual := (carsCount % 10) * 10000
	return uint(groupOfTen + individual)  
}
