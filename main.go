package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/rkoesters/xdg/basedir"
)

const ProjectName = "monpona"
const FilePerm = 0644

var region Region
var sanctuary Sanctuary

func main() {
	dataPath := filepath.Join(basedir.DataHome, ProjectName)
	regionPath := filepath.Join(dataPath, "region.json")
	regionBytes, err := os.ReadFile(regionPath)

	if err != nil {
		region = NewRegion("NewRegion")
		regEnc, _ := json.MarshalIndent(region, "", "\t")
		os.WriteFile(regionPath, regEnc, FilePerm)
	} else {
		json.Unmarshal(regionBytes, &region)
	}

	/* sanctuaryBytes, err := os.ReadFile(sanctuaryPath)

	if err != nil {
		sanctuary = NewSanctuary()
		sancEnc, _ := json.MarshalIndent(sanctuary, "", "\t")
		os.WriteFile(sanctuaryPath, sancEnc, FilePerm)
	} else {
		json.Unmarshal(sanctuaryBytes, &sanctuary)
	} */

	sanctuary = region.Sanctuaries[0]

	if len(sanctuary.Mons) == 0 {
		newMon := NewMon()
		newMon.Diet = FoodValues()
		region.Mons = append(region.Mons, newMon)
		sanctuary.Mons = append(sanctuary.Mons, newMon.Id)
		fmt.Println("New Mon ", newMon.Id, " ", newMon.MonForm)
	}

	// currentTime := time.Now()
	ticker := time.NewTicker(time.Second * 10)
	done := make(chan bool)

	for {
		select {
		case <-done:
			return
		case <-ticker.C:
			for _, mon := range region.Mons {
				mon.ChangeForm(sanctuary, region)
			}
		}
	}
}
