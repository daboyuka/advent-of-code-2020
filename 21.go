package aoc2020

import (
	"fmt"
	"sort"
	"strings"

	. "aoc2020/helpers"
)

type strset map[string]bool

func makeStrset(strs []string) strset {
	ss := make(strset, len(strs))
	for _, s := range strs {
		ss[s] = true
	}
	return ss
}

func (ss strset) Intersect(others ...strset) (ss3 strset) {
	ss3 = make(strset, len(ss))
loop:
	for s := range ss {
		for _, other := range others {
			if !other[s] {
				continue loop
			}
		}
		ss3[s] = true
	}
	return ss3
}

func (ss strset) Any() (s string) {
	for s = range ss {
		return s
	}
	return ""
}

type food struct {
	Ingreds   strset
	Allergens strset
}

func ParseFoods(lines []string) (foods []food) {
	for _, line := range lines {
		ingredsStr, allergensStr := Split2(line[:len(line)-1], " (contains ")
		foods = append(foods, food{
			Ingreds:   makeStrset(strings.Split(ingredsStr, " ")),
			Allergens: makeStrset(strings.Split(allergensStr, ", ")),
		})
	}
	return foods
}

func FindAllergenIngredientMapping(foods []food) (allergenToIngred, ingredToAllergen map[string]string) {
	allergenToIngred, ingredToAllergen = make(map[string]string), make(map[string]string)

	possibleIngredsByAllergen := make(map[string]strset) // [allergen] -> [possible ingreds]
	for _, f := range foods {
		for allergen := range f.Allergens {
			if ingreds, ok := possibleIngredsByAllergen[allergen]; ok {
				possibleIngredsByAllergen[allergen] = f.Ingreds.Intersect(ingreds)
			} else {
				possibleIngredsByAllergen[allergen] = f.Ingreds
			}
		}
	}

	for len(possibleIngredsByAllergen) > 0 {
		for allergen, possibleIngreds := range possibleIngredsByAllergen {
			if len(possibleIngreds) == 1 {
				ingred := possibleIngreds.Any()

				fmt.Println(allergen, "->", ingred)
				allergenToIngred[allergen] = ingred
				ingredToAllergen[ingred] = allergen

				delete(possibleIngredsByAllergen, allergen)
				for _, otherPossibleIngreds := range possibleIngredsByAllergen {
					delete(otherPossibleIngreds, ingred)
				}
			}
		}
	}

	return allergenToIngred, ingredToAllergen
}

func Problem21a(lines []string) {
	foods := ParseFoods(lines)

	_, ingredToAllergen := FindAllergenIngredientMapping(foods)

	unallergicIngredAppearances := 0
	for _, f := range foods {
		for ingred := range f.Ingreds {
			if _, ok := ingredToAllergen[ingred]; !ok {
				unallergicIngredAppearances++
			}
		}
	}
	fmt.Println(unallergicIngredAppearances)
}

func Problem21b(lines []string) {
	foods := ParseFoods(lines)

	allergenToIngred, _ := FindAllergenIngredientMapping(foods)

	allAllergens := make([]string, 0, len(allergenToIngred))
	for allergen := range allergenToIngred {
		allAllergens = append(allAllergens, allergen)
	}
	sort.Strings(allAllergens)

	dangerousIngreds := make([]string, len(allAllergens))
	for i, allergen := range allAllergens {
		dangerousIngreds[i] = allergenToIngred[allergen]
	}

	fmt.Println(strings.Join(dangerousIngreds, ","))
}
