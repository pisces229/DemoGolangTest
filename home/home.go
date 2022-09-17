package home

//go:generate mockgen -destination ../mock/home/animal.go -package mock_home demo.golang.test/home Animal

type Animal interface {
	GetName() string
	GetSpeed(base int) int
}

type Dog struct {
	Name string
}

func (the *Dog) GetName() string {
	return the.Name
}
func (the *Dog) GetSpeed(base int) int {
	return base * 20
}

type Cat struct {
	Name string
}

func (the *Cat) GetName() string {
	return the.Name
}
func (the *Cat) GetSpeed(base int) int {
	return base * 10
}

type Home struct {
	Animal Animal
}

func NewHome(animal Animal) *Home {
	return &Home{Animal: animal}
}
func (the *Home) GetAnimalName() string {
	// the.Animal.Get_Name()
	return the.Animal.GetName()
}
func (the *Home) GetAnimalSpeed(base int) int {
	the.Animal.GetName()
	// the.Animal.Get_Speed()
	return the.Animal.GetSpeed(base)
}
