package builderpattern

type IBuilder interface {
	SetWindowType(string)
	SetDoorType(string)
	SetNumOfFloor(int)
	GetHouse() House
}

func NewIBuilder(builderType string) IBuilder {
	if builderType == "normal" {
		return NewNormalBuilder()
	}

	if builderType == "igloo" {
		return NewIglooBuilder()
	}

	return nil
}
