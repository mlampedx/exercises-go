package unitconv

// CToF converts a Celsius temperature to Fahrenheit
func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

// FToC converts a Fahrenheit temperature to Celsius
func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

// MToFt converts a measurement in Meters to a measurement in Feet
func MToFt(m Meters) Feet {
	return Feet(m * 3)
}

// FtToM converts a measurement in Feet to a measurement in Meters
func FtToM(ft Feet) Meters {
	return Meters(ft / 3)
}

// KmToMi converts a measurement in Kilometers to a measurement in Miles
func KmToMi(km Kilometers) Miles {
	return Miles(km * 0.621371)
}

// MiToKm converts a measurement in Miles to a measurement in Kilometers
func MiToKm(mi Miles) Kilometers {
	return Kilometers(mi * 1.60934)
}

// LbsToKg converts a weight in Pounds to a weight in Kilograms
func LbsToKg(lbs Pounds) Kilograms {
	return Kilograms(lbs * 0.453592)
}

// KgToLbs converts a weight in Kilograms to a weight in Pounds
func KgToLbs(kg Kilograms) Pounds {
	return Pounds(kg * 2.20462)
}
