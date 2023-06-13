package abstractfactory

type Motorbike interface {
	GetMotorbikeType() int
}

type SportMotorbike struct{}

func (*SportMotorbike) GetMotorbikeType() int {
	return SportMotorbikeType
}

func (*SportMotorbike) NumWheels() int {
	return 4
}

func (*SportMotorbike) NumSeats() int {
	return 5
}

type CruiseMotorbike struct{}

func (*CruiseMotorbike) GetMotorbikeType() int {
	return CruiseMotorbikeType
}

func (*CruiseMotorbike) NumWheels() int {
	return 4
}

func (*CruiseMotorbike) NumSeats() int {
	return 5
}
