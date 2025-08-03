package main

import (
	"fmt"
	"github.com/google/uuid"
)

type SanctuaryId int

type Sanctuary struct {
	Id           SanctuaryId
	Name         string
	Houses       []House
	Mons         []MonId
	DepartedMons []MonId
	DeadMons     []MonId
	FoodSources  []FoodSource
}

func NewSanctuary() Sanctuary {
	id := region.NewSanctuaryId()
	return Sanctuary{
		Id:           id,
		Name:         fmt.Sprint("Sanctuary ", id),
		Houses:       make([]House, 0),
		Mons:         make([]MonId, 0),
		DepartedMons: make([]MonId, 0),
		DeadMons:     make([]MonId, 0),
		FoodSources:  make([]FoodSource, 0),
	}
}

type House struct {
	monId uuid.UUID
	food  []Food
}

func (sanc *Sanctuary) PlantFoodSource(foodType Food, amount int) {
	sanc.FoodSources = append(sanc.FoodSources, FoodSource{FoodType: foodType, Remaining: amount})
}
