package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/rkoesters/xdg/basedir"
)

const ProjectName = "monpona"
const FilePerm = 0644

var region Region
var sanctuary Sanctuary

func main() {
	newMon := NewMon()
	fmt.Println(newMon.MonType)

	dataPath := filepath.Join(basedir.DataHome, ProjectName)
	regionPath := filepath.Join(dataPath, "region.json")
	sanctuaryPath := filepath.Join(dataPath, "sanctuary.json")

	regionBytes, err := os.ReadFile(regionPath)

	if err != nil {
		region = NewRegion("NewRegion")
		regEnc, _ := json.MarshalIndent(region, "", "\t")
		os.WriteFile(regionPath, regEnc, FilePerm)
	} else {
		json.Unmarshal(regionBytes, &region)
	}

	sanctuaryBytes, err := os.ReadFile(sanctuaryPath)

	if err != nil {
		sanctuary = NewSanctuary()
		sancEnc, _ := json.MarshalIndent(sanctuary, "", "\t")
		os.WriteFile(sanctuaryPath, sancEnc, FilePerm)
	} else {
		json.Unmarshal(sanctuaryBytes, &sanctuary)
	}

	if len(sanctuary.Mons) == 0 {
		newMon = NewMon()
		sanctuary.Mons = append(sanctuary.Mons, NewMon())
	}

	for _, mon := range sanctuary.Mons {
		fmt.Println(mon)
	}
}
