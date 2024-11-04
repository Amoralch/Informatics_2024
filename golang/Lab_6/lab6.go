package lab6

import (
	"fmt"
)

type Cat struct {
	Name  string
	Color string
	Age   int
}

func NewCat(name, color string, age int) Cat {
	return Cat{
		Name:  name,
		Color: color,
		Age:   age,
	}
}

func (c Cat) GetAge() int {
	return c.Age
}

func (c *Cat) SetAge(age int) {
	if age >= 0 {
		c.Age = age
		fmt.Printf("Возраст кошки %s установлен в %d лет.\n", c.Name, c.Age)
	} else {
		fmt.Println("Возраст не может быть отрицательным.")
	}
}

func (c Cat) DisplayInfo() {
	fmt.Printf("Имя: %s\nЦвет: %s\nВозраст: %d лет\n", c.Name, c.Color, c.Age)
}

func RunLab6() {
	cat := NewCat("Мурка", "Серый", 2)

	cat.DisplayInfo()

	age := cat.GetAge()
	fmt.Printf("Текущий возраст кошки: %d лет\n", age)

	cat.SetAge(3)

	cat.DisplayInfo()
}
