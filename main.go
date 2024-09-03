package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"path/filepath"

	"github.com/ashestosea/monpona/food"
	"github.com/ashestosea/monpona/mon"
	"github.com/google/uuid"
	"github.com/rkoesters/xdg/basedir"
)

const ProjectName = "monpona"
const FilePerm = 0644

func main() {
	newMon := mon.New()
	fmt.Println(newMon.MonType)

	dataPath := filepath.Join(basedir.DataHome, ProjectName)
	regionPath := filepath.Join(dataPath, "region.json")
	sanctuaryPath := filepath.Join(dataPath, "sanctuary.json")

	var region Region
	regionBytes, err := os.ReadFile(regionPath)

	if err != nil {
		region = NewRegion("NewRegion")
		regEnc, _ := json.MarshalIndent(region, "", "\t")
		os.WriteFile(regionPath, regEnc, FilePerm)
	} else {
		json.Unmarshal(regionBytes, &region)
	}

	var sanctuary Sanctuary
	sanctuaryBytes, err := os.ReadFile(sanctuaryPath)

	if err != nil {
		sanctuary = NewSanctuary()
		sancEnc, _ := json.MarshalIndent(sanctuary, "", "\t")
		os.WriteFile(sanctuaryPath, sancEnc, FilePerm)
	} else {
		json.Unmarshal(sanctuaryBytes, &sanctuary)
	}
}

type Sanctuary struct {
	Id           uuid.UUID
	Name         string
	RegionId     uuid.UUID
	Houses       []House
	Mons         []mon.Mon
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
		Mons:         make([]mon.Mon, 0),
		DepartedMons: make([]uuid.UUID, 0),
		DeadMons:     make([]uuid.UUID, 0),
	}
}

type Region struct {
	Id          uuid.UUID
	GrowthChart map[mon.Form][]mon.Form
	FoodPrefs   map[mon.Form][]food.Type
}

func NewRegion(name string) (region Region) {
	region.Id, _ = uuid.NewUUID()
	region.GrowthChart = make(map[mon.Form][]mon.Form, 32)

	region.GrowthChart[mon.Egg] = []mon.Form{mon.Baby}
	region.GrowthChart[mon.Baby] = []mon.Form{mon.Head, mon.Serpentine}

	region.GrowthChart[mon.Serpentine] = []mon.Form{mon.Insectoid, mon.Winged}
	region.GrowthChart[mon.Head] = []mon.Form{mon.HeadArms, mon.HeadLegs, mon.Multibody}

	region.GrowthChart[mon.Insectoid] = []mon.Form{mon.Winged}
	region.GrowthChart[mon.HeadArms] = []mon.Form{mon.Bipedal, mon.Winged}
	region.GrowthChart[mon.HeadLegs] = []mon.Form{mon.Bipedal, mon.Quadruped, mon.Multiped, mon.Winged}
	region.GrowthChart[mon.Quadruped] = []mon.Form{mon.Multiped}

	region.FoodPrefs = make(map[mon.Form][]food.Type, 32)
	for _, form := range mon.FormValues() {
		region.FoodPrefs[form] = RandomFoodList(3)
	}

	return
}

func remove[L any](l []L, i int) []L {
	l[i] = l[len(l)-1]
	return l[:len(l)-1]
}

func RandomFoodList(maxCount int) (list []food.Type) {
	foodTypes := food.TypeValues()
	count := rand.Intn(maxCount) + 1

	for range count {
		randFood := rand.Intn(len(foodTypes))
		list = append(list, foodTypes[randFood])
		foodTypes = remove(foodTypes, randFood)
	}

	return
}
