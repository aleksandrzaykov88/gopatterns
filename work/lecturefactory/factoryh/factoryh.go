package factoryh

import (
	"fmt"
)

type IMobile interface {
	Prepare()
	SetTariff()
	Box()
}

type Mobile struct{}

func (m Mobile) Prepare() {
	fmt.Println("Preparing to sell")
}

func (m Mobile) SetTariff() {
	fmt.Println("Setting internet plan")
}

func (m Mobile) Box() {
	fmt.Println("Packing")
}

////////////////////////////////////////////

type Samsung struct {
	Mobile
}

func (s Samsung) SetTariff() {
	fmt.Println("Setting 5G internet plan")
}

////////////////////////////////////////////

type IPhone struct {
	Mobile
}

// Mobile Store

type MobileStore struct{}

func (ms *MobileStore) OrderPhone(brand string) {

	phone, _ := ms.GetMobile(brand)

	phone.Prepare()
	phone.SetTariff()
	phone.Box()
}

// Factory

func (ms *MobileStore) GetMobile(brand string) (IMobile, error) {
	var phone IMobile
	switch brand {
	case "samsung":
		phone = Samsung{}
	case "iphone":
		phone = IPhone{}
	default:
		return nil, fmt.Errorf("there is no such brand: %s", brand)
	}

	return phone, nil
}

// Center Store
type CenterSamsung struct {
	Samsung
}

func (cs CenterSamsung) SetTariff() {
	fmt.Println("Setting CENTER 5G internet plan")
}

type CenterIphone struct {
	IPhone
}

type CenterMobileStore struct {
	MobileStore
}

// Ural Store

type UralSamsung struct {
	Samsung
}

func (us UralSamsung) SetTariff() {
	fmt.Println("Setting URAL 5G internet plan")
}

type UralIphone struct {
	IPhone
}

type UralMobileStore struct {
	MobileStore
}
