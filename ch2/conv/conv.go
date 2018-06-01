package conv

func C2F(c Celsius) Fahrenheit {
	return Fahrenheit(c * 9 / 5 + 32)
}

func F2C(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

func C2K(c Celsius) Kelvin {
	return StartK + Kelvin(c)
}

func F2K(f Fahrenheit) Kelvin {
	return C2K(F2C(f))
}

func K2C(k Kelvin) Celsius  {
	return Celsius(k - StartK)
}

func K2F(k Kelvin) Fahrenheit {
	return C2F(K2C(k))
}

func Meter2Feet(m Meter) Feet {
	return Feet(m) * OneMeter
}

func Feet2Meter(f Feet) Meter {
	return Meter(f) * OneFeet
}

func Kilo2Pound(k Kilo) Pound {
	return Pound(k) * OneKilo
}

func Pound2Kilo(p Pound) Kilo {
	return Kilo(p) * OnePound
}