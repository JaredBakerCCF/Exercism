//Package cars contains source files for the cars module
package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
    fProductionRate := float64(productionRate)
	return fProductionRate * (successRate / 100)
	panic("CalculateWorkingCarsPerHour not implemented")
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
    fProductionRate := float64(productionRate)
	workingCarsPerHour := fProductionRate * (successRate / 100)
    workingCarsPerMin := int(workingCarsPerHour / 60)
    return workingCarsPerMin
	panic("CalculateWorkingCarsPerMinute not implemented")
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	if carsCount / 10 == 0 {
		var carsForDiscountPrice int = carsCount / 10
        var carsWithDiscountPrice int = carsForDiscountPrice * 95000
        var lonerCars int = carsCount % 10
        var lonerCarsPrice int = lonerCars * 10000
        var totalCost int = carsWithDiscountPrice + lonerCarsPrice
        return uint(totalCost)
	} else if carsCount / 10 != 0 {
    	var carsForDiscountPrice int = carsCount / 10
        var carsWithDiscountPrice int = carsForDiscountPrice * 95000
        var lonerCars int = carsCount % 10
        var lonerCarsPrice int = lonerCars * 10000
        var totalCost int = carsWithDiscountPrice + lonerCarsPrice
        return uint(totalCost)
    }
	panic("CalculateCost not implemented")
}
