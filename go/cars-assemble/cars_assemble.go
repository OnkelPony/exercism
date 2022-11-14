package cars

const minsPerHour = 60

// carsPerHour is amount of cars produced per hour for speed = 1
const carsPerHour = 221

// SuccessRate is used to calculate the ratio of an item being created without
// error for a given speed
func SuccessRate(speed int) float64 {
	if speed == 0 {
		return 0.0
	} else if 1 <= speed && speed <= 4 {
		return 1.0
	} else if speed >= 5 && speed <= 8 {
		return 0.9
	} else {
		return 0.77
	}
}

// CalculateProductionRatePerHour for the assembly line, taking into account
// its success rate
func CalculateProductionRatePerHour(speed int) float64 {
	return float64(speed) * SuccessRate(speed) * carsPerHour
}

// CalculateProductionRatePerMinute describes how many working items are
// produced by the assembly line every minute
func CalculateProductionRatePerMinute(speed int) int {
	return int(CalculateProductionRatePerHour(speed) / minsPerHour)
}

// CalculateLimitedProductionRatePerHour describes how many working items are
// produced per hour with an upper limit on how many can be produced per hour
func CalculateLimitedProductionRatePerHour(speed int, limit float64) float64 {
	rate := CalculateProductionRatePerHour(speed)
	if rate > limit {
		rate = limit
	}
	return rate
}
