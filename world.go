package main

import (
	"github.com/google/uuid"
)

type Region struct {
	Id          uuid.UUID
	GrowthChart map[Form][]Form
	FoodPrefs   map[Form][]Food
}

func NewRegion(name string) (region Region) {
	region.Id, _ = uuid.NewUUID()
	region.GrowthChart = make(map[Form][]Form, 32)

	region.GrowthChart[Egg] = []Form{Baby}
	region.GrowthChart[Baby] = []Form{Head, Serpentine}

	region.GrowthChart[Serpentine] = []Form{Insectoid, Winged}
	region.GrowthChart[Head] = []Form{HeadArms, HeadLegs, Multibody}

	region.GrowthChart[Insectoid] = []Form{Winged}
	region.GrowthChart[HeadArms] = []Form{Bipedal, Winged}
	region.GrowthChart[HeadLegs] = []Form{Bipedal, Quadruped, Multiped, Winged}
	region.GrowthChart[Quadruped] = []Form{Multiped}

	region.FoodPrefs = make(map[Form][]Food, 32)
	for _, form := range FormValues() {
		region.FoodPrefs[form] = RandomFoodList(3)
	}

	return
}
