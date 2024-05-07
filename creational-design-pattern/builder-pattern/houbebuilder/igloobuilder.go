package houbebuilder

type IglooBuilder struct {
	House
}

func (nb *IglooBuilder) SetWindowType(windowType string) {
	nb.WindowType = windowType
}

func (nb *IglooBuilder) SetDoorType(doorType string) {
	nb.House.DoorType = doorType
}

func (nb *IglooBuilder) SetNumOfFloor(floor int) {
	nb.Floor = floor
}

func (nb *IglooBuilder) GetHouse() House {
	return nb.House
}

func NewIglooBuilder() *IglooBuilder {
	return &IglooBuilder{}
}
