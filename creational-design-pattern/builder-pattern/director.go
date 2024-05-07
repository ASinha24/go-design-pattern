package builderpattern

type Director struct {
	builder IBuilder
}

func NewDirector(b IBuilder) *Director {
	return &Director{
		builder: b,
	}
}

func (d *Director) SetBuilder(b IBuilder) {
	d.builder = b
}

func (d *Director) BuildHouse(doorType, windowType string, noOfFloor int) House {
	d.builder.SetDoorType(doorType)
	d.builder.SetWindowType(windowType)
	d.builder.SetNumOfFloor(noOfFloor)

	return d.builder.GetHouse()
}
