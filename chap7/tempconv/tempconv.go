package main

import (
	"flag"
	"fmt"
)

type celciusFlag struct {
	celcius float64
}

func (c *celciusFlag) String() string {
	return fmt.Sprintf("%gC", c.celcius)
}

func (c *celciusFlag) Set(s string) error {
	var value float64
	var unit string

	fmt.Sscanf(s, "%f%s", &value, &unit)
	switch unit {
	case "C":
		c.celcius = value
		return nil
	case "F":
		c.celcius = value * 9.0 / 5.0 + 32
		return nil
	}

	return fmt.Errorf("invalid unit %s", unit)
}

func main() {
	defaultVal := 100.0
	fl := celciusFlag{defaultVal}
	flag.Var(&fl, "temp", "set temperature")
	flag.Parse()
	fmt.Println(fl.celcius)
}
