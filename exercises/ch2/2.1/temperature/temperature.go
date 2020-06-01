package temperature

import "fmt"

type Celsius float64
type Fahrenheit float64
type Kelvin float64

//return string
func (c Celsius) ToString() string {
	return fmt.Sprintf("%g°C", c)
}

func (f Fahrenheit) ToString() string {
	return fmt.Sprintf("%g°F", f)
}

func (k Kelvin) ToString() string {
	return fmt.Sprintf("%g°K", k)
}

//Convert functions
func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

func CToK(c Celsius) Kelvin {
	return Kelvin(c + 273.15)
}
