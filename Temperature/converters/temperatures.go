package converters

/*Fahrenheit temperature*/
type Fahrenheit float64

/*Celcious temperature*/
type Celcious float64

/*Kelvin temperature*/
type Kelvin float64

/*Celcious Fahrenheit to Celcious*/
func (f Fahrenheit) Celcious() Celcious {
	return Celcious((f - 32) / 1.8)
}

/*Fahrenheit Celcious to Fahrenheit*/
func (c Celcious) Fahrenheit() Fahrenheit {
	return Fahrenheit((c * 1.8) + 32)
}

/*Celcious Kelvin to Celcious*/
func (k Kelvin) Celcious() Celcious {
	return Celcious(k - 273.15)
}

/*Kelvin Celcious to Kelvin */
func (c Celcious) Kelvin() Kelvin {
	return Kelvin(c + 273.15)
}
