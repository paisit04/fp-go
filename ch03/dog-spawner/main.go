package main

import "fmt"

type (
	Name          string
	Breed         int
	Gender        int
	NameToDogFunc func(Name) Dog
)

// define possible breeds
const (
	Bulldog Breed = iota
	Havanese
	Cavalier
	Poodle
)

// define possible genders
const (
	Male Gender = iota
	Female
)

// define a Dog struct
type Dog struct {
	Name   Name
	Breed  Breed
	Gender Gender
}

var (
	maleHavaneseSpawner = DogSpawner(Havanese, Male)
	femalePoodleSpawner = DogSpawner(Poodle, Female)
)

func DogSpawner(breed Breed, gender Gender) NameToDogFunc {
	return func(name Name) Dog {
		return Dog{
			Name:   name,
			Breed:  breed,
			Gender: gender,
		}
	}
}

func main() {
	bucky := maleHavaneseSpawner("bucky")
	rocky := maleHavaneseSpawner("rocky")
	tipsy := femalePoodleSpawner("tipsy")

	fmt.Printf("%v\n", bucky)
	fmt.Printf("%v\n", rocky)
	fmt.Printf("%v\n", tipsy)
}

func createDogsWithoutPartialApplication() {
	bucky := Dog{
		Name:   "Bucky",
		Breed:  Havanese,
		Gender: Male,
	}

	rocky := Dog{
		Name:   "Rocky",
		Breed:  Havanese,
		Gender: Male,
	}

	tipsy := Dog{
		Name:   "Tipsy",
		Breed:  Poodle,
		Gender: Female,
	}
	_ = bucky
	_ = rocky
	_ = tipsy
}
