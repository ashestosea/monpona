package main

//go:generate go-enum --marshal --names --values --nocase

import (
	"github.com/google/uuid"
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

type Mon struct {
	Id        uuid.UUID
	Name      string
	MonType   Form
	Age       int    // Minutes
	Happiness byte   // 0-24
	Fitness   byte   // 0-24
	Fullness  byte   // 0-24
	Wellness  byte   // 0-24
	Diet      []Food // Food eaten this growth stage
}

func NewMon() Mon {
	// form := Form(rand.Intn(int(len(FormValues()))))
	form := Egg
	id, _ := uuid.NewUUID()

	return Mon{
		Id:        id,
		Name:      "Eggbert",
		MonType:   form,
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
	switch mon.MonType {
	case Egg:
		mon.MonType = Baby
	case Baby:

	case Head:
	case Serpentine:
	case Insectoid:
	}
}

const (
	EventGrow = iota
)
