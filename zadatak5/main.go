package main

import (
	"flag"
	"fmt"
	"regexp"
	"sort"
	"strconv"
)

func main() {
	inputPtr := flag.String("input", "", "Ulazni niz u formatu '[ 1, 2, 3, 4, 5 ]'")
	flag.Parse()

	if *inputPtr == "" {
		fmt.Println("Greška: Ulazni niz nije prosleđen.")
		return
	}

	// Regularni izraz za proveru ispravnosti formata unosa
	validInput := regexp.MustCompile(`^\[\s*(\d+\s*,\s*)*\d+\s*\]$`)

	if !validInput.MatchString(*inputPtr) {
		fmt.Println("Greška: Neispravan format unosa.")
		return
	}

	// Pretvaranje ulaznog niza u niz int-ova
	nums := extractNumbers(*inputPtr)

	// Deduplikacija u mestu (in-place)
	nums = deduplicate(nums)

	// Sortiranje niza
	sort.Ints(nums)

	// Prikazivanje izlaza na stdout
	fmt.Println(nums)
}

// Funkcija za izdvajanje brojeva iz ulaznog niza u obliku stringa i pretvaranje u niz int-ova
func extractNumbers(input string) []int {
	// Regularni izraz za pronalaženje brojeva u nizu
	numberRegex := regexp.MustCompile(`\d+`)
	matches := numberRegex.FindAllString(input, -1)

	nums := make([]int, len(matches))
	for i, match := range matches {
		num, err := strconv.Atoi(match)
		if err != nil {
			fmt.Printf("Greška pri pretvaranju broja: %v\n", err)
			return nil
		}
		nums[i] = num
	}

	return nums
}

// Funkcija za deduplikaciju niza u mestu (in-place)
func deduplicate(data []int)[]int {
	m := make(map[int]int8)
	currIndex := 0
	for i := 0; i < len(data); i++ {
		num := m[data[i]]
		if num == 0 {
			m[data[i]] = 1
			data[currIndex] = data[i]
			currIndex++
		}
	}
	return data[:currIndex]
}
