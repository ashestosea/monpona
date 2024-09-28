package main

import (
	"github.com/google/uuid"
)

type Sanctuary struct {
	Id           uuid.UUID
	Name         string
	RegionId     uuid.UUID
	Houses       []House
	Mons         []uuid.UUID
	DepartedMons []uuid.UUID
	DeadMons     []uuid.UUID
}

func NewSanctuary() Sanctuary {
	id, _ := uuid.NewUUID()
	return NewSanctuaryInRegion(id)
}

func NewSanctuaryInRegion(regionId uuid.UUID) Sanctuary {
	id, _ := uuid.NewUUID()
	return Sanctuary{
		Id:           id,
		Name:         "New Sanctuary",
		RegionId:     regionId,
		Houses:       make([]House, 0),
		Mons:         make([]uuid.UUID, 0),
		DepartedMons: make([]uuid.UUID, 0),
		DeadMons:     make([]uuid.UUID, 0),
	}
}

type House struct {
	monId uuid.UUID
	food  []Food
}
