package main

import (
	"fmt"
	"math/rand"
	"os"
	"path/filepath"

	"github.com/ashestosea/monpona/food"
	"github.com/ashestosea/monpona/mon"
	"github.com/rkoesters/xdg/basedir"
	"github.com/ugorji/go/codec"
)

const ProjectName = "monpona"

func main() {
	newMon := mon.New()
	fmt.Println(newMon.MonType)

	dataPath := filepath.Join(basedir.DataHome, ProjectName)

	region, err := os.ReadFile(filepath.Join(dataPath, "region.json"))

	if err != nil {
		panic(err.Error())
	}

	// var b []byte = make([]byte, 0, 64)
	// regionBytes := region.r
	var h codec.Handle = new(codec.JsonHandle)
	growthChartDec := make(map[mon.Form][]mon.Form, 32)
	var dec *codec.Decoder = codec.NewDecoderBytes(region, h)
	err = dec.Decode(growthChartDec)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("growth chart decode_____")
	for key, val := range growthChartDec {
		fmt.Println(mon.Form(key))
		for i := range len(val) {
			fmt.Println("  ", mon.Form(val[i]))
		}
	}
}

type Region struct {
	GrowthChart map[mon.Form][]mon.Form
	FoodPrefs   map[mon.Form][]food.Type
}

func NewRegion(name string) (region Region) {
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
	for _, f := range mon.FormValues() {
		region.FoodPrefs[f] = RandomFoodList(3)
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

	for range rand.Intn(count) {
		randFood := rand.Intn(len(foodTypes))
		list = append(list)
		foodTypes = remove(foodTypes, randFood)
	}

	return
}

func RandGrowthList(currentMonForm mon.Form, newEdgesCount int) []mon.Form {
	result := make([]mon.Form, 0, newEdgesCount)
	monForms := make([]mon.Form, 0, 10)
	monFormsLen := len(mon.FormValues())

	fmt.Println("rand_growth_list___________________________")
	fmt.Println("current mon.Form = ", currentMonForm)
	fmt.Println("from ", int(mon.Baby)+1, " to ", monFormsLen)
	for i := int(mon.Baby) + 1; i < monFormsLen; i++ {
		if i != int(currentMonForm) {
			fmt.Println("appending ", mon.Form(i), "(int", i, ")", " to monForms (len ", len(monForms), ")")
			monForms = append(monForms, mon.Form(i))
		}
	}

	for i := range monForms {
		fmt.Println(monForms[i])
	}

	for range newEdgesCount {
		randInt := rand.Intn(len(monForms))
		result = append(result, monForms[randInt])
		monForms = Remove(monForms, randInt)
		fmt.Println("Adding edge ", randInt, " to ", currentMonForm, "_________________")
		fmt.Println("monForms now equals")
		for i := range monForms {
			fmt.Println("  ", monForms[i])
		}
	}

	return result
}
