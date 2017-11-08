package animal

type Dog struct {
	Animal
}

func (d *Dog) Sound() string {
	return "I am a dog named " + d.name
}

func NewDog() *Dog {
	dog := new(Dog)
	dog.setThis(dog)
	return dog
}
