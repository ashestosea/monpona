package main

//go:generate go-enum --marshal --names --values --nocase

import (
	"math/rand"

	"fmt"
)

/*
	ENUM(

Egg,
Baby,
Head,
Serpentine,
HeadArms,
HeadLegs,
Quadruped,
Multiped,
Insectoid,
Winged,
Aquatic,
Bipedal,
Multibody,
)
*/
type Form byte

type MonId int

type Mon struct {
	Id        MonId
	Name      string
	MonForm   Form
	Age       int    // Minutes
	Happiness byte   // 0-24
	Fitness   byte   // 0-24
	Fullness  byte   // 0-24
	Wellness  byte   // 0-24
	Diet      []Food // Food eaten this growth stage
}

func NewMon() Mon {
	// form := Form(rand.Intn(int(len(FormValues()))))
	form := FormEgg
	id := region.NewMonId()

	return Mon{
		Id:        id,
		Name:      "Eggbert",
		MonForm:   form,
		Age:       0,
		Happiness: 12,
		Fitness:   12,
		Fullness:  12,
		Wellness:  24,
		Diet:      make([]Food, 0),
	}
}

func (mon *Mon) AgeUp(sanctuary Sanctuary, region Region) {
	mon.Age += 1

	switch mon.Age {
	case 2:
		mon.ChangeForm(sanctuary, region)
	}
}

func (mon *Mon) ChangeForm(sanctuary Sanctuary, region Region) {
	dietCounts := make(map[Food]int)
	for _, food := range mon.Diet {
		dietCounts[food] += 1
	}

	fmt.Println("CHANGE FORM____________________")

	if mon.MonForm == FormEgg {
		mon.MonForm = FormBaby
	} else {
		newFormWeights := make(map[Form]float32)
		var total float32

		for _, newForm := range region.GrowthChart[mon.MonForm] {
			for _, foodPref := range region.FoodPrefs[newForm] {
				weight := float32(dietCounts[foodPref])
				newFormWeights[newForm] = weight
			}
		}

		for _, weight := range newFormWeights {
			total += weight
		}

		fmt.Println("NewFormWeights")
		for form, weight := range newFormWeights {
			fmt.Println("  form = ", form, " weight = ", weight)
		}
		fmt.Println("   total = ", total)

		fmt.Println("Selecting new form_____")
		var remainingDist float32 = rand.Float32() * total
		for newForm, weight := range newFormWeights {
			remainingDist -= weight
			fmt.Println("  remainingDist = ", remainingDist, " form = ", newForm)
			if remainingDist < 0 {
				mon.MonForm = newForm
				break
			}
		}
	}

	fmt.Println("newForm = ", mon.MonForm)
}

const (
	EventGrow = iota
)
