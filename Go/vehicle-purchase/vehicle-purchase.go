package purchase

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	return kind == "car" || kind == "truck";
	panic("NeedsLicense not implemented")
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) string {
    olSaying := " is clearly the better choice."
    if option1 < option2 {
        return option1 + olSaying
    } else {
    	return option2 + olSaying
    }
	panic("ChooseVehicle not implemented")
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	if age < 3 {
        return originalPrice * 0.80
    } else if age < 10 {
    	return originalPrice * 0.70
    } else {
    	return originalPrice * 0.50
    }
	panic("CalculateResellPrice not implemented")
}
