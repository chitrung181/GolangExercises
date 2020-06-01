package convert

import "fmt"

func Convert(val float64, unit string) {
	switch unit {
	case "C":
		cToKAndF(val)
	case "m":
		mToFt(val)
	case "Kg":
		kgToLbs(val)
	default:
		cToKAndF(val)
		mToFt(val)
		kgToLbs(val)
	}
}

func cToKAndF(val float64) {
	fmt.Printf("%f°C = %f°K = %f°F\n", val, val*9/5+32, val+273.15)
}

func mToFt(val float64) {
	fmt.Printf("%fm = %fft\n", val, val*3.28)
}

func kgToLbs(val float64) {
	fmt.Printf("%fkg = %flbs\n", val, val*2.205)
}
