package main

import (
	"github.com/ashestosea/monpona/food"
	"github.com/google/uuid"
)

type House struct {
	monId uuid.UUID
	food  []food.Type
}

var houses []House
var departedMons []uuid.UUID
var deadMons []uuid.UUID
