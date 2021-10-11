package models

import (
	"github.com/rs/zerolog/log"
	"time"
)

type (
	FarmCell struct {
		Owner                *Farm
		outputType           ResourceType // This is what the farm produces
		outputAmount         int          // This is how much of the ResourceType the cell outputs upon harvesting
		isReadyToBeHarvested bool         // Can the cell be harvested
		grownAmount          int64        // How much growing the cell has done
	}
)

func GenerateFarmCells(farm *Farm, qty int) []FarmCell {
	log.Trace().
		Str("owner", farm.id.String()).
		Int("type", int(farm.resourceType)).
		Int("qty", qty).
		Msg("Generating cells")
	cells := make([]FarmCell, qty)
	for c := range cells {
		cells[c].Owner = farm
	}
	return cells
}

func (fc FarmCell) GetGrownDurationAmount() time.Duration {
	return time.Millisecond * time.Duration(fc.grownAmount)
}
