package conv

import (
	"fmt"
)

type Celsius float64
type Fahrenheit float64
type Kelvin float64
type Feet float64
type Meter float64
type Pound float64
type Kilo    float64

const (
	AbsoluteZeroC Celsius = -273.15
	StartK        Kelvin  = 273.15
	FreezingC     Celsius = 0
	Boiling       Celsius = 100
	OneMeter	  Feet    = 3.28083
	OneFeet       Meter   = 0.3048
	OnePound      Kilo    = 0.45359237
	OneKilo       Pound   = 2.20462262
)

func (c Celsius) String() string {
	return fmt.Sprintf("%g°C", c)
}

func (f Fahrenheit) String() string {
	return fmt.Sprintf("%g°F", f)
}

func (k Kelvin) String() string {
	return fmt.Sprintf("%g°K", k)
}