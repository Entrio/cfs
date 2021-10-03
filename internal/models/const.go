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

const (
	BuildingMaterial ResourceGroup = iota
	Edible
	Special
)
