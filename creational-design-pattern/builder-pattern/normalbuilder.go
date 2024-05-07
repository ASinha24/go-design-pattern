package builderpattern

type NormalBuilder struct {
	House
}

func (nb *NormalBuilder) SetWindowType(windowType string) {
	nb.WindowType = windowType
}

func (nb *NormalBuilder) SetDoorType(doorType string) {
	nb.House.DoorType = doorType
}

func (nb *NormalBuilder) SetNumOfFloor(floor int) {
	nb.Floor = floor
}

func (nb *NormalBuilder) GetHouse() House {
	return nb.House
}

func NewNormalBuilder() *NormalBuilder {
	return &NormalBuilder{}
}
