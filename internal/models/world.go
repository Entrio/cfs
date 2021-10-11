package models

import (
	"github.com/google/uuid"
	"time"
)

type (
	WorldMeta struct {
		Name        string
		Description string
	}

	resources struct {
	}

	World struct {
		WorldMeta
		quit        chan struct{}
		count       int
		farmManager *FarmManager
	}
)

var (
	timeDTPrev = time.Now()
)

func (w *World) Quit() chan struct{} {
	return w.quit
}

func (w *World) GetName() string {
	return w.WorldMeta.Name
}

func (w *World) GetDescription() string {
	return w.WorldMeta.Description
}

func NewWorld(meta WorldMeta) *World {
	return &World{
		WorldMeta:   meta,
		quit:        make(chan struct{}),
		farmManager: newFarmManager(),
	}
}

// ProcessUpdate finds all the objects in the world and updates them
func (w *World) ProcessUpdate() error {
	timeSinceLastFrame := time.Now().Sub(timeDTPrev).Milliseconds()

	//timeSinceLastFrame = float64(int(timeSinceLastFrame*100) / 100)

	// Update all dependencies
	w.farmManager.update(timeSinceLastFrame)

	timeDTPrev = time.Now()
	return nil
}

func (w *World) AddFarm(resourceType ResourceType) error {
	f := Farm{
		id:           uuid.New(),
		growthPeriod: time.Second * 20,
	}
	f.cells = GenerateFarmCells(&f, 50)
	f.SetCells(resourceType)

	return w.farmManager.AddFarm(f)
}
