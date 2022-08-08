package main

import "fmt"

func main() {

	var (
		name, lastname string
		age, number    int

		name2, lastname2 string
		age2             int
	)

	name, lastname, age, number = "Pablo", "Picasso", 95, 100
	name2, lastname2, age = "Sigmund", "Freud", 83

	fmt.Println("Picasso:", name, lastname, age, number)
	fmt.Println("Freud  :", name2, lastname2, age2)

	type animal struct { //定义结构体
		color  string
		weight int
		voice  string
	}
	cat := animal{ //实例化结构体
		color:  "black",
		weight: 4,
		voice:  "喵喵喵",
	}
	var dog animal
	dog.color = "gray"
	dog.weight = 6
	dog.voice = "汪汪"
	fmt.Println(cat)
	fmt.Printf("猫的颜色:%s\n猫的体重:%d\n猫的声音:%s\n", cat.color, cat.weight, cat.voice)
	type person struct {
		name, lastname string
		age            int
	}

	picasso := person{
		name:     "Pablo",
		lastname: "Picasso",
		age:      91,
	}

	var freud person
	freud.name = "Sigmund"
	freud.lastname = "Freud"
	freud.age = 83

	fmt.Printf("\n%s's age is %d\n", picasso.lastname, picasso.age)

	fmt.Printf("\nPicasso: %#v\n", picasso)
	fmt.Printf("Freud  : %#v\n", freud)
}
