package mon

//go:generate go-enum --marshal --names --values --nocase --noprefix

import (
	"math/rand"

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
	Age       int
	Happiness byte // 0-24
	Fitness   byte // 0-24
	Fullness  byte // 0-24
	Wellness  byte // 0-24
}

func New() Mon {
	form := Form(rand.Intn(13))
	id, _ := uuid.NewUUID()

	return Mon{
		Id:        id,
		Name:      string(form),
		MonType:   form,
		Age:       0,
		Happiness: 5,
		Fitness:   128,
		Fullness:  128,
		Wellness:  64,
	}
}

func (mon *Mon) AgeUp() {
	mon.Age += 1

	switch mon.Age {
	case 60:

	}
}

func (mon *Mon) ChangeForm() {
	switch mon.MonType {
	case Egg:
		mon.MonType = Baby
	case Baby:
		if mon.Happiness > 128 && mon.Fitness > 50 {
			mon.MonType = Serpentine
		} else {
			mon.MonType = Head
		}
	case Head:
		switch {
		case mon.Happiness > 50 && mon.Fitness < 128:
			mon.MonType = Multibody
		case mon.Fullness > 50 && mon.Fitness < 128:
			mon.MonType = HeadLegs
		case mon.Happiness > 128 && mon.Fitness > 50:
			mon.MonType = HeadArms
		default:
			// mon.MonType = Head
		}
	case Serpentine:
		if mon.Happiness > 128 && mon.Fitness > 200 {
			mon.MonType = Winged
		} else if mon.Happiness > 128 && mon.Fullness < 50 {
			mon.MonType = Insectoid
		} else {
			// mon.MonType = Serpentine
		}
	case Insectoid:
		if mon.Fitness > 200 {
			mon.MonType = Winged
		} else {
			// mon.MonType = Insectoid
		}
	}
}

const (
	EventGrow = iota
)
