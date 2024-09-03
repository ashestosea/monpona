package main

import (
	"math/rand"
)

//go:generate go-enum --marshal --names --values --nocase
/*ENUM(
RainbowGrass,
Ooznip,
Berry,
Groundnut,
Carnivorous
)
*/
type Food byte

func remove[L any](l []L, i int) []L {
	l[i] = l[len(l)-1]
	return l[:len(l)-1]
}

func RandomFoodList(maxCount int) (list []Food) {
	foodTypes := TypeValues()
	count := rand.Intn(maxCount) + 1

	for range count {
		randFood := rand.Intn(len(foodTypes))
		list = append(list, foodTypes[randFood])
		foodTypes = remove(foodTypes, randFood)
	}

	return
}
