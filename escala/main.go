package main

import "fmt"

const KELVIN_TEMP float64 = 373

func main() {
	tempCelsius := KELVIN_TEMP - 273

	fmt.Printf("A temperatura de ebulição da água em kelvin é %g°K e em Celsisus é %g°C.\n", KELVIN_TEMP, tempCelsius)
}
