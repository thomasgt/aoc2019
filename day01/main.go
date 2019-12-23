package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func computeModuleFuel(mass int) int {
	return mass/3 - 2
}

func computeModuleFuelRecursive(mass int) int {
	fuelMass := mass/3 - 2
	if fuelMass > 0 {
		return fuelMass + computeModuleFuelRecursive(fuelMass)
	} else {
		return 0
	}
}

func main() {
	if len(os.Args) != 3 {
		log.Fatal("Usage: input part")
	}

	file, fileErr := os.Open(os.Args[1])
	if fileErr != nil {
		log.Fatal(fileErr)
	}
	defer file.Close()

	var part int
	_, partErr := fmt.Sscan(os.Args[2], &part)
	if partErr != nil {
		log.Fatal(partErr)
	}

	scanner := bufio.NewScanner(file)

	totalFuel := 0
	for scanner.Scan() {
		var mass int
		_, moduleErr := fmt.Sscan(scanner.Text(), &mass)
		if moduleErr != nil {
			log.Fatal(moduleErr)
		}

		switch part {
		case 1:
			totalFuel += computeModuleFuel(mass)
			break
		case 2:
			totalFuel += computeModuleFuelRecursive(mass)
			break
		default:
			log.Fatalf("Unknown part %d", part)
		}

	}

	log.Printf("Total Fuel: %d", totalFuel)
}
