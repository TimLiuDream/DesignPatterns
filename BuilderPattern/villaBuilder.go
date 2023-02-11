package main

type VillaBuilder struct {
	windowType string
	doorType   string
	floor      int
}

func newVillaBuilder() *VillaBuilder {
	return &VillaBuilder{}
}

func (b *VillaBuilder) setWindowType() {
	b.windowType = "Snow Window"
}

func (b *VillaBuilder) setDoorType() {
	b.doorType = "Snow Door"
}

func (b *VillaBuilder) setNumFloor() {
	b.floor = 1
}

func (b *VillaBuilder) getHouse() House {
	return House{
		doorType:   b.doorType,
		windowType: b.windowType,
		floor:      b.floor,
	}
}
