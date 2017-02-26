package unitconv

import "fmt"

type Celsius float64
type Fahrenheit float64
type Meters float64
type Feet float64
type Kilometers float64
type Miles float64
type Pounds float64
type Kilograms float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func (c Celsius) String() string {
	return fmt.Sprintf("%g°C", c)
}

func (f Fahrenheit) String() string {
	return fmt.Sprintf("%g°F", f)
}

func (m Meters) String() string {
	return fmt.Sprintf("%g m", m)
}

func (ft Feet) String() string {
	return fmt.Sprintf("%g ft", ft)
}

func (km Kilometers) String() string {
	return fmt.Sprintf("%g km", km)
}

func (mi Miles) String() string {
	return fmt.Sprintf("%g mi", mi)
}

func (lbs Pounds) String() string {
	return fmt.Sprintf("%g lbs", lbs)
}

func (kg Kilograms) String() string {
	return fmt.Sprintf("%g kg", kg)
}
