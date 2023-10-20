package main

import "fmt"

type House struct {
	family     []Human
	appliances []Appliances
	furniture  []Furniture
}
type Human struct {
	name string
	age  int64
	role string
}

type Sizes struct {
	Meters      int64
	Centimeters int64
	Millimeters int64
}
type Appliances struct {
	name        string
	material    string
	appointment string
	width       Sizes
	height      Sizes
	depth       Sizes
}
type Furniture struct {
	name     string
	material string
	width    Sizes
	height   Sizes
	depth    Sizes
}

var appliances1 Appliances = Appliances{
	name:        "Стиральная машинка",
	material:    "Сталь",
	appointment: "Стирать вещи",
	width: Sizes{
		Meters:      1,
		Centimeters: 20,
		Millimeters: 0,
	},
	height: Sizes{
		Meters:      0,
		Centimeters: 80,
		Millimeters: 0,
	},
	depth: Sizes{
		Meters:      0,
		Centimeters: 60,
		Millimeters: 0,
	},
}
var appliances2 Appliances = Appliances{
	name:        "Холодильник",
	material:    "Пластик",
	appointment: "Хранить продукты",
	width: Sizes{
		Meters:      1,
		Centimeters: 50,
		Millimeters: 0,
	},
	height: Sizes{
		Meters:      1,
		Centimeters: 80,
		Millimeters: 0,
	},
	depth: Sizes{
		Meters:      0,
		Centimeters: 70,
		Millimeters: 0,
	},
}
var appliances3 Appliances = Appliances{
	name:        "Телевизор",
	material:    "Пластик",
	appointment: "Просмотр телепередач",
	width: Sizes{
		Meters:      1,
		Centimeters: 20,
		Millimeters: 9,
	},
	height: Sizes{
		Meters:      0,
		Centimeters: 80,
		Millimeters: 0,
	},
	depth: Sizes{
		Meters:      0,
		Centimeters: 20,
		Millimeters: 0,
	},
}
var appliances4 Appliances = Appliances{
	name:        "Микроволновая печь",
	material:    "Металл",
	appointment: "Подогрев еды",
	width: Sizes{
		Meters:      0,
		Centimeters: 50,
		Millimeters: 0,
	},
	height: Sizes{
		Meters:      0,
		Centimeters: 30,
		Millimeters: 0,
	},
	depth: Sizes{
		Meters:      0,
		Centimeters: 40,
		Millimeters: 0,
	},
}
var appliances5 Appliances = Appliances{
	name:        "Пылесос",
	material:    "Пластик",
	appointment: "Уборка",
	width: Sizes{
		Meters:      0,
		Centimeters: 40,
		Millimeters: 0,
	},
	height: Sizes{
		Meters:      0,
		Centimeters: 80,
		Millimeters: 0,
	},
	depth: Sizes{
		Meters:      0,
		Centimeters: 30,
		Millimeters: 0,
	},
}

var furniture1 Furniture = Furniture{
	name:     "Кровать",
	material: "Дерево",
	width: Sizes{
		Meters:      1,
		Centimeters: 80,
		Millimeters: 0,
	},
	height: Sizes{
		Meters:      0,
		Centimeters: 60,
		Millimeters: 0,
	},
	depth: Sizes{
		Meters:      2,
		Centimeters: 0,
		Millimeters: 0,
	},
}
var furniture2 Furniture = Furniture{
	name:     "Шкаф",
	material: "Дерево",
	width: Sizes{
		Meters: 1,
	},
	height: Sizes{
		Meters: 2,
	},
	depth: Sizes{
		Meters: 1,
	},
}
var furniture3 Furniture = Furniture{
	name:     "Стол",
	material: "Дерево",
	width: Sizes{
		Meters:      1,
		Centimeters: 0,
		Millimeters: 0,
	},
	height: Sizes{
		Meters:      0,
		Centimeters: 80,
		Millimeters: 0,
	},
	depth: Sizes{
		Meters:      0,
		Centimeters: 80,
		Millimeters: 0,
	},
}
var furniture4 Furniture = Furniture{
	name:     "Диван",
	material: "Ткань",
	width: Sizes{
		Meters:      2,
		Centimeters: 0,
		Millimeters: 0,
	},
	height: Sizes{
		Meters:      0,
		Centimeters: 90,
		Millimeters: 0,
	},
	depth: Sizes{
		Meters:      1,
		Centimeters: 0,
		Millimeters: 0,
	},
}
var furniture5 Furniture = Furniture{
	name:     "Стул",
	material: "Дерево",
	width: Sizes{
		Meters:      0,
		Centimeters: 50,
		Millimeters: 0,
	},
	height: Sizes{
		Meters:      0,
		Centimeters: 80,
		Millimeters: 0,
	},
	depth: Sizes{
		Meters:      0,
		Centimeters: 50,
		Millimeters: 0,
	},
}

var human1 Human = Human{
	name: "Фанил",
	age:  60,
	role: "муж",
}
var human2 Human = Human{
	name: "Ирина",
	age:  56,
	role: "жена",
}
var human3 Human = Human{
	name: "Тагир",
	age:  21,
	role: "ребенок",
}
var house House = House{
	furniture:  []Furniture{furniture1, furniture2, furniture3, furniture4, furniture5},
	appliances: []Appliances{appliances1, appliances2, appliances3, appliances4, appliances5},
	family:     []Human{human1, human2, human3},
}

func printHouseInfo(h House) {
	fmt.Println("Мебель:")
	for _, f := range h.furniture {
		fmt.Println("Название:", f.name)
		fmt.Println("Материал:", f.material)
		fmt.Println("Ширина:", f.width.Meters, "м", f.width.Centimeters, "см", f.width.Millimeters, "мм")
		fmt.Println("Высота:", f.height.Meters, "м", f.height.Centimeters, "см", f.height.Millimeters, "мм")
		fmt.Println("Глубина:", f.depth.Meters, "м", f.depth.Centimeters, "см", f.depth.Millimeters, "мм")
		fmt.Println()
	}

	fmt.Println("Бытовая техника:")
	for _, a := range h.appliances {
		fmt.Println("Название:", a.name)
		fmt.Println("Материал:", a.material)
		fmt.Println("Назначение:", a.appointment)
		fmt.Println("Ширина:", a.width.Meters, "м", a.width.Centimeters, "см", a.width.Millimeters, "мм")
		fmt.Println("Высота:", a.height.Meters, "м", a.width.Centimeters, "см", a.width.Millimeters, "мм")
		fmt.Println("Глубина:", a.depth.Meters, "м", a.width.Centimeters, "см", a.width.Millimeters, "мм")
		fmt.Println()
	}

	fmt.Println("Семья:")
	for _, p := range h.family {
		fmt.Println("Имя:", p.name)
		fmt.Println("Возраст:", p.age)
		fmt.Println("Роль:", p.role)
		fmt.Println()
	}
}

func main() {
	printHouseInfo(house)
}
