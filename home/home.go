package home

//go:generate mockgen -destination ../mock/animal.go -package mock demo.golang.test/home Animal

type Animal interface {
	Get_Name() string
	Get_Speed(base int) int
}

type Dog struct {
	Name string
}

func (the *Dog) Get_Name() string {
	return the.Name
}
func (the *Dog) Get_Speed(base int) int {
	return base * 20
}

type Cat struct {
	Name string
}

func (the *Cat) Get_Name() string {
	return the.Name
}
func (the *Cat) Get_Speed(base int) int {
	return base * 10
}

type Home struct {
	Animal Animal
}

func NewHome(animal Animal) *Home {
	return &Home{Animal: animal}
}
func (the *Home) Get_AnimalName() string {
	// the.Animal.Get_Name()
	return the.Animal.Get_Name()
}
func (the *Home) Get_AnimalSpeed(base int) int {
	the.Animal.Get_Name()
	// the.Animal.Get_Speed()
	return the.Animal.Get_Speed(base)
}
