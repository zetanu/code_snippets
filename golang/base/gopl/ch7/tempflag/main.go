package main

import (
	"flag"
	"fmt"
)

type celsiusFlag struct {
	d float64
}

func (f *celsiusFlag) Set(s string) error {
	var uint string
	var value float64
	fmt.Scanf(s, "%f%s", &value, &uint)
	switch uint {
	case "C", "c":
		f.d = value
		return nil
	case "F", "f":
		f.d = value * 2
		return nil
	}

	return fmt.Errorf("invalid %s", s)
}

func (f *celsiusFlag) String() string {
	return ""
}

func CelsiusFlag(name string, value float64, usage string) *float64 {
	f := celsiusFlag{value}
	flag.CommandLine.Var(&f, name, usage)
	return &f.d

}

var temp = CelsiusFlag("temp", 2, "the temperature")

func main() {
	flag.Parse()
	fmt.Println(*temp)
}
