package main

import (
	"github.com/google/uuid"
)

type Region struct {
	Id          uuid.UUID
	GrowthChart map[Form][]Form
	FoodPrefs   map[Form][]Food
	Sanctuaries []Sanctuary
	Mons        []Mon
}

func NewRegion(name string) (region Region) {
	region.Id, _ = uuid.NewUUID()
	region.GrowthChart = make(map[Form][]Form, 32)

	region.GrowthChart[FormEgg] = []Form{FormBaby}
	region.GrowthChart[FormBaby] = []Form{FormHead, FormSerpentine}

	region.GrowthChart[FormSerpentine] = []Form{FormInsectoid, FormWinged}
	region.GrowthChart[FormHead] = []Form{FormHeadArms, FormHeadLegs, FormMultibody}

	region.GrowthChart[FormInsectoid] = []Form{FormWinged}
	region.GrowthChart[FormHeadArms] = []Form{FormBipedal, FormWinged}
	region.GrowthChart[FormHeadLegs] = []Form{FormBipedal, FormQuadruped, FormMultiped, FormWinged}
	region.GrowthChart[FormQuadruped] = []Form{FormMultiped}

	region.FoodPrefs = make(map[Form][]Food, 32)
	for _, form := range FormValues() {
		region.FoodPrefs[form] = RandomFoodList(3)
	}

	region.Sanctuaries = make([]Sanctuary, 1)
	region.Sanctuaries[0] = NewSanctuary()

	return
}

func (region Region) NewMonId() MonId {
	var id MonId
	monCount := len(region.Mons)
	if monCount == 0 {
		id = 0
	} else {
		id = region.Mons[monCount-1].Id
	}

	return id + 1
}

func (region Region) NewSanctuaryId() SanctuaryId {
	var id SanctuaryId
	sancCount := len(region.Sanctuaries)
	if sancCount == 0 {
		id = 0
	} else {
		id = region.Sanctuaries[sancCount-1].Id
	}

	return id + 1
}
