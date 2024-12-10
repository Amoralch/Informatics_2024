package lab7

import (
	"fmt"
)

type Product interface {
	GetPrice() float64
	SetPrice(float64)
	ApplyDiscount(float64)
	GetDetails() string
}

type BaseProduct struct {
	Name  string
	Price float64
}

func (p *BaseProduct) GetPrice() float64 {
	return p.Price
}

func (p *BaseProduct) SetPrice(price float64) {
	p.Price = price
}

func (p *BaseProduct) ApplyDiscount(discount float64) {
	if discount < 0 || discount > 100 {
		fmt.Println("Скидка должна быть от 0 до 100%")
		return
	}
	p.Price *= (1 - discount/100)
}

func (p *BaseProduct) GetDetails() string {
	return fmt.Sprintf("Название: %s, Цена: %.2f", p.Name, p.Price)
}

type Electronics struct {
	BaseProduct
	Warranty int
}

func (e *Electronics) GetDetails() string {
	return fmt.Sprintf("%s, Гарантия: %d года", e.BaseProduct.GetDetails(), e.Warranty)
}

type Clothing struct {
	BaseProduct
	Size  string
	Color string
}

func (c *Clothing) GetDetails() string {
	return fmt.Sprintf("%s, Размер: %s, Цвет: %s", c.BaseProduct.GetDetails(), c.Size, c.Color)
}

func TotalPrice(products []Product) float64 {
	var total float64
	for _, p := range products {
		total += p.GetPrice()
	}
	return total
}

func RunLab7() {
	laptop := &Electronics{
		BaseProduct: BaseProduct{Name: "Ноутбук", Price: 50000},
		Warranty:    2,
	}
	tshirt := &Clothing{
		BaseProduct: BaseProduct{Name: "Футболка", Price: 1500},
		Size:        "M",
		Color:       "Черный",
	}

	products := []Product{laptop, tshirt}

	fmt.Println("Список товаров:")
	for _, p := range products {
		fmt.Println(p.GetDetails())
	}

	fmt.Printf("\nОбщая стоимость до скидок: %.2f\n", TotalPrice(products))

	laptop.ApplyDiscount(10)
	tshirt.ApplyDiscount(20)

	fmt.Println("\nСписок товаров после скидок:")
	for _, p := range products {
		fmt.Println(p.GetDetails())
	}

	fmt.Printf("\nОбщая стоимость после скидок: %.2f\n", TotalPrice(products))
}
