package main

func main() {
	animal1 := getAnimal(AnimalCategory1)
	animal2 := getAnimal(AnimalCategory2)
	animal3 := getAnimal(AnimalCategory(7))
	UseAnimal(animal1)
	UseAnimal(animal2)
	UseAnimal(animal3)

}

func getAnimal(category AnimalCategory) Animal {
	switch category {
	case AnimalCategory1:
		return &Dog{
			Nick: "tiaotiao",
			Age:  0,
		}
	case AnimalCategory2:
		return &Pig{
			Nick: "tiaotiao",
			Age:  0,
		}
	default:

	}
	return &Undefined{Base{
		Nick: "Base",
		Age:  0,
	}}

}

func UseAnimal[T Animal](a T) {
	a.Eat("suibian")
	a.Name()
}
