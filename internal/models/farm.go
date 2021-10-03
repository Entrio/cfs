package models

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"time"
)

/*
Each farm split into "cells" and each cell has to be worked before anything can be grown
If the cell hasn't been worked, do not update it, otherwise update the grown amount
*/

type (
	Farm struct {
		cells        []FarmCell
		resourceType ResourceType
		growthPeriod time.Duration
	}
)

func (f *Farm) GetName() string {
	return "farm"
}

func (f *Farm) GetDescription() string {
	return "we are making tons of food"
}

// Update updates the farm and "grows" slots
func (f *Farm) Update(delta int64) {
	for c := range f.cells {

		if f.cells[c].isReadyToBeHarvested {
			// no point in doing anything we are ready to be harvested
			log.Trace().Str("cell", fmt.Sprintf("cell_%d", c)).Msg("ready to be harvested")
			continue
		}

		f.cells[c].grownAmount += delta
		if c == -1 {
			log.Trace().
				Str("cell", fmt.Sprintf("farm_%s_cell_%d", c)).
				Int64("grown", f.cells[c].grownAmount).
				Int64("total", int64(f.growthPeriod)).
				Msg("growing")
		}

		if f.cells[c].GetGrownDurationAmount() >= f.growthPeriod {
			// ready to be harvested
			f.cells[c].grownAmount = int64(f.growthPeriod)
			f.cells[c].isReadyToBeHarvested = true
			continue
		}

	}
}

func (f *Farm) SetCells(rt ResourceType) {
	f.resourceType = rt
}
