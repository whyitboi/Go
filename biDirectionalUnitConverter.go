package main

import (
	"fmt"
	"math"
)

//Named type declaration for units
type (
	Feet        float64
	Centimeters float64
	Celsius     float64
	Fahrenheit  float64
	Minutes     float64
	Seconds     float64
	Radian      float64
	Degrees     float64
	Kilograms   float64
	Pounds      float64
)

//Converter that defines data types of units
type Converter struct {
	feet       Feet
	centimeter Centimeters
	celsius    Celsius
	fahrenheit Fahrenheit
	minute     Minutes
	second     Seconds
	radian     Radian
	degree     Degrees
	kilogram   Kilograms
	pound      Pounds
}

//CelsiusToFahrenheit for celsius to degree conversion
func (cvr Converter) CelsiusToFahrenheit(c Celsius) Fahrenheit {
	return Fahrenheit((c * 1.8) + 32)
}

//FahrenheitToCelsius for degree to celsius conversion
func (cvr Converter) FahrenheitToCelsius(f Fahrenheit) Celsius {
	return Celsius((f - 32.0) * 0.5555555)

}

//MinutesToSeconds for Minutes to Second conversion
func (cvr Converter) MinutesToSeconds(m Minutes) Seconds {
	return Seconds(m * 60)
}

//SecondsToMinutes for Seconds to Minute conversion
func (cvr Converter) SecondsToMinutes(s Seconds) Minutes {
	return Minutes(s / 60)
}

//FeetToCentimeter for Feet to Centimeter conversion
func (cvr Converter) FeetToCentimeter(f Feet) Centimeters {
	return Centimeters(f * 30.48)
}

//CentimeterToFeet for Centimeter to Feet conversion
func (cvr Converter) CentimeterToFeet(c Centimeters) Feet {
	return Feet(c / 30.48)
}

//RadianToDegree for radian to degree conversion
func (cvr Converter) RadianToDegree(r Radian) Degrees {
	return Degrees(r * (180 / math.Pi))
}

//DegreeToRadian for degree to radian conversion
func (cvr Converter) DegreeToRadian(d Degrees) Radian {
	return Radian(d * (math.Pi / 180))
}

//KilogramToPound for kilogram to pound conversion
func (cvr Converter) KilogramToPound(k Kilograms) Pounds {
	return Pounds(k * 2.205)
}

//PoundToKilogram for pound to kilogram conversion
func (cvr Converter) PoundToKilogram(p Pounds) Kilograms {
	return Kilograms(p / 2.205)
}

var cvr Converter

func main() {

	fmt.Println(cvr.FeetToCentimeter(5))
	fmt.Println(cvr.CentimeterToFeet(152.4))
	fmt.Println(cvr.MinutesToSeconds(5))
	fmt.Println(cvr.SecondsToMinutes(120))
	fmt.Println(cvr.CelsiusToFahrenheit(23))
	fmt.Println(cvr.FahrenheitToCelsius(73.4))
	fmt.Println(cvr.RadianToDegree(23))
	fmt.Println(cvr.DegreeToRadian(100))
	fmt.Println(cvr.PoundToKilogram(100))
	fmt.Println(cvr.KilogramToPound(100))

}
