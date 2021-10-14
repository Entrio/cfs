package models

type ResourceType int
type ResourceGroup int

const (
	Water ResourceType = iota
	Wood
	WoodenPlanks
	Potatoes
	Carrots
	Berries
	CookedPotato
	CookedCarrot
)

var (
	ResourceTypeString = map[ResourceType]string{
		Water:        "water",
		Wood:         "wood",
		WoodenPlanks: "planks",
		Potatoes:     "potatoes",
		Carrots:      "carrots",
		Berries:      "berries",
		CookedPotato: "grilled potatoes",
		CookedCarrot: "steamed carrot",
	}
)

const (
	BuildingMaterial ResourceGroup = iota
	Edible
	Special
)
