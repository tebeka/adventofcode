package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strconv"
)

func main() {
	masses, err := parseFile(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	fuel := 0.0
	for _, mass := range masses {
		v := calcFuel(mass)
		fuel += v + fuelFuel(v)
	}
	fmt.Println(int(fuel))
}

func parseFile(r io.Reader) ([]float64, error) {
	var masses []float64
	s := bufio.NewScanner(r)
	for s.Scan() {
		mass, err := strconv.ParseFloat(s.Text(), 64)
		if err != nil {
			return nil, err
		}
		masses = append(masses, mass)
	}

	return masses, s.Err()
}

func calcFuel(mass float64) float64 {
	return math.Floor(mass/3.0) - 2
}

func fuelFuel(fuel float64) float64 {
	total := 0.0
	for {
		fuel = calcFuel(fuel)
		if fuel <= 0 {
			break
		}
		total += fuel
	}
	return total
}
