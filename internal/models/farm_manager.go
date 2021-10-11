package models

import (
	"github.com/rs/zerolog/log"
	"sync"
)

type (
	FarmManager struct {
		farms []Farm
		mu    sync.Mutex
	}
)

func newFarmManager() *FarmManager {
	return &FarmManager{
		farms: make([]Farm, 0),
		mu:    sync.Mutex{},
	}
}

func (fm *FarmManager) update(delta int64) {
	for f := range fm.farms {
		fm.farms[f].Update(delta)
	}
}

// AddFarm creates a new farm and adds to the manager
func (fm *FarmManager) AddFarm(farm Farm) error {
	fm.mu.Lock()
	defer fm.mu.Unlock()
	//TODO: Validate the new farm
	fm.farms = append(fm.farms, farm)
	log.Info().
		Str("name", farm.GetID().String()).
		Str("description", farm.GetDescription()).
		Int("cells", len(farm.cells)).
		Msg("new farm added")
	log.Debug().Int("farms", len(fm.farms)).Msg("farms statistics")
	return nil
}
