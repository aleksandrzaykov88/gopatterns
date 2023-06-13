package composite

import "testing"

func TestComposite(t *testing.T) {
	swimmer := CompositeSwimmerA{
		MySwim: Swim,
	}

	swimmer.MyAthlete.Train()
	swimmer.MySwim()

	fish := Shark{
		Swim: Swim,
	}

	fish.Eat()
	fish.Swim()

	swimmerB := CompositeSwimmerB{
		&Athlete{},
		&SwimmerImpl{},
	}

	swimmerB.Train()
	swimmerB.Swim()
}
